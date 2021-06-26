package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}
type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// fill response pointer
	log.Println("Getting data")

	*reply = time.Now().Unix()

	return nil
}


func main() {
	time_server := new(TimeServer)

	rpc.Register(time_server)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")

	defer func() {
		log.Println("closing connections...")
		l.Close()
	}()

	if e != nil {
		log.Fatal("listen error: ", e)
	}

	http.Serve(l, nil)
}