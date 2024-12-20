package main

import (
	"log"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	zctx, _ := zmq.NewContext()

	s, _ := zctx.NewSocket(zmq.REP)

	s.Bind("tcp://*:5555")

	log.Println("listening on port 5555")

	for {
		msg, _ := s.Recv(0)
		log.Printf("Received %s\n", msg)

		time.Sleep(time.Second * 1)

		s.Send("world", 0)
	}
}
