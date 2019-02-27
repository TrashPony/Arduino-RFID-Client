package checkUser

import (
	"github.com/TrashPony/Arduino-RFID-Client/dbConnect"
	"log"
)

type User struct {
	Name string
	UUID string
}

func UUIDCheck(uuid string) (*User, bool) {
	rows, err := dbConnect.GetDBConnect().Query("SELECT name, uuid FROM users WHERE uuid=$1", uuid)
	if err != nil {
		log.Fatal("get user: " + err.Error())
	}
	defer rows.Close()

	user := &User{}

	for rows.Next() {
		err = rows.Scan(&user.Name, &user.UUID)
		if err != nil {
			log.Fatal("get user: " + err.Error())
		}
	}

	if user.Name != "" {
		return user, true
	} else {
		return nil, false
	}
}
