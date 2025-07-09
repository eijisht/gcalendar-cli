package cmd

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

// TODO: Figure out how to handle parsing time from parameters

// TODO: move timezone to config file
var TIMEZONE string = time.Local.String()

const TIME_LAYOUT string = "2006-01-02 15:04"

func Create(srv *calendar.Service, calendarID string, summary string, descritpion string, endDate string, startDate string) {
	startEventDateTime, err := ParseDate(startDate)
	if err != nil {
		log.Fatalf("Error parsing date: %s\n", err)
	}

	endEventDateTime, err := ParseDate(endDate)
	if err != nil {
		log.Fatalf("Error parsing date: %s\n", err)
	}

	event := &calendar.Event{
		Summary:     summary,
		Description: descritpion,
		//		Start:       startEventDateTime,
		//		End:         endEventDateTime,
	}

	//	srv.Events.Insert(calendarID, event)
	fmt.Println(startEventDateTime, endEventDateTime)
	fmt.Println(event.Summary, event.Description, event.Start, event.End)
}

// TODO:
func ParseDate(value string) (time.Time, error) {
	if len(value) < 18 {
		value += TIMEZONE
	}
	date, err := time.Parse(TIME_LAYOUT, value)

	if err != nil {
		log.Fatalf("Error parsing date: %s\n", err)
	}

	fmt.Println(date.Zone())
	return date, err

	//	datetime := &calendar.EventDateTime{
	//		Date: date.Format("RFC3339"),
	//	}

	// return datetime, err
}
