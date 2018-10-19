package port

import (
	"strconv"
	"github.com/tarm/serial"
	"time"
)

func SelectPort() {
	println("Поиск портов")
	portClass := []string{"/dev/ttyS", "/dev/ttyACM", "/dev/ttyUSB"}

	for {
		for _, nameClass := range portClass {
			for i := 0; i < 10; i++ {

				portName := nameClass + strconv.Itoa(i)

				if RFIDPort == nil {
					RFIDPort = FindRFID(portName)
				}

				if RFIDPort != nil {
					println("Контролер пропусков подключен")
					return
				}
			}
		}
	}
}

func FindRFID(portName string) (port *Port) {
	portConfig := &serial.Config{Name: portName,
		Baud: 115200,
		ReadTimeout: time.Millisecond * 200,
	}

	port = &Port{Name: portName, Config: portConfig}
	connect := port.Connect()
	if connect == nil {
		return nil
	}

	_, err := connect.Write([]byte{0x95})
	if err != nil {
		connect.Close()
		return nil
	}

	buf := make([]byte, 5)

	_, err = connect.Read(buf)

	println(portName)
	if err != nil {
		println(err.Error())
		connect.Close()
		return nil
	} else {
		if buf[0] == 127 {
			println("Контроллер подключен к порту " + portName)
			return port
		} else {
			return nil
		}
	}
}
