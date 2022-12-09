package main

import (
	errors "errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room //key:name
	commands chan command
}

func newServer() *server {
	fmt.Println("server start")

	return &server{
		rooms:    make(map[string]*room, 0),
		commands: make(chan command),
	}
}

// 서버가 돈다는 것은 무엇일까? 지속적으로 무언가를
func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICKNAME:
			s.setNickname(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client)
		case CMD_SEND:
			s.send(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *server) setNickname(client *client, args []string) {
	if len(args) <= 1 {
		fmt.Printf("invalid args: %v\n", len(args))
	}
	nickname := args[1]
	client.nickname = nickname
	client.msg(fmt.Sprintf("hi, %v \n", client.nickname))
}

func (s *server) join(c *client, args []string) {
	if len(args) <= 1 {
		fmt.Printf("invalid args: %v\n", len(args))
	}

	roomName := args[1]

	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}

	r.members[c.connection.RemoteAddr()] = c

	s.makeQuitCurrentRoom(c)

	c.room = r

	r.broadcast(c, fmt.Sprintf("%v has joined the room\n", c.nickname))
	c.msg(r.clientNicknames(c))
}

func (s *server) listRooms(client *client) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}
	client.msg(fmt.Sprintf("%v", strings.Join(rooms, ",")))
}

func (s *server) send(client *client, args []string) {
	if len(args) <= 1 {
		fmt.Printf("invalid args: %v", len(args))
	}
	if client.room == nil {
		client.err(errors.New("you mus join the room first"))
	}

	client.room.broadcast(client, fmt.Sprintf("%v", args[1]))
}

func (s *server) quit(client *client) {
	fmt.Printf("%v has quit: %v", client.nickname, client.connection.RemoteAddr().String())

	s.makeQuitCurrentRoom(client)

	client.msg("see you")
	client.connection.Close()
}

func (s *server) makeQuitCurrentRoom(client *client) {
	if client.room != nil {
		delete(client.room.members, client.connection.RemoteAddr())
		client.room.broadcast(client, fmt.Sprintf("%v has quit", client.nickname))
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has connected: %s", conn.RemoteAddr().String())

	c := &client{
		connection: conn,
		nickname:   "anonymous",
		room:       nil,
		commands:   s.commands,
	}

	c.readInput()
}
