package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"

	"github.com/nfnt/resize"
)

func main() {
	for100()
}
func cImage() {
	// imageBytes, err := Download("https://img0.baidu.com/it/u=234305478,3590860473&fm=253&fmt=auto&app=120&f=JPEG?w=550&h=756")
	// if err != nil {
	// 	return
	// }
	imageBytes, err := ioutil.ReadFile("/Users/pzxy/WorkSpace/Go/src/step/demo1/2_1.jpeg")
	if err != nil {
		return
	}

	ioutil.WriteFile("./1.jpeg", imageBytes, 0644)

	resize, err := ConvertImage(imageBytes, 70, 80, 100)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ioutil.WriteFile("./2.jpeg", resize, 0644)
}

func for100() {
	imageBytes, err := Download("https://img0.baidu.com/it/u=234305478,3590860473&fm=253&fmt=auto&app=120&f=JPEG?w=550&h=756")
	if err != nil {
		return
	}
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return
	}
	var b []byte
	// curr := uint(img.Bounds().Dx())
	for i := 0; i < 100; i++ {
		b, err = resizeImage(uint(img.Bounds().Dx()), img)
		if err != nil {
			return
		}
	}
	ioutil.WriteFile("./3.jpeg", b, 0644)
}

func ConvertImage(content []byte, minKB uint, maxKB uint, n int) ([]byte, error) {
	if n <= 0 {
		// 未指定次数，默认10次，防止未知特殊情况。
		n = 10
	}
	img, name, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	if name == "jpeg" && compareSize(content, minKB, maxKB) == 0 {
		return content, nil
	}
	// 大小不合格，或者不是jpg，都转一下
	b, err := resizeImage(uint(img.Bounds().Dx()), img)
	if err != nil {
		return nil, err
	}
	// 先前的值
	preVal, curVal := uint(img.Bounds().Dx()), uint(img.Bounds().Dx())
	for i := 0; i < n; i++ {
		ret := compareSize(b, minKB, maxKB)
		// 太大了，要缩小
		if ret > 0 {
			curVal, preVal = scaleDown(curVal, preVal), curVal
			b, err = resizeImage(curVal, img)
			if err != nil {
				return nil, err
			}
			fmt.Println(len(b)>>10, "KB")
			continue
		}
		// 太小了，要放大
		if ret < 0 {
			curVal, preVal = scaleUp(curVal, preVal), curVal
			b, err = resizeImage(curVal, img)
			if err != nil {
				return nil, err
			}
			fmt.Println(len(b)>>10, "KB")
			continue
		}
		return b, nil
	}
	return nil, fmt.Errorf("image size should be between %d(kb) and %d(kb)", minKB, maxKB)
}

func scaleUp(currVal uint, preVal uint) uint {
	if currVal >= preVal {
		return currVal << 1
	}
	return (currVal + preVal) >> 1
}

func scaleDown(currVal uint, preVal uint) uint {
	if currVal <= preVal {
		return currVal >> 1
	}
	return (currVal + preVal) >> 1
}

func compareSize(content []byte, minKB uint, maxKB uint) int {
	padding := uint(10)
	imgSize := uint(len(content) >> 10)
	// 接近图片大小, 再保守一点。
	if maxKB-minKB <= padding*2 {
		maxKB = maxKB + padding
		minKB = minKB - padding
	}
	if imgSize >= maxKB-padding {
		return 1
	}
	if imgSize <= minKB+padding {
		return -1
	}
	return 0
}

func resizeImage(weight uint, img image.Image) ([]byte, error) {
	buf := bytes.Buffer{}
	m := resize.Resize(weight, 0, img, resize.NearestNeighbor)
	if err := jpeg.Encode(&buf, m, &jpeg.Options{Quality: 100}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Download(urlPath string) (data []byte, err error) {
	var (
		request  *http.Request
		response *http.Response
	)
	urlPath = fmt.Sprintf("%s", urlPath)
	request, err = http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		//  gLog.Error("NewRequest err", err)
		return nil, err
	}
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	response, err = doHTTPReqWithRetry(client, request, 3)
	if err != nil {
		fmt.Println("download err：", err)
		return nil, err
	}
	defer response.Body.Close()
	data, err = ioutil.ReadAll(response.Body)
	return data, err
}

func doHTTPReqWithRetry(client *http.Client, request *http.Request, retry int) (resp *http.Response, err error) {
	for ; retry > 0; retry-- {
		resp, err = client.Do(request)
		if err == nil {
			return
		}
	}
	return
}
