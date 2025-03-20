package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// TODO: add input sanitization
// TODO: get the user auth code from the http request automatically

func getClient(config *oauth2.Config) *http.Client {
	token := "token.json"
	tok, err := tokenFromFile(token)

	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(token, tok)
	}

	return config.Client(context.Background(), tok)
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

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Open this URL in a browser and authorize: \n%s\n", authURL)

	var authCode string
	fmt.Print("Enter the authorization code: ")
	fmt.Scan(&authCode)

	tok, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	return tok
}

func saveToken(file string, token *oauth2.Token) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func GetCalendarService() (*calendar.Service, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to parse client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)

	if err != nil {
		log.Fatalf("Unable to parse client secret file: %v", err)
	}

	client := getClient(config)

	return calendar.NewService(context.TODO(), option.WithHTTPClient(client))
}
