package webSocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"../transferData"
	rfidLog "../log"
)

var usersWs = make(map[*websocket.Conn]*Clients)

var upgrader = websocket.Upgrader{}

type Clients struct {
	Id int
}

type Message struct {
	Event string `json:"event"`
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
	}
	usersWs[ws] = &Clients{}

	Reader(ws)
}

func Reader(ws *websocket.Conn) {
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil { // Если есть ошибка при чтение из сокета вероятно клиент отключился, удаляем его сессию
			println(err.Error())
			break
		}

		if msg.Event == "OpenDoor" {
			transferData.OpenDoor()
		}
	}
}

func Sender() {
	for {

		logs := rfidLog.GetLogs()
		for ws := range usersWs {
			err := ws.WriteJSON(logs)
			if err != nil {
				ws.Close()
				delete(usersWs, ws)
			}
		}

		time.Sleep(time.Millisecond * 300)
	}
}
