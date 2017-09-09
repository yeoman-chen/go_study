package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {

	ddd, _ := base64.StdEncoding.DecodeString("123456.jpg") //
	err2 := ioutil.WriteFile("./output.jpg", ddd, 0666)
	fmt.Println(err2)
}
