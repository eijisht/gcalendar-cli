package internal

import (
	"flag"
	"log"
	"time"
)

const (
	createCommand string = "create"
	readCommand   string = "read"
	updateCommand string = "update"
	deleteCommand string = "remove"
	resetCommand  string = "reset"
	initCommand   string = "init"
)

func ParseCommand(args []string) string {
	if len(args) < 2 {
		return ""
	}

	commandMap := map[string]string{
		createCommand: "c",
		readCommand:   "r",
		updateCommand: "u",
		deleteCommand: "d",
		resetCommand:  "reset",
		initCommand:   "init",
	}

	if short, exists := commandMap[args[1]]; exists {
		return short
	}

	return args[1]
}

type ReadRequest struct {
	Calendar *string
	Count    *int64
	Days     *int64
}

func ParseReadRequest(args []string) ReadRequest {

	readFlags := flag.NewFlagSet("readFlags", flag.ContinueOnError)
	readFlags.ErrorHandling()

	calendar := readFlags.String("n", "primary", "usage: -c <calendar id>")
	count := readFlags.Int64("c", 10, "usage: -n <int>")
	day := readFlags.Int64("d", -1, "usage: -d <int>")

	// TODO: Improve usage messages
	// TODO: program panics if given the wrong type of flag (e.g string count)

	err := readFlags.Parse(args[2:])

	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	return ReadRequest{
		calendar,
		count,
		day,
	}
}

type CreateRequest struct {
	Calendar    *string
	Summary     *string
	Description *string
	End         *string // can be just a date
	Start       *string // can be just a date

	// TODO:
	// Attendees
	// Color ID
}

func ParseCreateRequest(args []string) CreateRequest {
	createFlags := flag.NewFlagSet("createFlags", flag.ContinueOnError)

	now := time.Now().Format("2006-01-02 15:04")
	calendar := createFlags.String("n", "primary", "")
	summary := createFlags.String("s", "CLI Event", "")
	description := createFlags.String("desc", "", "")
	startTime := createFlags.String("start", now, "")
	endTime := createFlags.String("end", now, "")

	err := createFlags.Parse(args[2:])

	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	return CreateRequest{
		calendar,
		summary,
		description,
		endTime,
		startTime,
	}
}
