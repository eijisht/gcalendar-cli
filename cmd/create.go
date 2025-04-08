package cmd

import (
	"time"

	"google.golang.org/api/calendar/v3"
)

// TODO: Figure out how to handle parsing time from parameters

func Create(srv *calendar.Service, calendarID string, summary string, descritpion string, endDate string, startDate string) {
	startEventDateTime := parseDate(startDate)
	endEventDateTime := parseDate(endDate)
	event := &calendar.Event{
		Summary:     summary,
		Description: descritpion,
		Start:       startEventDateTime,
		End:         endEventDateTime,
	}
	srv.Events.Insert(calendarID, event)
}

// TODO:
func parseDate(date string) *calendar.EventDateTime {
	time.Now()
	return &calendar.EventDateTime{}
}
