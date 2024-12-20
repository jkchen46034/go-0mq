package main

import (
	"fmt"
	"log"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	fmt.Println("Connecting to the server...\n")
	s, _ := zctx.NewSocket(zmq.REQ)
	s.Connect("tcp://localhost:5555")

	for i := 0; i < 10; i++ {
		log.Printf("Sending request %d...\n", i)
		s.Send("hello", 0)
		msg, _ := s.Recv(0)
		log.Printf("Received reply %d [%s]\n", i, msg)
	}
}
