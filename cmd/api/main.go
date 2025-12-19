package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kaizakin/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true) // setting up logrus to tell which line the error occured in logs.

	var r *chi.Mux = chi.NewRouter()// setting up a new router.
	handlers.Handler(r);// handlers custom

	fmt.Println("Starting GO API service...")

	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

 	err :=  http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err) // if there's any error print the log with error level logrus takes care of this.
	}
}
