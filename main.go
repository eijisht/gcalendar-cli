package main

import (
	"fmt"
	"log"
	//	"flag"

	"gcal-cli/cmd"
	"gcal-cli/internal"
	"github.com/joho/godotenv"
)

// TODO: add caching for API requests (could use SQLite)
// TODO: figure out argument and flag parsing

func main() {
	fmt.Println("This is a CLI interface for the Google Calendar API written in Go!")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	srv, err := internal.GetCalendarService()
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar service: %v", err)
	}

	calendars, err := srv.CalendarList.List().Do()

	if err != nil {
		log.Fatalf("Unable to retrieve Calendar service: %v", err)
	}

	for _, calendar := range calendars.Items {
		println(calendar.Summary)
	}

	cmd.Read(srv, cmd.DefaultCalendar, cmd.DefaultCount)
}
