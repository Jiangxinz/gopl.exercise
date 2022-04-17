package common

import (
	"log"
	"net"
)

func MustListen(protocol string, servIp string) net.Listener {
	ret, err := net.Listen(protocol, servIp)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}
