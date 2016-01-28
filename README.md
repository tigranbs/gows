# About
This project is simple and light but very powerful implementation of Browser side Websockets with Go server. It is used 
in <a href="https://treescale.com">TreeScale.com</a> as a part of dashboard live Docker container builds and manage.

# Server Example
Take a look on this simple implementation. It is super easy to get started and make awesome evented things almost with the same 
simple API as <code>Socket.io</code> have.
```go
package main

import (
	"fmt"
	"time"
	"treebuild/tree_socket"
)

func main() {
	tree_socket.On("connection", func(i_sock interface{}) {
		sock := i_sock.(*tree_socket.Socket)
		sock.On("api_init", func(data interface{}) {
			fmt.Println(data)
			i := 0
			for {
				sock.Emit("test_emit", map[string]int{"net_data": i})
				i++
				time.Sleep(time.Second * 1)
			}
		})

		sock.On("disconnect", func(i_sock interface{}) {
			fmt.Println("Disconnected")
		})
	})

	tree_socket.Listen(":3000", "/")
}

```

# Client Side Example
With JavaScript in <a href="https://treescale.com">TreeScale.com</a>  we are using plain Websocket implementation so it will work
on most browsers and you don't need to include 90Kb+ <code>Socket.io</code> library to do basic evented systems.
```javascript
var socket = require("./lib/socket");  // in this case we have socket.js file in lib directory

socket.connect("localhost:3000");
socket.on("test_emit", function (data) {
    console.log(data);
});

socket.on("connection", function (data) {
    console.log("connected", data);
    socket.emit("api_init", {test: 15});
});
```

# Is this useful for you ?
Let me know if you want to add Websocket event <code>Rooms, Groups or any other special things</code> for make this light 
thing more functional. I'm ready to help.<br/>
<strong>With :heart: from <a href="https://treescale.com">TreeScale Inc.</a></strong>
