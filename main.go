package main

import (
	"flag"
	"fmt"
	"log"

	"gcal-cli/cmd"
	"gcal-cli/internal"

	"github.com/joho/godotenv"
)

// TODO: add caching for API requests (could use SQLite)
// TODO: figure out argument and flag parsing

func main() {
	flags := initFlags()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	srv, err := internal.GetCalendarService()
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar service: %v", err)
	}

	cmd.Read(srv, *flags.Read, *flags.Count, *flags.Days)
}

type arguments struct {
	Read  *string
	Count *int64
	Days  *int64
}

func initFlags() arguments {
	readFlag := flag.String("r", "primary", "usage: -r <calendar id>")
	countFlag := flag.Int64("c", 10, "usage: -u <int>")
	dayFlag := flag.Int64("d", -1, "usage: -d <int>")

	// TODO: Improve usage messages

	flag.Parse()

	return arguments{
		readFlag,
		countFlag,
		dayFlag,
	}
}
