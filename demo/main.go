package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Printf("error listening: %s", err)
		return
	}
	defer listener.Close()

	fmt.Printf("listening on :1234\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting connection: %s", err)
			continue
		}
		go echoHandler(conn)
	}
}

func echoHandler(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("error reading: %s", err)
			return
		}

		if n == 0 {
			log.Printf("client disconnected")
			return
		}

		_, err = conn.Write(buffer[:n])
		if err != nil {
			log.Printf("error writing: %s", err)
			return
		}
	}
}
