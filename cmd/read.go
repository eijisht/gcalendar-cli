package cmd

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

// TODO: Filter by calendar
// TODO: Export events

const DefaultCalendar string = "primary"
const DefaultCount int64 = 10
const DefaultMaxDays int64 = -1

func Read(srv *calendar.Service, calendar string, maxResults int64) {
	events, err := requestHandler(srv, calendar, maxResults, DefaultMaxDays)
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

	if maxDays == DefaultMaxDays {
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

	return *events, err
}
