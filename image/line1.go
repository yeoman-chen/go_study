package main

import (
	"fmt"
	"image/jpeg"
	"os"
)

func main() {
	file, err := os.Open("test.jpg")

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Create("test1.jpg")

	if err != nil {
		fmt.Println(err)
	}

	defer file1.Close()

	img, err := jpeg.Decode(file) //解码
	if err != nil {
		fmt.Println(err)
	}
	jpeg.Encode(file1, img, &jpeg.Options{5}) //编码，但是将图像质量从100改成5

}
