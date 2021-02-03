package server

import (
	"net"
	"tcpio/event"
)

type Socket struct {
	connection net.Conn
	Events     map[string]event.Handler
	Id         int
}

func (s *Socket) Emit(eventName string, data []byte) {
	eventNameB := append([]byte(eventName), byte('\n'))
	data = append(data, byte('\n'))
	s.connection.Write(append(eventNameB, data...))
}

func (s *Socket) On(eventName string, cb event.Handler) {
	s.Events[eventName] = cb
}
