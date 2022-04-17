package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func mustFprintf(w io.Writer, format string, s string) {
	if _, err := fmt.Fprintf(w, format, s); err != nil {
		log.Fatal(err)
	}
}

func echo(conn *net.TCPConn, s string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	mustFprintf(conn, "\t%s\n", strings.ToUpper(s))
	time.Sleep(delay)
	mustFprintf(conn, "\t%s\n", s)
	time.Sleep(delay)
	mustFprintf(conn, "\t%s\n", strings.ToLower(s))
}

// TODO: 如果客户端直接发EOF
func handleEcho(tcpConn *net.TCPConn, input *bufio.Scanner, wg *sync.WaitGroup) {
	starter := make(chan struct{})
	go func() {
		<-starter
		fmt.Println("closer start")
		wg.Wait()
		fmt.Println("tcp connection close write")
		tcpConn.CloseWrite()
	}()
	var firstTime bool = true
	for input.Scan() {
		if firstTime {
			starter <- struct{}{}
			firstTime = false
		}
		wg.Add(1)
		go echo(tcpConn, input.Text(), 1*time.Second, wg)
	}
	if firstTime {
		fmt.Println("tcp connection close read&write")
		tcpConn.Close()
	} else {
		fmt.Println("tcp connection close read")
		tcpConn.CloseRead()
	}
}

func handleConn(conn net.Conn) {
	tcpConn := conn.(*net.TCPConn)
	input := bufio.NewScanner(tcpConn)
	var wg sync.WaitGroup
	handleEcho(tcpConn, input, &wg)
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
