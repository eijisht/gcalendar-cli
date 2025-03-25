package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gcal-cli/cmd"
	"gcal-cli/internal"

	"github.com/joho/godotenv"
)

// TODO: add caching for API requests (could use SQLite)
// TODO: help commands

var CREATE_COMMAND string = "write"
var READ_COMMAND string = "read"
var UPDATE_COMMAND string = "update"
var DELETE_COMMAND string = "remove"

func main() {
	command := parseCommand()

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
		fmt.Printf("Hello from create case!\n")
	case "r":
		flags := parseReadRequest()
		cmd.Read(srv, *flags.Calendar, *flags.Count, *flags.Days)
	case "":
		fmt.Printf("No command given. For help, use <gcal help>")
	default:
		fmt.Printf("Unknown command %s", command)
	}

}

func parseCommand() string {
	if len(os.Args) < 2 {
		return ""
	}

	commandMap := map[string]string{
		CREATE_COMMAND: "c",
		READ_COMMAND:   "r",
		UPDATE_COMMAND: "u",
		DELETE_COMMAND: "d",
	}

	if short, exists := commandMap[os.Args[1]]; exists {
		return short
	}

	return os.Args[1]
}

type readRequest struct {
	Calendar *string
	Count    *int64
	Days     *int64
}

func parseReadRequest() readRequest {
	calendar := flag.String("c", "primary", "usage: -c <calendar id>")
	count := flag.Int64("n", 10, "usage: -n <int>")
	day := flag.Int64("d", -1, "usage: -d <int>")

	// TODO: Improve usage messages

	flag.Parse()

	return readRequest{
		calendar,
		count,
		day,
	}
}
