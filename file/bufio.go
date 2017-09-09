package main 

import (
	"bufio"
	"io"
	"os"
	"time"
)

func main() {
	fi,err := os.Open("input.txt")
	if err != nil {panic(err)}
	defer fi.Close()
	r := bufio.NewReader(fi)//创建一个读取缓冲流

	fo,err := os.Create("output.txt")
	if err != nil { panic(err) }

	defer fo.Close()
	w := bufio.NewWriter(fo)//创建输出缓冲流

	buf := make([]byte,1024)
	for {
		n,err := r.Read(buf)
		if err != nil && err != io.EOF { panic(err)}
		if n == 0 { break }

		if n2,err := w.Write(buf[:n]); err != nil {
			panic(err)
		}else if n2 != n {
			panic("error in writing")
		}
	}

	time.Sleep(time.Second)
	if err  = w.Flush(); err != nil {panic(err)}
}