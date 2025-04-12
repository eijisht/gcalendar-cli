package internal

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var CREATE_COMMAND string = "create"
var READ_COMMAND string = "read"
var UPDATE_COMMAND string = "update"
var DELETE_COMMAND string = "remove"

var RESET_COMMAND string = "reset"

func ParseCommand(args []string) string {
	if len(args) < 2 {
		return ""
	}

	commandMap := map[string]string{
		CREATE_COMMAND: "c",
		READ_COMMAND:   "r",
		UPDATE_COMMAND: "u",
		DELETE_COMMAND: "d",
		RESET_COMMAND:  "reset",
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

	calendar := readFlags.String("c", "primary", "usage: -c <calendar id>")
	count := readFlags.Int64("n", 10, "usage: -n <int>")
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
	Calendar *string
	Summary  *string
	End      *string // can be just a date
	Start    *string // can be just a date

	// TODO:
	// Attendees
	// Color ID
}

func ParseCreaterequest(args []string) CreateRequest {
	createFlags := flag.NewFlagSet("createFlags", flag.ContinueOnError)

	calendar := createFlags.String("c", "primary", "")
	summary := createFlags.String("n", "CLI Event", "")
	endTime := createFlags.String("end", time.Now().Format(time.RFC3339), "")
	startTime := createFlags.String("start", time.Now().Format(time.RFC3339), "")

	err := createFlags.Parse(args[2:])

	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

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
