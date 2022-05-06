package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"math"
	"net/http"
)

func main() {
	faceImageBin, err := Download("https://img0.baidu.com/it/u=234305478,3590860473&fm=253&fmt=auto&app=120&f=JPEG?w=550&h=756")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(len(faceImageBin)/1024, "KB")
	imageBytes, err := ioutil.ReadFile("/Users/pzxy/WorkSpace/golang/tedge/door2/guanlin_door_driver/internal/guanlin/util/logo/256.png")
	if err != nil {
		t.Error(err)
		return
	}
	ioutil.WriteFile("/Users/pzxy/WorkSpace/golang/tedge/door2/guanlin_door_driver/internal/guanlin/util/logo/1.jpeg", imageBytes, 0644)

	resize, err := ConvertImage(imageBytes, 80, 200)
	if err != nil {
		t.Error(err)
		return
	}

	ioutil.WriteFile("/Users/pzxy/WorkSpace/golang/tedge/door2/guanlin_door_driver/internal/guanlin/util/logo/2.jpeg", resize, 0644)

}

func ConvertImage(content []byte, minKB uint, maxKB uint) ([]byte, error) {
	// 判断类型
	imageType := http.DetectContentType(content)
	switch imageType {
	case "image/jpeg", "image/png", "image/jpg":
	default:
		return nil, fmt.Errorf("invalid image type:%s", imageType)
	}
	// 转换
	img, n, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	// 压缩
	context.Context()
	fmt.Println(n)
	imgSize := img.Bounds().Dy() * img.Bounds().Dx()
	fmt.Println(img.Bounds().Dx())
	fmt.Println(img.Bounds().Dy())
	// 分辨率*位深/8
	fmt.Println(imgSize * 3 / 8 / 1024)

	// 判断尺寸
	if imgSize > int(minKB<<10) && imgSize < int(maxKB<<10) {
		if imageType == "image/jpeg" || imageType == "image/jpg" {
			return content, nil
		}
		buf := bytes.Buffer{}
		if err = jpeg.Encode(&buf, img, nil); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
	// 计算缩放比例，先选出合适大小，根据以前x，y比例选出最合适的宽度
	v := (int(minKB+maxKB) / 2) << 10
	vv := float64(v * img.Bounds().Dx() / img.Bounds().Dy())
	width := uint(math.Sqrt(vv))
	fmt.Println(width)
	buf := bytes.Buffer{}
	m := resize.Resize(width, 0, img, resize.NearestNeighbor)
	if err = jpeg.Encode(&buf, m, nil); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// todo 图片拉伸
