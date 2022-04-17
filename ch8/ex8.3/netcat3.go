package main

import (
	"flag"
	"fmt"
	"gopl-exercise/common"
	"log"
	"net"
	"os"
)

func main() {
	var port = common.PortFlag("port", 8000, "the port")
	flag.Parse()
	servIp := fmt.Sprintf("localhost:%d", *port)
	conn := common.MustDail("tcp", servIp)
	tcpConn := conn.(*net.TCPConn)
	done := make(chan struct{})
	go func() {
		common.MustCopy(os.Stdout, conn)
		tcpConn.CloseRead()
		log.Println("close tcpConn read")
		done <- struct{}{}
	}()
	go func() {
		common.MustCopy(conn, os.Stdin)
		log.Println("close tcpConn write")
		tcpConn.CloseWrite()
	}()
	<-done
}
