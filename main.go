package main

import (
	"github.com/TrashPony/Arduino-RFID-Client/port"
	"github.com/TrashPony/Arduino-RFID-Client/transferData"
	"github.com/TrashPony/Arduino-RFID-Client/webSocket"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
