package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {

	fmt.Println(runtime.GOOS)     //取得操作系统版本
	fmt.Println(runtime.Compiler) //取得编译器名称-gc
	fmt.Println(runtime.NumCPU()) //取得CPU多核的数目
	fmt.Println(runtime.Version())
	//fmt.Println(os.Environ())

	err := os.Setenv("XIAO", "xiaochen") //临时设置 系统环境变量
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(os.Getenv("XIAO")) //获取环境变量
	fmt.Println(os.Getenv("GOPATH"))
	fmt.Println(os.Getenv("OS"))

	//获取全部系统环境变量 获取的是 key=val 的[]string
	for _, v := range os.Environ() {
		str := strings.Split(v, "=")
		fmt.Printf("key=%s,val=%s\n", str[0], str[1])
	}
}
