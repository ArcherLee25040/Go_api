package main

import (
	"fmt"
	"net/http"

	"github.com/google/martian/log"
)

func main() {
	log.SetResponsetCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Start Sever at now")

	fmt.Println("GO API")

	error := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(error)
	}
}
