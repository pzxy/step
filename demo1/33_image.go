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
	//imageBytes, err := Download("https://img0.baidu.com/it/u=234305478,3590860473&fm=253&fmt=auto&app=120&f=JPEG?w=550&h=756")
	//if err != nil {
	//	return
	//}
	imageBytes, err := ioutil.ReadFile("/Users/pzxy/WorkSpace/Go/src/step/demo1/1.jpeg")
	if err != nil {
		return
	}
	ioutil.WriteFile("/Users/pzxy/WorkSpace/Go/src/step/demo1/1_1.jpeg", imageBytes, 0644)

	resize, err := ConvertImage2(imageBytes, 80, 200)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ioutil.WriteFile("/Users/pzxy/WorkSpace/Go/src/step/demo1/2_1.jpeg", resize, 0644)

}

func ConvertImage(content []byte, minKB uint, maxKB uint) ([]byte, error) {
	img, n, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	if n == "jpeg" && compareSize(content, minKB, maxKB) {
		return content, nil
	}

	// 大小不合格，或者不是jpg，都转一下
	if
	b, err := resizeImage(uint(img.Bounds().Dx()), img)
	if err != nil {
		return nil, err
	}
	if compareSize(b, minKB, maxKB) {
		return b, nil
	}
	return b, nil
}

func binarySearch(arr []uint, target uint) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func compareSize(content []byte, minKB uint, maxKB uint) bool {
	padding := uint(10)
	imgSize := uint(len(content) >> 10)
	//接近图片大小kb, 同时允许一定差值
	if maxKB-minKB <= padding*2 {
		maxKB = maxKB + padding
		minKB = minKB - padding
	}
	if imgSize > minKB+padding && imgSize < maxKB-padding {
		return true
	}
	return false
}

func resizeImage(weight uint, img image.Image) ([]byte, error) {
	buf := bytes.Buffer{}
	m := resize.Resize(weight, 0, img, resize.NearestNeighbor)
	if err := jpeg.Encode(&buf, m, nil); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ConvertImage2(content []byte, minKB uint, maxKB uint) ([]byte, error) {
	// 转换
	img, _, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	fmt.Println(img.Bounds().Dx())
	fmt.Println(img.Bounds().Dy())
	fmt.Println(len(content) / 1024)
	// 大小不合格，或者不是jpg，都转一下
	buf := bytes.Buffer{}
	m := resize.Resize(uint(img.Bounds().Dx()*2), 0, img, resize.NearestNeighbor)
	if err = jpeg.Encode(&buf, m, nil); err != nil {
		return nil, err
	}
	b := buf.Bytes()
	// 转换
	fmt.Println(len(b) / 1024)

	return b, nil

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
