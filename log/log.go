package log

import (
	"../checkUser"
	"../dbConnect"
	"time"
	"log"
)

type Log struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Event string `json:"event"`
	Time  string `json:"time"`
	UUID  string `json:"uuid"`
}

func OpenDoorLog(user *checkUser.User) {
	dbConnect.GetDBConnect().Exec("INSERT INTO log (uuid, name, event, time) "+
		"VALUES ($1, $2, $3, $4)", user.UUID, user.Name, "OpenDoor", time.Now())
}

func GetLogs() []Log {

	rows, err := dbConnect.GetDBConnect().Query("SELECT id, uuid, name, event, time FROM log")
	if err != nil {
		log.Fatal("get user: " + err.Error())
	}
	defer rows.Close()

	logs := make([]Log, 0)

	for rows.Next() {

		Log := Log{}

		err = rows.Scan(&Log.ID, &Log.UUID, &Log.Name, &Log.Event, &Log.Time)
		if err != nil {
			log.Fatal("get user: " + err.Error())
		}

		logs = append(logs, Log)
	}

	return logs
}
