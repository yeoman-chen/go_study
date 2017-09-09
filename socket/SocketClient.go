package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func sender(conn net.Conn) {
	words := "hello world"
	conn.Write([]byte(words))
	fmt.Println("send over")
}

//测试粘包问题
func sender1(conn net.Conn) {
	for i := 0; i < 100; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		conn.Write([]byte(words))
	}
}
func main() {
	//解析端口
	server := "127.0.0.1:8090"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Param Fatal error: %s", err.Error())
		os.Exit(1)
	}
	//建立连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Conn Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")

	//sender(conn)
	//time.Sleep(10 * 1e9)

	//测试粘包问题
	defer conn.Close()
	go sender1(conn)
	for {
		time.Sleep(1 * 1e9)
	}
}
