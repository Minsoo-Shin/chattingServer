package main

import (
	"fmt"
	"net"
	"strings"
)

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(c *client, msg string) {
	for addr, member := range r.members {
		if addr != c.connection.RemoteAddr() {
			member.msg(msg)
		}
	}
}

func (r *room) clientNicknames(c *client) string {
	var nicknames []string
	for _, member := range r.members {
		if c.nickname != member.nickname {
			nicknames = append(nicknames, member.nickname)
		}
	}
	return fmt.Sprintf("welcome, there is %v", strings.Join(nicknames, ", "))
}
