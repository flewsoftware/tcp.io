package client

import (
	"net"
	"sync"
	"tcpio/event"
	"tcpio/events"
	"tcpio/utils"
)

type Client struct {
	Config     Config
	Events     map[string]ConnectionHandler
	connection net.Conn
	mutex      sync.Mutex
}

type Config struct {
	Addr string
}

// connects to the addr defined in Client.Config
func (c *Client) Connect() net.Conn {
	conn, err := net.Dial("tcp", c.Config.Addr)
	if err != nil {
		return conn
	}
	c.connection = conn
	socket := Socket{
		connection: conn,
		Events:     map[string]event.Handler{},
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if val, ok := c.Events[events.Connection]; ok {
		val(socket)
	}

	go func() {
		for {
			err, eventName, message := utils.ReadData(c.connection)
			if err != nil {
				break
			}
			if val, ok := socket.Events[eventName]; ok {
				val(message)
			}
		}
	}()
	return nil
}

type ConnectionHandler func(Socket)

// sets a event handler
func (c *Client) On(eventName string, cb ConnectionHandler) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Events[eventName] = cb
}

// emits data to event listener's
func (c *Client) Emit(eventName string, data []byte) {
	c.connection.Write(append(append([]byte(eventName), []byte("\n")...), data...))
}

// creates a new Client
func Create(config Config) Client {
	return Client{Config: config, Events: make(map[string]ConnectionHandler)}
}
