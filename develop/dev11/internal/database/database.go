package database

import (
	"L2/develop/dev11/internal/config"
	"L2/develop/dev11/internal/constants"
	"L2/develop/dev11/internal/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct{
	Db *sql.DB
}

func NewDatabase(conf config.PostgresConfig) (Database, error){
	s := fmt.Sprintf("dbname=%s user=%s password=%s port=%s sslmode = disable", conf.DbName, conf.User, conf.Password, conf.Port)
	pg, err := sql.Open("postgres", s)
	if err != nil{
		return Database{}, err
	}
	return Database{ Db: pg}, nil
}

func(pg Database)EventForDay(param models.EventDayRequest) ([]models.EventData, error){
	var EventsData []models.EventData
	query := `SELECT data_title, data_description FROM events
			WHERE user_id = $1 AND day_number = $2 AND month_name = $3 AND year_number = $4`
	rows, err := pg.Db.Query(query, param.UserID, param.Date.Day, param.Date.Month, param.Date.Year)
	if err != nil{
		return []models.EventData{}, err
	}
	for rows.Next(){
		var EventData models.EventData
		if err := rows.Scan(&EventData.Title, &EventData.Description); err != nil{
			return []models.EventData{}, err
		}
		EventsData = append(EventsData, EventData)
	}
	if len(EventsData) == 0{
		return []models.EventData{}, constants.ErrEventNotFound
	}
	return EventsData, nil
}


func(pg Database)EventForMonth(param models.EventDayRequest) ([]models.EventData, error){
	var EventsData []models.EventData
	query := `SELECT data_title, data_description FROM events
			WHERE user_id = $1 AND month_name = $3 AND year_number = $4`
	rows, err := pg.Db.Query(query, param.UserID, param.Date.Month, param.Date.Year)
	if err != nil{
		return []models.EventData{}, err
	}
	for rows.Next(){
		var EventData models.EventData
		if err := rows.Scan(&EventData.Title, &EventData.Description); err != nil{
			return []models.EventData{}, err
		}
		EventsData = append(EventsData, EventData)
	}
	if len(EventsData) == 0{
		return []models.EventData{}, constants.ErrEventNotFound
	}
	return EventsData, nil
}

func(pg Database)EventForWeek(param models.EventDayRequest, WeekNumber int) ([]models.EventData, error){
	var EventsData []models.EventData
	query := `SELECT data_title, data_description FROM events
			WHERE user_id = $1 AND week_number = $3 AND year_number = $4`
	rows, err := pg.Db.Query(query, param.UserID, WeekNumber, param.Date.Year)
	if err != nil{
		return []models.EventData{}, err
	}
	for rows.Next(){
		var EventData models.EventData
		if err := rows.Scan(&EventData.Title, &EventData.Description); err != nil{
			return []models.EventData{}, err
		}
		EventsData = append(EventsData, EventData)
	}
	if len(EventsData) == 0{
		return []models.EventData{}, constants.ErrEventNotFound
	}
	return EventsData, nil
}

func(pg Database)CreateEvent(NewEvent models.EventCreate, paramDate models.Date, WeekNumber int)error{
	query := `INSERT INTO events (user_id, week_number, day_number, month_name, year_number, data_title, data_description) VALUES 
	($1, $2, $3, $4, $5, $6, $7)`
	_, err := pg.Db.Exec(query, NewEvent.UserID, WeekNumber, paramDate.Day, paramDate.Month, paramDate.Year, NewEvent.DataEvent.Title, NewEvent.DataEvent.Description)
	if err != nil{
		return err
	}
	return nil
}

func(pg Database)UpdateEvent(NewEvent models.EventChange, paramDate models.Date, WeekNumber int)error{
	query := `UPDATE events
			SET user_id = $2 AND week_number = $3, day_number = $4, month_name = $5, year_number = $6, data_title = $7, data_description = $8
			WHERE event_id = $1`
	_, err := pg.Db.Exec(query, NewEvent.EventID, NewEvent.UserID, WeekNumber, paramDate.Day, paramDate.Month, paramDate.Year, NewEvent.DataEvent.Title, NewEvent.DataEvent.Description)
	if err != nil{
		return err
	}
	return nil
}

func(pg Database)DeleteEvent(eventID string)error{
	query := `DELETE FROM events
			WHERE event_id = $1`
	_, err := pg.Db.Exec(query, eventID)
	if err != nil{
		return err
	}
	return nil
}

func(pg Database)IsUserExist(userID string)(bool, error){
	query := `SELECT user_id FROM events
				WHERE user_id = $1
				GROUP BY user_id`
	var buf string
	if err := pg.Db.QueryRow(query, userID).Scan(buf); err == constants.ErrNoRow{
		return false, nil
	}else if err != nil{
		return false, err
	}
	return true, nil
}

func(pg Database)IsUserEventExist(eventID, UserID string)(bool, error){
	query := `SELECT event_id, user_id FROM events
				WHERE event_id = $1 AND user_id = $2
				`
	var buf1, buf2 string
	if err := pg.Db.QueryRow(query, eventID, UserID).Scan(buf1, buf2); err == constants.ErrNoRow{
		return false, nil
	}else if err != nil{
		return false, err
	}
	return true, nil
}
