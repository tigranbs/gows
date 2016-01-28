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
