package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("failed to accept: %v", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	log.Printf("TCP Session Open")
	b := make([]byte, 4096)

	for {
		// Read from TCP Buffer
		n, err := conn.Read(b)
		if err != nil {
			fmt.Println("Failed to receive data : ", err)
			break
		}

		if n > 0 {
			//server reply 'server got'
			fmt.Println(string(b[:n]))
			conn.Write(append([]byte("server got: "), b[:n]...))
		}
	}
}
