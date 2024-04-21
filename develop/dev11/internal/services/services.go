package services

import (
	"L2/develop/dev11/internal/constants"
	"L2/develop/dev11/internal/database"
	"L2/develop/dev11/internal/models"
	"strconv"
	"time"
)


type Service struct{
	db database.Database
}

func NewService(db database.Database) Service{
	return Service{ db : db}
}

func(s Service)EventForDay(userID string, date string) ([]models.EventData, error){
	var param models.EventDayRequest
	var err error
	var t time.Time
	exist, err := s.db.IsUserExist(userID)
	if err != nil{
		return []models.EventData{}, err
	}
	if !exist{
		return []models.EventData{}, constants.ErrUserIDNotFound
	}
	param.UserID, err = strconv.Atoi(userID)
	if err != nil{
		return []models.EventData{}, constants.ErrIncorrectInputData
	}
	if param.UserID < 1{
		return []models.EventData{}, constants.ErrIncorrectInputData
	}
	if t, err = s.ParseSringToDate(date); err != nil{
		return []models.EventData{}, err
	}
	param.Date = s.SplitDate(t)
	return s.db.EventForDay(param)
}

func(s Service)EventForMonth(userID string, date string) ([]models.EventData, error){
	var param models.EventDayRequest
	var err error
	var t time.Time
	exist, err := s.db.IsUserExist(userID)
	if err != nil{
		return []models.EventData{}, err
	}
	if !exist{
		return []models.EventData{}, constants.ErrUserIDNotFound
	}
	param.UserID, err = strconv.Atoi(userID)
	if err != nil{
		return []models.EventData{}, constants.ErrIncorrectInputData
	}
	if param.UserID < 1{
		return []models.EventData{}, constants.ErrIncorrectInputData
	}
	if t, err = s.ParseSringToDate(date); err != nil{
		return []models.EventData{}, err
	}
	param.Date = s.SplitDate(t)
	return s.db.EventForMonth(param)
}

func(s Service)EventForWeek(userID string, date string) ([]models.EventData, error){
	var param models.EventDayRequest
	var err error
	var t time.Time
	exist, err := s.db.IsUserExist(userID)
	if err != nil{
		return []models.EventData{}, err
	}
	if !exist{
		return []models.EventData{}, constants.ErrUserIDNotFound
	}
	param.UserID, err = strconv.Atoi(userID)
	if err != nil{
		return []models.EventData{}, constants.ErrIncorrectInputData
	}
	if param.UserID < 1{
		return []models.EventData{}, constants.ErrIncorrectInputData
	}
	if t, err = s.ParseSringToDate(date); err != nil{
		return []models.EventData{}, err
	}
	param.Date = s.SplitDate(t)
	_, WeekNumber := t.ISOWeek()
	return s.db.EventForWeek(param, WeekNumber)
}

func(s Service)CreateEvent(newEvent models.EventCreate) error{
	var t time.Time
	var err error
	var dateParam models.Date
	if t, err = s.ParseSringToDate(newEvent.Date); err != nil{
		return constants.ErrIncorrectInputData
	}
	dateParam = s.SplitDate(t)
	_, WeekNumber := t.ISOWeek()
	if err = s.db.CreateEvent(newEvent, dateParam, WeekNumber); err != nil{
		return err
	}
	return nil
}

func(s Service)UpdateEvent(newEvent models.EventChange) error{
	var t time.Time
	var err error
	var dateParam models.Date
	exist, err := s.db.IsUserEventExist(newEvent.UserID, newEvent.EventID)
	if err != nil{
		return err
	}
	if !exist{
		return constants.ErrUserIDNotFound
	}
	if t, err = s.ParseSringToDate(newEvent.Date); err != nil{
		return constants.ErrIncorrectInputData
	}
	dateParam = s.SplitDate(t)
	_, WeekNumber := t.ISOWeek()
	if err = s.db.UpdateEvent(newEvent, dateParam, WeekNumber); err != nil{
		return err
	}
	return nil
}

func(s Service)DeleteEvent(userID, eventID string) error{
	var err error
	exist, err := s.db.IsUserEventExist(userID, eventID)
	if err != nil{
		return err
	}
	if !exist{
		return constants.ErrUserIDNotFound
	}
	if err = s.db.DeleteEvent(eventID); err != nil{
		return err
	}
	return nil
}

func(s Service)ParseSringToDate(date string) (time.Time, error){
	var (
		eventDate time.Time
		err       error
	)

	eventDate, err = time.Parse("2006-01-02T15:04", date)
	if err != nil {
		eventDate, err = time.Parse("2006-01-02", date)
		if err != nil {
			eventDate, err = time.Parse("2006-01-02T15:04:00Z", date)
			if err != nil {
				return time.Time{}, constants.ErrIncorrectInputData
			}
		}
	}
	

	return eventDate, nil
}

func (s Service)SplitDate(t time.Time) models.Date{
	var	month time.Month
	var NewParam models.Date
	NewParam.Year, month, NewParam.Day = t.Date()
	NewParam.Month = month.String()
	return NewParam
}

