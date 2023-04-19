package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type ComputerEvent struct {
	gorm.Model
	Date      string `json:"date"`
	CompIdent string `json:"compid" gorm:"uniqueIndex"`
	T8        string `json:"t8"`
	T9        string `json:"t9"`
	T10       string `json:"t10"`
	T11       string `json:"t11"`
	T12       string `json:"t12"`
	T13       string `json:"t13"`
	T14       string `json:"t14"`
	T15       string `json:"t15"`
	T16       string `json:"t16"`
	T17       string `json:"t17"`
	T18       string `json:"t18"`
	T19       string `json:"t19"`
	T20       string `json:"t20"`
	T21       string `json:"t21"`
	T22       string `json:"t22"`
}

func (env *Env) CheckEventExists(response http.ResponseWriter, event *ComputerEvent) {
	db := env.db
	if err := db.Where("comp_ident = ? AND date = ?", event.CompIdent, event.Date).First(&event).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Write([]byte("No event on that date exists for that computer: " + err.Error()))
			return
		} else {
			panic("something terrible has happened")
		}
	}
}

func (env *Env) GetEventAvailability(response http.ResponseWriter, request *http.Request) {
	fmt.Println("RETRIEVING EVENT AVAILABILITY")
	response.Header().Set("Content-Type", "application/json")
	var event ComputerEvent = ComputerEvent{}
	json.NewDecoder(request.Body).Decode(&event)

	env.CheckEventExists(response, &event)

	var Availabilities []string

	if event.T8 == "8AM" {
		Availabilities = append(Availabilities, "8 a.m.")
	}
	if event.T9 == "9AM" {
		Availabilities = append(Availabilities, "9 a.m.")
	}
	if event.T10 == "10AM" {
		Availabilities = append(Availabilities, "10 a.m.")
	}
	if event.T11 == "11AM" {
		Availabilities = append(Availabilities, "11 a.m.")
	}
	if event.T12 == "12PM" {
		Availabilities = append(Availabilities, "12 p.m.")
	}
	if event.T13 == "1PM" {
		Availabilities = append(Availabilities, "1 p.m.")
	}
	if event.T14 == "2PM" {
		Availabilities = append(Availabilities, "2 p.m.")
	}
	if event.T15 == "3PM" {
		Availabilities = append(Availabilities, "3 p.m.")
	}
	if event.T16 == "4PM" {
		Availabilities = append(Availabilities, "4 p.m.")
	}
	if event.T17 == "5PM" {
		Availabilities = append(Availabilities, "5 p.m.")
	}
	if event.T18 == "6PM" {
		Availabilities = append(Availabilities, "6 p.m.")
	}
	if event.T19 == "7PM" {
		Availabilities = append(Availabilities, "7 p.m.")
	}
	if event.T20 == "8PM" {
		Availabilities = append(Availabilities, "8 p.m.")
	}
	if event.T21 == "9PM" {
		Availabilities = append(Availabilities, "9 p.m.")
	}
	if event.T22 == "10PM" {
		Availabilities = append(Availabilities, "10 p.m.")
	}

	j, err := json.Marshal(Availabilities)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	response.Write(j)
}

func (env *Env) CreateEvent(response http.ResponseWriter, request *http.Request) {
	fmt.Println("CREATING NEW EVENT")
	response.Header().Set("Content-Type", "application/json")
	var event ComputerEvent = ComputerEvent{}
	var hold ComputerEvent = ComputerEvent{}
	json.NewDecoder(request.Body).Decode(&event)

	db := env.db

	if err := db.Where("comp_ident = ? AND date = ?", event.CompIdent, event.Date).First(&hold).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false).Create(&event)
			response.Write([]byte(`Event created`))
			fmt.Println("EVENT CREATION SUCCESS")
			return
		} else {
			panic("something terrible has happened")
		}
	} else {
		response.Write([]byte(`Event already exists`))
		return
	}
}

func (env *Env) UpdateEvent(response http.ResponseWriter, request *http.Request) {
	fmt.Println("UPDATING EVENT")
	response.Header().Set("Content-Type", "application/json")
	var event ComputerEvent = ComputerEvent{}
	var dbEvent ComputerEvent = ComputerEvent{}
	db := env.db
	json.NewDecoder(request.Body).Decode(&event)

	dbEvent.Date = event.Date
	dbEvent.CompIdent = event.CompIdent
	env.CheckEventExists(response, &dbEvent)

	dbEvent.T8 = event.T8
	dbEvent.T9 = event.T9
	dbEvent.T10 = event.T10
	dbEvent.T11 = event.T11
	dbEvent.T12 = event.T12
	dbEvent.T13 = event.T13
	dbEvent.T14 = event.T14
	dbEvent.T15 = event.T15
	dbEvent.T16 = event.T16
	dbEvent.T17 = event.T17
	dbEvent.T18 = event.T18
	dbEvent.T19 = event.T19
	dbEvent.T20 = event.T20
	dbEvent.T21 = event.T21
	dbEvent.T22 = event.T22

	db.Save(&dbEvent)
	response.Write([]byte(`{Successful}`))
	fmt.Println("EVENT UPDATED")
}
