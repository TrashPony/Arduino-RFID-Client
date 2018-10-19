package transferData

import (
	"../port"
	"time"
	"encoding/hex"
	"../checkUser"
	doorLog "../log"
)

func RfidController()  {

	for {


		port.RFIDPort.Connection.Write([]byte{0x90})

		uuid := make([]byte, 4)
		n, err := port.RFIDPort.Connection.Read(uuid)

		if n == 4 && err == nil {
			uuidString := hex.EncodeToString(uuid)
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

func OpenDoor()  {
	port.RFIDPort.Connection.Write([]byte{0x91})
}