package main

import (
	"io"
	"log"
	"net"
	"os"
)

func MustDail(protocal string, ipAddr string) net.Conn {
	ret, err := net.Dial(protocal, ipAddr)
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

func main() {
	conn := MustDail("tcp", "localhost:8000")
	tcpConn := conn.(*net.TCPConn)
	done := make(chan struct{})
	go func() {
		MustCopy(os.Stdout, conn)
		tcpConn.CloseRead()
		log.Println("close tcpConn read")
		done <- struct{}{}
	}()
	go func() {
		MustCopy(conn, os.Stdin)
		// type assertion
		log.Println("close tcpConn write")
		tcpConn.CloseWrite()
	}()
	<-done
}
