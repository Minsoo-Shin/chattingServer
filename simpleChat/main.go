package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("failed to accept: %v", err.Error())
			continue
		}
		go s.newClient(conn)
	}

}
