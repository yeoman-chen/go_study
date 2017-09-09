package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	file, err := os.Create("dst.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file1, err := os.Open("20.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()

	img, _ := jpeg.Decode(file1)

	//剪裁图片
	//jpg := image.NewNRGBA(image.Rect(0, 0, 100, 100))
	//draw.Draw(jpg, img.Bounds().Add(image.Pt(10, 10)), img, img.Bounds().Min, draw.Src) //截取图片的一部分
	//转为灰色
	jpg := image.NewGray(img.Bounds())                            //NewGray
	draw.Draw(jpg, jpg.Bounds(), img, img.Bounds().Min, draw.Src) //原始图片转换为灰色图片
	jpeg.Encode(file, jpg, nil)
}
