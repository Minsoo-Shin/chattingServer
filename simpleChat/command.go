package main

type commandID int

const (
	CMD_NICKNAME commandID = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_SEND
	CMD_QUIT
)

type command struct {
	id     commandID // action define
	client *client   // Who
	args   []string  // What
}
