package main

import (
	"fmt"
	"io"
	"net/http"
)

type HandlerController struct {
	dbInstance *Database
}

func NewHandlerController(db *Database) *HandlerController {
	return &HandlerController{
		dbInstance: db,
	}
}

func (h *HandlerController) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.PostFormValue("user_id")
	eventDate := r.PostFormValue("event_date")
	eventName := r.PostFormValue("event_name")

	if userID == "" || eventDate == "" || eventName == "" {
		http.Error(w, `{"error": "empty user_id, event_date or event_name"}`, http.StatusBadRequest)
		return
	}

	h.dbInstance.Set(userID, Event{Date: eventDate, Name: eventName})

	io.WriteString(w, fmt.Sprintf("Successfully created event %s %s", eventDate, eventName))
}

func (h *HandlerController) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.PostFormValue("user_id")
	eventDate := r.PostFormValue("event_date")
	eventName := r.PostFormValue("event_name")

	if userID == "" || eventDate == "" || eventName == "" {
		http.Error(w, `{"error": "empty user_id, event_date or event_name"}`, http.StatusBadRequest)
		return
	}

	h.dbInstance.Set(userID, Event{Date: eventDate, Name: eventName})

	io.WriteString(w, fmt.Sprintf("Successfully updated event %s %s", eventDate, eventName))
}

func (h *HandlerController) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.PostFormValue("user_id")
	eventDate := r.PostFormValue("event_date")

	if userID == "" || eventDate == "" {
		http.Error(w, `{"error": "empty user_id, event_date or event_name"}`, http.StatusBadRequest)
		return
	}

	h.dbInstance.Del(userID, eventDate)

	io.WriteString(w, fmt.Sprintf("Successfully deleted event at %s", eventDate))
}

func (h *HandlerController) getDayEvents(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlerController) getWeekEvents(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlerController) getMonthEvents(w http.ResponseWriter, r *http.Request) {

}
