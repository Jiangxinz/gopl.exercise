package main

import (
	"flag"
	"fmt"
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

type portFlag struct {
	Port uint16
}

func (p *portFlag) String() string {
	return fmt.Sprintf("port: %d", p.Port)
}

func (p *portFlag) Set(s string) error {
	_, err := fmt.Sscanf(s, "%d", &p.Port)
	fmt.Printf("s:(%s), p.Port=%d\n", s, p.Port)
	return err
}

// 必须返回一个指针指向对应的变量，如果返回拷贝，则后续调用的Set函数会无法有效更新参数
func PortFlag(name string, value uint16, usage string) *uint16 {
	p := portFlag{value}
	flag.CommandLine.Var(&p, name, usage)
	fmt.Printf("after PortFlag: %d\n", p.Port)
	return &p.Port
}

// 首先调用该函数，将变量port注册到flag.CommandLine.Var中，
// 然后返回一个指针指向port
var port = PortFlag("port", 8899, "the port")

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
