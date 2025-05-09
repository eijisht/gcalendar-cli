package cmd

import (
	"fmt"
	"log"
	"time"

	//	"gcal-cli/internal"

	"google.golang.org/api/calendar/v3"
)

// TODO: Filter by calendar
// TODO: Export events

func Read(srv *calendar.Service, calendar string, maxResults int64, maxDays int64) {
	events, err := requestHandler(srv, calendar, maxResults, maxDays)
	if err != nil {
		log.Fatalf("Unable to retrieve events: %v", err)
	}

	for _, item := range events.Items {
		date := item.Start.DateTime
		if date == "" {
			date = item.Start.Date
		}

		fmt.Printf("%s: %s\n", date, item.Summary)
	}
}

func requestHandler(srv *calendar.Service, calendarName string, maxResults int64, maxDays int64) (calendar.Events, error) {
	var err error
	var events *calendar.Events

	_, err = srv.Calendars.Get(calendarName).Do()
	// could cache the user events and update periodically to reduce api calls and error check

	if err != nil {
		log.Fatalf("Unable to get events: Calendar does not exist\n")
	}

	if maxDays == -1 {
		events, err = srv.Events.List(calendarName).
			TimeMin(time.Now().Format(time.RFC3339)).
			MaxResults(maxResults).
			SingleEvents(true).
			OrderBy("startTime").
			Do()

	} else {
		maxTime := time.Now().AddDate(0, 0, int(maxDays))
		events, err = srv.Events.List(calendarName).
			TimeMin(time.Now().Format(time.RFC3339)).
			TimeMax(maxTime.Format(time.RFC3339)).
			MaxResults(maxResults).
			SingleEvents(true).
			OrderBy("startTime").
			Do()

	}

	if err != nil {
		log.Fatalf("Could not retrieve events: %s\n", err)
	}

	return *events, err
}
