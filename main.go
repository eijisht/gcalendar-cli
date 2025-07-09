package main

import (
	"fmt"
	"log"
	"os"

	"gcal-cli/cmd"
	"gcal-cli/internal"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(0)
	command := internal.ParseCommand(os.Args)

	if command == "reset" {
		err := os.Remove("token.json")
		if err != nil {
			fmt.Printf("Could not remove token.json, %s\n", err)
			return
		}
		fmt.Printf("Removed token.json\n")
		return
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	srv, err := internal.GetCalendarService()
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar service: %v", err)
	}

	// test for invalid token
	_, err = srv.Events.List("primary").MaxResults(1).Do()
	if err != nil {
		log.Fatalf("Token invalid or expired: %v\nTry running 'gcalcli reset' to get another token.", err)
	}

	// returns anonymus function to handle the command
	handlers := map[string]func(){
		"c": func() {
			flags := internal.ParseCreaterequest(os.Args)
			fmt.Println(*flags.Calendar, *flags.End, *flags.Start, *flags.Summary)
		},

		"r": func() {
			flags := internal.ParseReadRequest(os.Args)
			cmd.Read(srv, *flags.Calendar, *flags.Count, *flags.Days)
		},
	}

	if handler, found := handlers[command]; found {
		handler()
	} else if command == "" {
		fmt.Println("No command given. For help, use <gcal help>")
	} else {
		fmt.Printf("Unknown command %s\n", command)
	}

}
