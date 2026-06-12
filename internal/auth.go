package internal

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// TODO: add input sanitization
// TODO: get the user auth code from the http request automatically

func getClient(config *oauth2.Config) (*http.Client, error) {
	token := "token.json"
	tok, err := tokenFromFile(token)

	if err != nil {
		tok, err = getTokenFromWeb(config)
		if err != nil {
			return nil, fmt.Errorf("could not get token from web: %s\n", err)
		}
		saveToken(token, tok)
	}
	ctx := context.Background()
	tokenSource := config.TokenSource(ctx, tok)
	client := oauth2.NewClient(ctx, tokenSource)

	return client, nil
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	// Listen on a random free loopback port for Google's redirect.
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, fmt.Errorf("could not start local server: %w", err)
	}
	defer listener.Close()

	config.RedirectURL = fmt.Sprintf("http://%s/callback", listener.Addr().String())

	// Random state value to guard against CSRF on the callback.
	state, err := randomState()
	if err != nil {
		return nil, err
	}

	codeCh := make(chan string, 1)
	errCh := make(chan error, 1)

	mux := http.NewServeMux()
	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if query.Get("state") != state {
			http.Error(w, "state mismatch", http.StatusBadRequest)
			errCh <- fmt.Errorf("state mismatch in oauth callback")
			return
		}
		code := query.Get("code")
		if code == "" {
			http.Error(w, "missing authorization code", http.StatusBadRequest)
			errCh <- fmt.Errorf("no authorization code in oauth callback")
			return
		}
		fmt.Fprintln(w, "Authentication complete. You can close this tab and return to the terminal.")
		codeCh <- code
	})

	server := &http.Server{Handler: mux}
	go server.Serve(listener)
	defer server.Shutdown(context.Background())

	authURL := config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	printWelcomeMessage()
	fmt.Printf("Open this URL in your browser to authorize:\n%s\n\nWaiting for authorization...\n", authURL)

	var code string
	select {
	case code = <-codeCh:
	case err = <-errCh:
		return nil, err
	}

	tok, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("could not exchange user token: %w", err)
	}

	return tok, nil
}

func randomState() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("could not generate state: %w", err)
	}
	return hex.EncodeToString(b), nil
}

func saveToken(file string, token *oauth2.Token) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("Unable to retrieve token from web: %v", err)
	}

	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return nil
}

func loadConfig() (*oauth2.Config, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, fmt.Errorf("unable to read credentials.json: %w", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file: %w", err)
	}

	return config, nil
}

// InitToken runs the OAuth flow and writes token.json, replacing any existing
// token. Use it for first-time setup or to re-authenticate.
func InitToken() error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	tok, err := getTokenFromWeb(config)
	if err != nil {
		return err
	}

	return saveToken("token.json", tok)
}

func GetCalendarService() (*calendar.Service, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	client, err := getClient(config)
	if err != nil {
		return nil, fmt.Errorf("unable to get client: %s", err)
	}

	return calendar.NewService(context.Background(), option.WithHTTPClient(client))
}

func printWelcomeMessage() {
	fmt.Printf("")
}
