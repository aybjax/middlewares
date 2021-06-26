package main

import (
	"fmt"
	"log"
	"net/http"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request")

		handler.ServeHTTP(w, r)

		fmt.Println("Executing middleware after response")
	})
}


func mainLogic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing main handler")

	w.Write([]byte("OK"))
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)

	http.Handle("/", middleware(mainLogicHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}