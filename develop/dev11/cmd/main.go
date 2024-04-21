package main

import (
	"L2/develop/dev11/internal/config"
	"L2/develop/dev11/internal/database"
	"L2/develop/dev11/internal/server"
	"L2/develop/dev11/internal/services"
	"log"
)



func main(){
	v, err := config.LoadConfig()
	if err != nil{
		log.Fatal(err)
	}
	conf, err := config.ParseConfig(v)
	if err != nil{
		log.Fatal(err)
	}
	db, err := database.NewDatabase(conf.Postgres)
	if err != nil{
		log.Fatal(err)
	}
	service := services.NewService(db)
	server := server.NewServer(conf.Server, service)
	server.StartServer()
}