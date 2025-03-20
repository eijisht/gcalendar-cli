package main

import (
	"flag"
	"log"
	"os"

	"gcal-cli/cmd"
	"gcal-cli/internal"

	"github.com/joho/godotenv"
)

// TODO: add caching for API requests (could use SQLite)
// TODO: figure out argument and flag parsing

func main() {
	flags := parseStdin()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	srv, err := internal.GetCalendarService()
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar service: %v", err)
	}

	if flags.Read {
		cmd.Read(srv, *flags.Calendar, *flags.Count, *flags.Days)
	}
}

type arguments struct {
	Read  bool
	Write bool

	Calendar *string
	Count    *int64
	Days     *int64
}

func parseStdin() arguments {
	read := false
	write := false

	if len(os.Args) >= 2 {
		command := os.Args[1]
		read = command == "read"
		write = command == "write"

	}

	calendar := flag.String("r", "primary", "usage: -r <calendar id>")
	count := flag.Int64("c", 10, "usage: -u <int>")
	day := flag.Int64("d", -1, "usage: -d <int>")

	// TODO: Improve usage messages

	flag.Parse()

	return arguments{
		read,
		write,
		calendar,
		count,
		day,
	}
}
