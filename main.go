package main

import (
	"fmt"
	"log"
	"os"

	"gcal-cli/cmd"
	"gcal-cli/internal"
)

func main() {
	log.SetFlags(0)
	command := internal.ParseCommand(os.Args)

	if command == "" || command == "help" {
		printUsage()
		return
	}

	if command == "reset" {
		err := os.Remove("token.json")
		if err != nil {
			fmt.Printf("Could not remove token.json, %s\n", err)
			return
		}
		fmt.Printf("Removed token.json\n")
		return
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
			flags := internal.ParseCreateRequest(os.Args)
			if err := cmd.Create(srv, *flags.Calendar, *flags.Summary, *flags.Description, *flags.Start, *flags.End); err != nil {
				log.Fatalf("create failed: %v", err)
			}
		},

		"r": func() {
			flags := internal.ParseReadRequest(os.Args)
			if err := cmd.Read(srv, *flags.Calendar, *flags.Count, *flags.Days); err != nil {
				log.Fatalf("read failed: %v", err)
			}
		},

		"d": func() {
			if len(os.Args) < 3 {
				log.Fatalf("usage: gcal remove <eventID>")
			}
			if err := cmd.Delete(srv, "primary", os.Args[2]); err != nil {
				log.Fatalf("delete failed: %v", err)
			}
		},
	}

	if handler, found := handlers[command]; found {
		handler()
	} else {
		fmt.Printf("Unknown command. Run 'gcal help' for usage.\n")
	}
}

func printUsage() {
	fmt.Println(`gcal - Google Calendar CLI

Usage:
  gcal read   [-n calendar] [-c count] [-d days]
  gcal create -s summary [-desc text] -start "2006-01-02 15:04" -end "2006-01-02 15:04"
  gcal remove <eventID>
  gcal reset  (delete the cached auth token)
  gcal help`)
}
