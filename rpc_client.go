package main

import (
	"log"
	"net/rpc"
)

type Args struct{}

func main() {
	var reply int64

	args := Args{}

	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")

	defer func() {
		log.Println("closing client...")
		client.Close()
	}()

	if err != nil {
		log.Fatal("dialing: ", err)
	}

	err = client.Call("TimeServer.GiveServerTime", args, &reply)
	
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	log.Printf("%d\n", reply)
}