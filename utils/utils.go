package utils

import (
	"math/rand"
	"net"
	"time"
)

// reads data from a connection
func ReadData(c net.Conn) (error, string, []byte) {
	var eventName = ""
	var message []byte

	mode := false
	for {
		tempS := make([]byte, 1)
		_, err := c.Read(tempS)
		if err != nil {
			return err, "", nil
		}
		if tempS[0] == byte('\n') && mode == false {
			mode = true
		} else if mode == true {
			if tempS[0] == byte('\n') {
				break
			} else {
				message = append(message, tempS...)
			}
		} else if mode == false {
			eventName += string(tempS)
		}
	}
	return nil, eventName, message
}

// generates a random id
func RandomID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}
