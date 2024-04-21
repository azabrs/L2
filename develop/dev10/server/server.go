package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"

)

func handleConnection(conn net.Conn){
	defer conn.Close()
	for{
		buf, err := bufio.NewReader(conn).ReadBytes('\n')
		fmt.Println(string(buf))
		if err == io.EOF{
			break
		} else if err != nil{
			log.Fatal(err)
		}
		nn, err := conn.Write(append([]byte("Your messages was "),  buf...))
		fmt.Printf("Was written")
		if nn == 0 || err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Server connection was closed\n")
}

func main(){
	fmt.Println("Server listening")
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil{
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			conn.Close()
			log.Fatal(err)
		}

		fmt.Printf("Connection with %s established\n", conn.RemoteAddr())
		go handleConnection(conn)
	}
}