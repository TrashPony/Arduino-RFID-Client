package transferData

import (
	"encoding/hex"
	"github.com/TrashPony/Arduino-RFID-Client/checkUser"
	doorLog "github.com/TrashPony/Arduino-RFID-Client/log"
	"github.com/TrashPony/Arduino-RFID-Client/port"
	"time"
)

func RfidController() {

	for {

		//port.RFIDPort.Connection.Write([]byte{0x90})

		uuid := make([]byte, 4)
		n, err := port.RFIDPort.Connection.Read(uuid)

		if n == 4 && err == nil {
			uuidString := hex.EncodeToString(uuid)
			println(uuidString)
			user, success := checkUser.UUIDCheck(uuidString)
			if success {
				doorLog.OpenDoorLog(user)
				OpenDoor()
			} else {
				// todo
			}
		} else {
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func OpenDoor() {
	port.RFIDPort.Connection.Write([]byte{0x91})
}
