package server

import (
	"L2/develop/dev11/internal/constants"
	"L2/develop/dev11/internal/models"
	"encoding/json"
	"net/http"
)

type ResultResponse struct {
	Result []models.EventData `json:"result"`
}

type ErrorResponse struct {
	Err string `json:"error"`
}

func(s Server)EventsForDay(w http.ResponseWriter, r *http.Request){
	userID := r.URL.Query().Get("user_id")
	Date := r.URL.Query().Get("date")
	events, err := s.service.EventForDay(userID, Date)
	if err == constants.ErrIncorrectInputData{
		w.Header().Set("Content-Type", "application/json")
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusBadRequest)
		return
	} else if err != nil{
		w.Header().Set("Content-Type", "application/json")
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	result, _ := json.MarshalIndent(&ResultResponse{Result: events}, " ", "")
	_, err = w.Write(result)
	if err != nil {
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusInternalServerError)
		return
	}
	
}

func(s Server)EventsFoMonth(w http.ResponseWriter, r *http.Request){
	userID := r.URL.Query().Get("user_id")
	Date := r.URL.Query().Get("date")
	events, err := s.service.EventForMonth(userID, Date)
	if err == constants.ErrIncorrectInputData{
		w.Header().Set("Content-Type", "application/json")
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusBadRequest)
		return
	} else if err != nil{
		w.Header().Set("Content-Type", "application/json")
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	result, _ := json.MarshalIndent(&ResultResponse{Result: events}, " ", "")
	_, err = w.Write(result)
	if err != nil {
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusInternalServerError)
		return
	}
	
}

func(s Server)EventsFoWeek(w http.ResponseWriter, r *http.Request){
	userID := r.URL.Query().Get("user_id")
	Date := r.URL.Query().Get("date")
	events, err := s.service.EventForWeek(userID, Date)
	if err == constants.ErrIncorrectInputData{
		w.Header().Set("Content-Type", "application/json")
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusBadRequest)
		return
	} else if err != nil{
		w.Header().Set("Content-Type", "application/json")
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	result, _ := json.MarshalIndent(&ResultResponse{Result: events}, " ", "")
	_, err = w.Write(result)
	if err != nil {
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusInternalServerError)
		return
	}
	
}

func(s Server)CreateEvent(w http.ResponseWriter, r *http.Request){
	var newEvent models.EventCreate
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&newEvent); err != nil{
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusBadRequest)
		return
	}

	if err := s.service.CreateEvent(newEvent); err == constants.ErrIncorrectInputData{
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusBadRequest)
		return
	}else if err != nil{
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func(s Server)UpdateEvent(w http.ResponseWriter, r *http.Request){
	var newEvent models.EventChange
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&newEvent); err != nil{
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusBadRequest)
		return
	}
	if err := s.service.UpdateEvent(newEvent); err == constants.ErrIncorrectInputData{
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusBadRequest)
		return
	}else if err != nil{
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
}


func(s Server)DeleteEvent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var newEvent models.EventChange
	if err := json.NewDecoder(r.Body).Decode(&newEvent); err != nil{
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusBadRequest)
		return
	}
	if err := s.service.DeleteEvent(newEvent.UserID, newEvent.EventID); err == constants.ErrIncorrectInputData{
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusBadRequest)
		return
	}else if err != nil{
		jsonErr, _ := json.MarshalIndent(&ErrorResponse{Err: err.Error()}, " ", " ")
		http.Error(w, string(jsonErr), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}