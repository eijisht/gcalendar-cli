package cmd

import (
	"fmt"
	"time"

	//	"gcal-cli/internal"

	"google.golang.org/api/calendar/v3"
)

// TODO: Filter by calendar
// TODO: Export events

func Read(srv *calendar.Service, calendar string, maxResults int64, maxDays int64) error {
	events, err := requestHandler(srv, calendar, maxResults, maxDays)
	if err != nil {
		return fmt.Errorf("error reading from calendar: %s\n", err)
	}

	for _, item := range events.Items {
		date := item.Start.DateTime
		if date == "" {
			date = item.Start.Date
		}

		fmt.Printf("%s  %s  %s\n", item.Id, date, item.Summary)
	}

	return nil
}

func requestHandler(srv *calendar.Service, calendarName string, maxResults int64, maxDays int64) (calendar.Events, error) {
	var err error
	var events *calendar.Events

	_, err = srv.Calendars.Get(calendarName).Do()
	// could cache the user events and update periodically to reduce api calls and error check

	if err != nil {
		return calendar.Events{}, fmt.Errorf("could not fetch calendar: %s\n", err)
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
		return *events, fmt.Errorf("could not retrieve events: %s\n", err)
	}

	return *events, err
}
