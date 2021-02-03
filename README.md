## What is TCP.IO
TCP.IO enables real-time bidirectional event-based communication. Inspired by Socket.IO,

## What can I use it for
* To create messaging apps
* To create real-time apps

## How to use
server:
```go
package main
import (
    "fmt"
	"log"
"tcpio/events"
    "tcpio/server"
)

func main() {
    io := server.Create(server.Config{Addr: ":8000"})
    io.On(events.Connection, func(socket server.Socket) {
    	socket.On("response", func(bytes []byte) {
            fmt.Println(string(bytes))
        })
        socket.Emit("request", []byte("ask"))
    })
    err := io.Listen()
    if err != nil {
        log.Fatalln(err)
    }
}
```

client:
```go
package main
import (
    "fmt"
    "tcpio/client"
    "tcpio/events"
)
    
func main() {
	c := client.Create(client.Config{Addr: ":8000"})
	c.On(events.Connection, func(socket client.Socket) {
		socket.On("request", func(bytes []byte) {
            fmt.Println(string(bytes))
            socket.Emit("response", []byte("give"))
        })
    })
    c.Connect()
}
```