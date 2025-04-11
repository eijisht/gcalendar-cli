package main

import (
	"fmt"
	"log"
	"os"

	"gcal-cli/cmd"
	"gcal-cli/internal"

	"github.com/joho/godotenv"
)

// TODO: add caching for API requests (could use SQLite)
// TODO: help commands

func main() {
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

	switch command {
	case "c":
		flags := internal.ParseCreaterequest(os.Args)
		fmt.Println(*flags.Calendar, *flags.End, *flags.Start, *flags.Summary)
	case "r":
		flags := internal.ParseReadRequest(os.Args)
		// Read request segfaults instead of throwing an error if token is too old
		cmd.Read(srv, *flags.Calendar, *flags.Count, *flags.Days)
	case "":
		fmt.Printf("No command given. For help, use <gcal help>\n")
	default:
		fmt.Printf("Unknown command %s\n", command)
	}

}
