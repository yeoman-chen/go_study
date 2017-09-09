package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//建立socket，监听端口
	netListen, err := net.Listen("tcp", "localhost:8090")
	CheckError(err)
	defer netListen.Close()
	Log("Waitring for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), "tcp connect success")
		go handleConnection(conn)
	}

}

//处理连接
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	Log("Waiting for clients")
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), "connectiion error:", err)
			return
		}
		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))

	}
}

func Log(v ...interface{}) {
	log.Println(v...)
}
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
