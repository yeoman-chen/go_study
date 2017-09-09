package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	file, err := os.Create("dst-merge.jpg")
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
	file2, err := os.Open("1.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file2.Close()

	img2, _ := jpeg.Decode(file2)

	jpg := image.NewNRGBA(image.Rect(0, 0, 300, 300))

	draw.Draw(jpg, jpg.Bounds(), img2, img2.Bounds().Min, draw.Over)                   //首先将一个图片信息存入jpg
	draw.Draw(jpg, jpg.Bounds(), img, img.Bounds().Min.Sub(image.Pt(0, 0)), draw.Over) //将另外一张图片信息存入jpg
	jpeg.Encode(file, jpg, nil)
}
