package main

import (
	"flag"
	"fmt"
	"gopl-exercise/common"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

// 首先调用该函数，将变量port注册到flag.CommandLine.Var中，
// 然后返回一个指针指向port
var port = common.PortFlag("port", 8899, "the port")

func main() {
	// 在这里会更新事先注册的参数port
	flag.Parse()
	// 读取到更新的port
	servIp := fmt.Sprintf("localhost:%d", *port)
	fmt.Printf("servIp:[%s]\n", servIp)
	listener, err := net.Listen("tcp", servIp)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
