package main

import (
	"encoding/json"
	"os"
	"sort"
	"sync"
)

type Event struct {
	Date string
	Name string
}

type Data struct {
	Events map[string]map[string]string `json:"users"`
}

type Database struct {
	sync.Mutex
	data     Data
	filePath string
}

func NewDatabase(filePath string) (*Database, error) {
	db := &Database{
		data:     Data{},
		filePath: filePath,
	}

	err := db.loadFromFile()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *Database) loadFromFile() error {
	data, err := os.ReadFile(db.filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &db.data)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) saveToFile() error {
	data, err := json.MarshalIndent(db.data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(db.filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetAllEvents(userID string) []Event {
	db.Lock()

	events := db.data.Events[userID]

	eventsSlice := make([]Event, 0, len(events))

	for key, val := range events {
		eventsSlice = append(eventsSlice, Event{Date: key, Name: val})
	}

	db.Unlock()

	sort.Slice(eventsSlice, func(i, j int) bool {
		return eventsSlice[i].Date < eventsSlice[j].Date
	})

	return eventsSlice
}

func (db *Database) Get(userID string, date string) (*Event, bool) {
	db.Lock()
	defer db.Unlock()

	eventName, ok := db.data.Events[userID][date]

	event := &Event{
		Date: date,
		Name: eventName,
	}

	return event, ok
}

func (db *Database) Set(userID string, event Event) {
	db.Lock()
	defer db.Unlock()

	db.data.Events[userID][event.Date] = event.Name
	db.saveToFile()
}

func (db *Database) Del(userID string, date string) {
	db.Lock()
	defer db.Unlock()

	delete(db.data.Events, date)
}
