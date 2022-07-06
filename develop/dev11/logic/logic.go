package logic

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Event struct {
	EventId int    `json:"event_id"`
	Date    string `json:"date"`
	Event   string `json:"event"`
}

//Create event logic
func CreateEvent(event *Event, userId string, storage map[string][]*Event) (string, error) {
	//Validate date format
	_, err := time.Parse("2006-01-02", event.Date)
	if err != nil {
		return "", err
	}

	//Set Event ID
	event.EventId = len(storage[userId]) + 1
	//Add event to memory storage
	storage[userId] = append(storage[userId], event)

	return fmt.Sprintf("%+v", *event), nil
}

//Update event logic
func UpdateEvent(event *Event, userId string, storage map[string][]*Event) (string, error) {
	//Validate date format
	_, err := time.Parse("2006-01-02", event.Date)
	if err != nil {
		return "", err
	}

	//Find event in memory and set new parameters
	var found bool
	for _, v := range storage[userId] {
		if v.EventId == event.EventId {
			v.Date = event.Date
			v.Event = event.Event
			found = true
		}
	}

	if !found {
		return "", fmt.Errorf("event %d not found", event.EventId)
	}

	return fmt.Sprintf("%+v", *event), nil
}

//Delete event
func DeleteEvent(event *Event, userId string, storage map[string][]*Event) (string, error) {
	//Find event in memory by event ID and delete it
	var found bool
	for i, v := range storage[userId] {
		if v.EventId == event.EventId {
			storage[userId] = append(storage[userId][:i], storage[userId][i+1:]...)
			found = true
		}
	}

	if !found {
		return "", fmt.Errorf("event %d not found", event.EventId)
	}
	return fmt.Sprintf("%+v", *event), nil
}

//Events for day
func EventsForDay(userId string, date string, storage map[string][]*Event) (string, error) {
	//Validate date format
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}

	//Search for all events on the required date
	var eventsForDay string
	var found bool
	for _, v := range storage[userId] {
		if v.Date == date {
			eventsForDay += fmt.Sprintf("%+v", *v)
			found = true
		}
	}

	if !found {
		return "", fmt.Errorf("no events on this date: %s", date)
	}

	return fmt.Sprintf("%+v", eventsForDay), nil

}

//Events on the specific week
func EventsForWeek(userId string, week string, storage map[string][]*Event) (string, error) {
	var eventsForWeek string
	var weekNum int
	var found bool

	for _, v := range storage[userId] {
		//Split date and find the day
		nums := strings.Split(v.Date, "-")
		d, err := strconv.Atoi(nums[2])
		if err != nil {
			return "", fmt.Errorf("invalid day: %s", nums[2])
		}
		//Find what week does event take place on
		if d >= 28 {
			weekNum = 4
		} else {
			weekNum = d/7 + 1
		}
		//If the event week is on the specified week add to week events
		if week == strconv.Itoa(weekNum) {
			eventsForWeek += fmt.Sprintf("%+v", *v)
			found = true
		}

	}

	if !found {
		return "", fmt.Errorf("no events this week: %s", week)
	}

	return fmt.Sprintf("%+v", eventsForWeek), nil

}

//Events for month
func EventsForMonth(userId string, month string, storage map[string][]*Event) (string, error) {
	var eventsForMonth string
	var found bool

	//Split date, find month and all the event that month
	for _, v := range storage[userId] {
		nums := strings.Split(v.Date, "-")

		if month == nums[1] {
			eventsForMonth += fmt.Sprintf("%+v", *v)
			found = true
		}

	}

	if !found {
		return "", fmt.Errorf("no events this month: %s", month)
	}

	return fmt.Sprintf("%+v", eventsForMonth), nil

}
