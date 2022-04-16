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
	done := make(chan struct{})
	go func() {
		MustCopy(os.Stdout, conn)
		log.Println("Done")
		done <- struct{}{}
	}()
	MustCopy(conn, os.Stdin)
	log.Println("conn.Close()")
	// type assertion
	tcpConn := conn.(*net.TCPConn)
	tcpConn.CloseWrite()
	log.Println("read channal Done")
	<-done
	tcpConn.CloseRead()
}
