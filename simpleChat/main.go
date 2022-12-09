package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("%v", err.Error())
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("%v", err.Error())
			continue
		}
		go s.newClient(conn)
	}

}
