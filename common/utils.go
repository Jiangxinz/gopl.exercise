package common

import (
	"io"
	"log"
	"net"
)

func MustDail(protocol string, ipAddr string) net.Conn {
	ret, err := net.Dial(protocol, ipAddr)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

func MustListen(protocol string, ipAddr string) net.Listener {
	ret, err := net.Listen(protocol, ipAddr)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

func MustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
