package main

import (
	"WBL2/develop/dev11/logic"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//Init memory storage
var EventStore = make(map[string][]*logic.Event)

func main() {
	fmt.Println("The Server runs with http://localhost:3000/")

	//Handlers for endpoints with middleware
	createEventFunc := http.HandlerFunc(createEvent)
	http.Handle("/create_event", Logger(createEventFunc))

	updateEventFunc := http.HandlerFunc(updateEvent)
	http.Handle("/update_event", Logger(updateEventFunc))

	deleteEventFunc := http.HandlerFunc(deleteEvent)
	http.Handle("/delete_event", Logger(deleteEventFunc))

	eventsForDayFunc := http.HandlerFunc(eventsForDay)
	http.Handle("/events_for_day", Logger(eventsForDayFunc))

	eventsForWeekFunc := http.HandlerFunc(eventsForWeek)
	http.Handle("/events_for_week", Logger(eventsForWeekFunc))

	eventsForMonthFunc := http.HandlerFunc(eventsForMonth)
	http.Handle("/events_for_month", Logger(eventsForMonthFunc))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("HTTP 500 Server Error: %s ", err)
	}

}

//Logging request middleware
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.String())
		log.Println(r.Method)
		log.Println(r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}

//Request result json
type RequestResult struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

func NewRequestResult(result string, err string) *RequestResult {
	return &RequestResult{
		Result: result,
		Error:  err,
	}
}

//Writing response in json
func WriteResponce(rr *RequestResult, w http.ResponseWriter) {
	j, _ := json.Marshal(rr)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	//Checking content type
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	//Parsing form with user_id, date and event details
	err := r.ParseForm()
	if err != nil {
		rr := NewRequestResult("nil", "HTTP 400 "+err.Error())
		WriteResponce(rr, w)
		return
	}

	event := &logic.Event{
		Date:  r.FormValue("date"),
		Event: r.FormValue("event"),
	}

	//Creating event
	s, err := logic.CreateEvent(event, r.FormValue("user_id"), EventStore)
	if err != nil {
		rr := NewRequestResult("nil", "HTTP 503 "+err.Error())
		WriteResponce(rr, w)
		return
	}

	//Writing response
	rr := NewRequestResult("User "+r.FormValue("user_id")+", created event: "+s, "nil")
	WriteResponce(rr, w)

}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	//Checking content type
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	//Parsing form with user_id, event_id, new date and new event details
	err := r.ParseForm()
	if err != nil {
		rr := NewRequestResult("nil", "HTTP 400 "+err.Error())
		WriteResponce(rr, w)
		return
	}

	eId, err := strconv.Atoi(r.FormValue("event_id"))
	if err != nil {
		rr := NewRequestResult("nil", "HTTP 400 - invalid event id "+err.Error())
		WriteResponce(rr, w)
		return
	}

	event := &logic.Event{
		EventId: eId,
		Date:    r.FormValue("date"),
		Event:   r.FormValue("event"),
	}

	//Update event
	s, err := logic.UpdateEvent(event, r.FormValue("user_id"), EventStore)
	if err != nil {
		rr := NewRequestResult("nil", "HTTP 503 "+err.Error())
		WriteResponce(rr, w)
		return
	}

	//Write response
	rr := NewRequestResult("User "+r.FormValue("user_id")+", updated event: "+s, "nil")
	WriteResponce(rr, w)

}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	//Checking content type
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	//Parsing form with user_id and event ID
	err := r.ParseForm()
	if err != nil {
		rr := NewRequestResult("nil", "HTTP 400 "+err.Error())
		WriteResponce(rr, w)
		return
	}

	eId, err := strconv.Atoi(r.FormValue("event_id"))
	if err != nil {
		rr := NewRequestResult("nil", "HTTP 400 - invalid event id "+err.Error())
		WriteResponce(rr, w)
		return
	}
	event := &logic.Event{
		EventId: eId,
	}

	//Delete event
	s, err := logic.DeleteEvent(event, r.FormValue("user_id"), EventStore)
	if err != nil {
		rr := NewRequestResult("nil", "HTTP 503 "+err.Error())
		WriteResponce(rr, w)
		return
	}

	//write response
	rr := NewRequestResult("User "+r.FormValue("user_id")+", deleted event: "+s, "nil")
	WriteResponce(rr, w)
}

func eventsForDay(w http.ResponseWriter, r *http.Request) {
	//Get query parameters
	userId := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	//Find event for certain date
	s, err := logic.EventsForDay(userId, date, EventStore)
	if err != nil {
		rr := NewRequestResult("nil", "HTTP 503 "+err.Error())
		WriteResponce(rr, w)
		return
	}

	//Write response
	rr := NewRequestResult("User "+userId+", events for "+date+" are "+s, "nil")
	WriteResponce(rr, w)
}

func eventsForWeek(w http.ResponseWriter, r *http.Request) {
	//Get query parameters
	userId := r.URL.Query().Get("user_id")
	week := r.URL.Query().Get("week")

	//Find event for certain week
	s, err := logic.EventsForWeek(userId, week, EventStore)
	if err != nil {
		rr := NewRequestResult("nil", "HTTP 503 "+err.Error())
		WriteResponce(rr, w)
		return
	}
	//Write response
	rr := NewRequestResult("User "+userId+", events for  week "+week+" are "+s, "nil")
	WriteResponce(rr, w)

}

func eventsForMonth(w http.ResponseWriter, r *http.Request) {
	//Get query parameters
	userId := r.URL.Query().Get("user_id")
	month := r.URL.Query().Get("month")

	//Find event for certain month
	s, err := logic.EventsForMonth(userId, month, EventStore)
	if err != nil {
		rr := NewRequestResult("nil", "HTTP 503 "+err.Error())
		WriteResponce(rr, w)
		return
	}
	//Write response
	rr := NewRequestResult("User "+userId+", events for  month "+month+" are "+s, "nil")
	WriteResponce(rr, w)

}
