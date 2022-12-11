package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type client struct {
	connection net.Conn
	nickname   string
	room       *room
	commands   chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.connection).ReadString('\n')
		if err != nil {
			log.Printf("failed to read:  %v\n", err.Error())
			return
		}
		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")

		if len(args) <= 1 {
			return
		}

		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/nickname":
			c.commands <- command{
				id:     CMD_NICKNAME,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
				args:   args,
			}
		case "/msg":
			c.commands <- command{
				id:     CMD_SEND,
				client: c,
				args:   args,
			}
		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args,
			}
		default:
			c.err(fmt.Errorf("unkown cmd: %v", cmd))
		}
	}
}

func (c *client) err(err error) {
	c.connection.Write([]byte("ERR: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.connection.Write([]byte("> " + msg + "\n"))
}
