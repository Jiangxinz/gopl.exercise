package main

import (
	"bufio"
	"flag"
	"fmt"
	"gopl-exercise/common"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(tcpConn *net.TCPConn, s string, delay time.Duration, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	fmt.Fprintln(tcpConn, "\t", strings.ToUpper(s))
	time.Sleep(delay)
	fmt.Fprintln(tcpConn, "\t", s)
	time.Sleep(delay)
	fmt.Fprintln(tcpConn, "\t", strings.ToLower(s))
}

func resetTimer(active chan struct{}) {
	active <- struct{}{}
}

func handleConn(conn net.Conn) {
	var tcpConn = conn.(*net.TCPConn)
	var input = bufio.NewScanner(tcpConn)
	var active = make(chan struct{})
	// TODO: let expired be func parameter and os.args
	var expired = 10 * time.Second
	var wg sync.WaitGroup

	go func() {
		for input.Scan() {
			go resetTimer(active)
			go echo(tcpConn, input.Text(), 1*time.Second, &wg)
		}
		fmt.Println("sender finished")
	}()

	var ticker = time.NewTicker(expired)
	for {
		select {
		case <-ticker.C:
			fmt.Println("reach expired time")
			// FIXME: Close() or CloseWrite()?
			defer tcpConn.CloseWrite()
			// defer tcpConn.Close()
			fmt.Println("Wait send residual msg")
			wg.Wait()
			return
		case <-active:
			// reset timer
			ticker.Stop()
			ticker = time.NewTicker(expired)
		}
	}
}

func main() {
	var port = common.PortFlag("port", 8000, "the port")
	flag.Parse()
	var servIp = fmt.Sprintf("localhost:%d", *port)
	var listener = common.MustListen("tcp", servIp)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}
