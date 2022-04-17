package common

import "fmt"
import "flag"

type portFlag struct {
	Port uint16
}

func (p *portFlag) String() string {
	return fmt.Sprintf("port: %d", p.Port)
}

func (p *portFlag) Set(s string) error {
	_, err := fmt.Sscanf(s, "%d", &p.Port)
	return err
}

// 必须返回一个指针指向对应的变量，如果返回拷贝，则后续调用的Set函数会无法有效更新参数
func PortFlag(name string, value uint16, usage string) *uint16 {
	p := portFlag{value}
	flag.CommandLine.Var(&p, name, usage)
	return &p.Port
}

type depthFlag struct {
	Depth uint16
}

func (d *depthFlag) String() string {
	return fmt.Sprintf("depth: %d", d.Depth)
}

func (d *depthFlag) Set(s string) error {
	_, err := fmt.Sscanf(s, "%d", &d.Depth)
	return err
}

func DepthFlag(name string, value uint16, usage string) *uint16 {
	d := depthFlag{value}
	flag.CommandLine.Var(&d, name, usage)
	return &d.Depth
}
