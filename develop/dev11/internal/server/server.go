package server

import (
	"L2/develop/dev11/internal/config"
	"L2/develop/dev11/internal/services"
	"log"
	"net/http"
)


type Server struct{
	Addr string
	service services.Service
}

func NewServer(conf config.ServerConfig, service services.Service) Server{
	return Server{
				Addr: conf.Host + ":" + conf.Port,
				service: service,
				}
}

func (s Server)RegisterHandler(mux *http.ServeMux){
	mux.HandleFunc("/events_for_day", s.EventsForDay)
	mux.HandleFunc("/events_for_month", s.EventsFoMonth)
	mux.HandleFunc("/events_for_week", s.EventsFoWeek)
	mux.HandleFunc("/update_event", s.UpdateEvent)
	mux.HandleFunc("/create_event", s.CreateEvent)
	mux.HandleFunc("/delete_event", s.DeleteEvent)
}

func (s Server)StartServer(){
	mux := http.NewServeMux()
	s.RegisterHandler(mux)
	err := http.ListenAndServe(s.Addr, mux)
	if err != nil{
		log.Fatal(err)
	}

}