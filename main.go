package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gcal-cli/cmd"
	"gcal-cli/internal"

	"github.com/joho/godotenv"
)

// TODO: add caching for API requests (could use SQLite)
// TODO: help commands

var CREATE_COMMAND string = "create"
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
		flags := parseCreaterequest()
		fmt.Println(*flags.Calendar, *flags.End, *flags.Start, *flags.Summary)
	case "r":
		flags := parseReadRequest()
		// Read request segfaults instead of throwing an error if token is too old
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

type ReadRequest struct {
	Calendar *string
	Count    *int64
	Days     *int64
}

// Flags are broken

func parseReadRequest() ReadRequest {
	calendar := flag.String("c", "primary", "usage: -c <calendar id>")
	count := flag.Int64("n", 10, "usage: -n <int>")
	day := flag.Int64("d", -1, "usage: -d <int>")

	// TODO: Improve usage messages

	flag.Parse()

	return ReadRequest{
		calendar,
		count,
		day,
	}
}

type CreateRequest struct {
	Calendar *string
	Summary  *string
	End      *string // can be just a date
	Start    *string // can be just a date

	// TODO:
	// Attendees
	// Color ID
}

func parseCreaterequest() CreateRequest {
	calendar := flag.String("c", "primary", "")
	summary := flag.String("n", "CLI Event", "")
	endTime := flag.String("end", time.Now().Format(time.RFC3339), "")
	startTime := flag.String("start", time.Now().Format(time.RFC3339), "")

	flag.Parse()
	for _, arg := range flag.Args() {
		fmt.Println(arg)
	}

	fmt.Println("calendar = ", *calendar)

	return CreateRequest{
		calendar,
		summary,
		endTime,
		startTime,
	}
}
