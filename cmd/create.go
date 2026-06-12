package cmd

import (
	"fmt"
	"time"

	"google.golang.org/api/calendar/v3"
)

const timeLayout string = "2006-01-02 15:04"

func Create(srv *calendar.Service, calendarID, summary, description, startStr, endStr string) error {
	start, err := ParseDate(startStr)
	if err != nil {
		return fmt.Errorf("parsing start date: %w", err)
	}

	end, err := ParseDate(endStr)
	if err != nil {
		return fmt.Errorf("parsing end date: %w", err)
	}

	event := &calendar.Event{
		Summary:     summary,
		Description: description,
		Start:       &calendar.EventDateTime{DateTime: start.Format(time.RFC3339)},
		End:         &calendar.EventDateTime{DateTime: end.Format(time.RFC3339)},
	}

	created, err := srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		return fmt.Errorf("creating event: %w", err)
	}

	fmt.Printf("Created event: %s\n", created.HtmlLink)
	return nil
}

// ParseDate interprets a "2006-01-02 15:04" string as local wall-clock time.
func ParseDate(value string) (time.Time, error) {
	return time.ParseInLocation(timeLayout, value, time.Local)
}
