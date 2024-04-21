package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)


type ConfigParam struct{
	addr string
	timeout *time.Duration
}

func ReadFromSocket(errChan chan error, conn net.Conn){
	for{
		text, err :=  bufio.NewReader(conn).ReadBytes('\n')
		fmt.Println(string(text))
		if err != nil {
			errChan <- fmt.Errorf("remoute server stopped: %v", err)
			return
		}
	}
}

func WriteToSocket(errChan chan error, conn net.Conn){
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadBytes('\n')
		if err != nil{
			errChan <- err
			return
		}
		_, err = conn.Write(text)
		if err != nil{
			errChan <- err
			return
		}
	}
}

func ParseCommandLine() (ConfigParam, error){
	var conf ConfigParam
	conf.timeout = flag.Duration("timeout", 10*time.Second, "timeout for connection")
	
	flag.Parse()
	if len(os.Args) < 3{
		return ConfigParam{}, fmt.Errorf("usage: go-telnet [--timeout=<timeout>] host port")
	} 

	conf.addr = flag.Arg(0) + ":" + flag.Arg(1)

	return conf, nil
}

func Client(conf ConfigParam) error{
	conn, err := net.DialTimeout("tcp", conf.addr, *conf.timeout)
	if err != nil{
		return err
	}
	sigs := make(chan os.Signal, 1)
	errChan := make(chan error)
	signal.Notify(sigs, syscall.SIGINT)
	go ReadFromSocket(errChan, conn)
	go WriteToSocket(errChan, conn)
	select{
	case e := <- errChan:
		fmt.Println("Connection stopped by", e)
	case s := <- sigs:
		fmt.Println("\nConnection stopped by signal:", s)
	}
	return nil
}

func main(){
	conf, err := ParseCommandLine()
	if err != nil{
		log.Fatal(err)
	}
	err = Client(conf)
	if err != nil{
		log.Fatal(err)
	}
}