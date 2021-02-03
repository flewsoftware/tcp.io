package client

import (
	"net"
	"tcpio/event"
)

type Socket struct {
	connection net.Conn
	Events     map[string]event.Handler
}

// emits data to event listener's
func (s *Socket) Emit(eventName string, data []byte) {
	eventNameB := append([]byte(eventName), byte('\n'))
	data = append(data, byte('\n'))
	s.connection.Write(append(eventNameB, data...))
}

// sets a event handler
func (s *Socket) On(eventName string, cb event.Handler) {
	s.Events[eventName] = cb
}
