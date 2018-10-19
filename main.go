package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"./webSocket"
	"./port"
	"./transferData"
)

func main() {

	port.SelectPort()

	go transferData.RfidController()

	router := mux.NewRouter()
	router.HandleFunc("/ws", webSocket.HandleConnections)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	go webSocket.Sender()

	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Panic(err)
	}
}
