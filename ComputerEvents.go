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
	t8        bool   `json:"t8"`
	t9        bool   `json:"t9"`
	t10       bool   `json:"t10"`
	t11       bool   `json:"t11"`
	t12       bool   `json:"t12"`
	t13       bool   `json:"t13"`
	t14       bool   `json:"t14"`
	t15       bool   `json:"t15"`
	t16       bool   `json:"t16"`
	t17       bool   `json:"t17"`
	t18       bool   `json:"t18"`
	t19       bool   `json:"t19"`
	t20       bool   `json:"t20"`
	t21       bool   `json:"t21"`
	t22       bool   `json:"t22"`
}

func (env *Env) CheckEventExists(response http.ResponseWriter, event *ComputerEvent) {
	db := env.db
	if err := db.Where("CompIdent = ? AND Date = ?", event.CompIdent, event.Date).First(&event).Error; err != nil {
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

	if event.t8 {
		Availabilities = append(Availabilities, "8 a.m.")
	}
	if event.t9 {
		Availabilities = append(Availabilities, "9 a.m.")
	}
	if event.t10 {
		Availabilities = append(Availabilities, "10 a.m.")
	}
	if event.t11 {
		Availabilities = append(Availabilities, "11 a.m.")
	}
	if event.t12 {
		Availabilities = append(Availabilities, "12 p.m.")
	}
	if event.t13 {
		Availabilities = append(Availabilities, "1 p.m.")
	}
	if event.t14 {
		Availabilities = append(Availabilities, "2 p.m.")
	}
	if event.t15 {
		Availabilities = append(Availabilities, "3 p.m.")
	}
	if event.t16 {
		Availabilities = append(Availabilities, "4 p.m.")
	}
	if event.t17 {
		Availabilities = append(Availabilities, "5 p.m.")
	}
	if event.t18 {
		Availabilities = append(Availabilities, "6 p.m.")
	}
	if event.t19 {
		Availabilities = append(Availabilities, "7 p.m.")
	}
	if event.t20 {
		Availabilities = append(Availabilities, "8 p.m.")
	}
	if event.t21 {
		Availabilities = append(Availabilities, "9 p.m.")
	}
	if event.t22 {
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

	if err := db.Where("CompIdent = ? AND Date = ?", event.CompIdent, event.Date).First(&hold).Error; err != nil {
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

	dbEvent.t8 = event.t8
	dbEvent.t9 = event.t9
	dbEvent.t10 = event.t10
	dbEvent.t11 = event.t11
	dbEvent.t12 = event.t12
	dbEvent.t13 = event.t13
	dbEvent.t14 = event.t14
	dbEvent.t15 = event.t15
	dbEvent.t16 = event.t16
	dbEvent.t17 = event.t17
	dbEvent.t18 = event.t18
	dbEvent.t19 = event.t19
	dbEvent.t20 = event.t20
	dbEvent.t21 = event.t21
	dbEvent.t22 = event.t22

	db.Save(&dbEvent)
	response.Write([]byte(`{Successful}`))
	fmt.Println("EVENT UPDATED")
}
