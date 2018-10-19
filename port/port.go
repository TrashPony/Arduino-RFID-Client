package port

import (
	"github.com/tarm/serial"
)

var RFIDPort *Port

type Port struct {
	Name string
	Config *serial.Config
	Connection *serial.Port
}

func (p *Port) Connect () (connect *serial.Port)   {
	connect, err := serial.OpenPort(p.Config)
	if err != nil {
		return nil
	} else {
		p.Connection = connect
		return connect
	}
}
