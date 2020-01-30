package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {
	demo1()
}

func demo1() {
	const size = 300
	pic := image.NewGray(image.Rect(0, 0, size, size))
	//设置北京为白色
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	//绘制正选函数的轨迹
	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Sin(s)*size/2
		pic.SetGray(x, int(y), color.Gray{0})
	}
	//写入图片文件
	file,err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file,pic)
	file.Close()

}
