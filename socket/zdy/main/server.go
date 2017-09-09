package main

import (
	"fmt"
	"net"
	"os"
	"protocol"
	"time"
)

func main() {
	netListen, err := net.Listen("tcp", "localhost:6060")
	CheckError(err)

	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		//短连接的timeout,长连接见serverHeartBeat.go
		//conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))

		Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn)
		//模拟长连接心跳检测
		//go handleConnectionHeart(conn, 5)
	}
}

//长连接入口
func handleConnection(conn net.Conn) {

	// 缓冲区，存储被截断的数据
	tmpBuffer := make([]byte, 0)

	//接收解包，等待解包处理完成接收解包数据
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		//解包
		tmpBuffer = protocol.Depack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
	defer conn.Close()
}

func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			Log(string(data))
		}
	}
}

//长连接入口，模拟心跳
func handleConnectionHeart(conn net.Conn, timeout int) {

	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)

		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		Data := (buffer[:n])
		messnager := make(chan byte)
		//postda := make(chan byte)
		//心跳计时
		go HeartBeating(conn, messnager, timeout)
		//检测每次Client是否有数据传来
		go GravelChannel(Data, messnager)
		Log("receive data length:", n)
		Log(conn.RemoteAddr().String(), "receive data string:", string(Data))

	}
}

//心跳计时，根据GravelChannel判断Client是否在设定时间内发来信息
func HeartBeating(conn net.Conn, readerChannel chan byte, timeout int) {
	select {
	case fk := <-readerChannel:
		Log(conn.RemoteAddr().String(), "receive data string:", string(fk))
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		//conn.SetReadDeadline(time.Now().Add(time.Duration(5) * time.Second))
		break
	case <-time.After(time.Second * 5):
		Log("It's really weird to get Nothing!!!")
		conn.Close()
	}

}

func GravelChannel(n []byte, mess chan byte) {
	for _, v := range n {
		mess <- v
	}
	close(mess)
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
