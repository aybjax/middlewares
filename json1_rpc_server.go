package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Args struct{
	Id string
}

type Book struct {
	Id string `json:string,omitempty`
	Name string `json:name,omitempty`
	Author string `json:author,omitempty`
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	log.Printf("%#v", args)
	log.Printf("%#v", reply)
	log.Printf("%#v", r)

	var books []Book

	raw, readerr := ioutil.ReadFile("./books.json")

	if readerr != nil {
		log.Println("error: ", readerr)
	}

	marshalerr := jsonparse.Unmarshal(raw, &books)

	if marshalerr != nil {
		log.Println("error:", marshalerr)

		os.Exit(1)
	}

	for _, book := range books {
		if book.Id == args.Id {
			*reply = book
			break
		}
	}

	return nil
}


func main() {
	s := rpc.NewServer()

	s.RegisterCodec(json.NewCodec(), "application/json")

	s.RegisterService(new(JSONServer), "")

	r := mux.NewRouter()

	r.Handle("/rpc", s)

	http.ListenAndServe(":1234", r)
}