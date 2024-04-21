package models

import (

)

type EventTest struct{
	UserID string `json:"user_id" binding "required"`
	Title string `json:"title" binding "required"`
	Descr string `json:"descr" binding "required"`
	Date string `json:"date" binding "required"`
}

type EventCreate struct{
	Date string `json:"date" binding "required"`
	DataEvent EventData `json:"data_event" binding "required"`
	UserID string `json:"user_id" binding "required"`
}

type EventChange struct{
	Date string `json:"date" binding "required"`
	DataEvent EventData
	UserID string `json:"user_id" binding "required"`
	EventID string `json:"event_id" binding "required"`
}


type EventDayRequest struct{
	UserID int
	Date Date
}

type Date struct{
	Year int
	Month string
	Day int
}

type EventData struct{
	Title string `json:"title" binding "required"`
	Description string `json:"description"`
}
