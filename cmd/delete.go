package cmd

import (
	"fmt"

	"google.golang.org/api/calendar/v3"
)

func Delete(srv *calendar.Service, calendarID, eventID string) error {
	if err := srv.Events.Delete(calendarID, eventID).Do(); err != nil {
		return fmt.Errorf("deleting event %s: %w", eventID, err)
	}

	fmt.Printf("Deleted event %s\n", eventID)
	return nil
}
