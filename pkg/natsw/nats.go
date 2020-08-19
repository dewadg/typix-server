package natsw

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func ConnectNATS(host, port string) (*nats.EncodedConn, error) {
	url := fmt.Sprintf("nats://%s:%s", host, port)
	natsConn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)
}
