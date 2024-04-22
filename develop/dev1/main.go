package main

import (
	"fmt"
	"log"
	"time"
	"github.com/beevik/ntp"
)
func Time() (time.Time, error){
	return ntp.Time("0.beevik-ntp.pool.ntp.org")
}

func main(){
	time, err := Time()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(time.Format("15:04:05"))
}