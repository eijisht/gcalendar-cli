package internal

import (
	"context"
	"encoding/json"
	"fmt"
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
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	printWelcomeMessage()
	fmt.Printf("Open this URL in a browser and authorize: \n%s\n", authURL)

	var authCode string
	fmt.Print("Enter the authorization code: ")
	fmt.Scan(&authCode)

	tok, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		return nil, fmt.Errorf("could not exchange user token: %s\n", err)
	}

	return tok, nil
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

func GetCalendarService() (*calendar.Service, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, fmt.Errorf("unable to read credentials.json: %w", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)

	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file: %w", err)
	}

	client, err := getClient(config)
	if err != nil {
		return nil, fmt.Errorf("unable to get client: %s", err)
	}

	srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))

	return srv, err
}

func printWelcomeMessage() {
	fmt.Printf("")
}
