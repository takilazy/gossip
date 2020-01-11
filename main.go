package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/takilazy/gossip/controllers"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir(".web/")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(".web/static/"))))


	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user", controllers.Get).Methods("GET")


	port := os.Getenv("PORT")

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, handlers.LoggingHandler(os.Stdout, router))
	if err != nil {
		fmt.Print(err)
	}
}