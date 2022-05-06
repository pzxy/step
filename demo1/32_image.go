package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"math"
	"net/http"
	"testing"
)

package util

import (
"fmt"
"io/ioutil"
"testing"
)

func TestLoginKey(t *testing.T) {
	key := LoginKey("admin", "123")
	if key != "ab965a188646d3fb86c762991a01ab66" {
		t.Errorf("LoginKey() = %v, want %v", key, "ab965a188646d3fb86c762991a01ab66")
	}
}

func TestBusinessKey(t *testing.T) {
	got := BusinessKey("admin", "1534602542690", "b1f1b3f5-956f-4e60-abda-e095317d5fc3")
	if got != "3bbbe578ac2470240700ea89dade8256" {
		t.Errorf("BusinessKey() = %v, want %v", got, "3bbbe578ac2470240700ea89dade8256")
	}
}

// func TestEncrypt(t *testing.T) {
//    plaintext := []byte(`{"User":"system","Verify":"e10adc3949ba59abbe56e057f20f883e"}`)
//    kkk := "a5cc893b0db32f4f48963d516dff25cb"
//    key := []byte(kkk) //hex.DecodeString(hex.EncodeToString([]byte(kkk)))
//    iv := []byte(kkk[:16])
//
//    //加密
//    ciphertext, err := encrypt.AesEncrypt(plaintext, key, iv)
//    if err != nil {
//        panic(err)
//    }
//
//    //打印加密base64后密码
//    baseStr := base64.StdEncoding.EncodeToString(ciphertext)
//    fmt.Println(baseStr)
//
//    ct, _ := base64.StdEncoding.DecodeString(baseStr)
//    fmt.Println(ct)
//
//    //解密
//    plaintext, err = encrypt.AesDecrypt(ct, key, iv)
//    if err != nil {
//        panic(err)
//    }
//
//    //打印解密明文
//    fmt.Println(string(plaintext))
//
//    user := "system"
//    password := "123456"
//
//    aesKey := LoginKey(user, password)
//    params := make(map[string]interface{})
//    params["User"] = user
//    params["Verify"] = Md5(password)
//
//    plainText := `oiJ6a0f+7UhEbl+sjads0wu0/fiWaQlchoMY1ODr8O4GmBtRDnn4qEDSglYeStPbXdJd+Ic4Wu7pIzRuSjL8NQ==`
//    dd, err := Decrypt(plainText, aesKey)
//    t.Log(dd, err)
//
//    aesKey = "a5cc893b0db32f4f48963d516dff25cb"
//    text, err := Encrypt(params, aesKey)
//    if err != nil {
//        t.Error(err)
//    } else if text != "oiJ6a0f+7UhEbl+sjads0wu0/fiWaQlchoMY1ODr8O4GmBtRDnn4qEDSglYeStPbXdJd+Ic4Wu7pIzRuSjL8NQ==" {
//        t.Error("not expect", text)
//    }
// }

func TestEncodeImage(t *testing.T) {
	imageBytes, err := ioutil.ReadFile("/Users/pzxy/WorkSpace/golang/tedge/door2/guanlin_door_driver/internal/guanlin/util/logo/256.png")
	if err != nil {
		t.Error(err)
		return
	}
	_, err = EncodeImage(imageBytes)
	if err != nil {
		t.Error(err)
		return
	}

}

func TestImageResize(t *testing.T) {
	faceImageBin, err := Download("https://img0.baidu.com/it/u=234305478,3590860473&fm=253&fmt=auto&app=120&f=JPEG?w=550&h=756")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(len(faceImageBin)/1024, "KB")
	ioutil.WriteFile("/Users/pzxy/WorkSpace/golang/tedge/door2/guanlin_door_driver/internal/guanlin/util/logo/1.jpeg", faceImageBin, 0644)

	// imageBytes, err := ioutil.ReadFile("/Users/pzxy/WorkSpace/golang/tedge/door2/guanlin_door_driver/internal/guanlin/util/logo/256.png")
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	resize, err := ImageResize(faceImageBin, 1000)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(len(resize)/1024, "KB")

	ioutil.WriteFile("/Users/pzxy/WorkSpace/golang/tedge/door2/guanlin_door_driver/internal/guanlin/util/logo/2.jpeg", resize, 0644)

}

func TestConvertImage(t *testing.T) {

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
package util

import (
"bytes"
"context"
"fmt"
"image"
"image/jpeg"
_ "image/png"
"math"
"net/http"

"github.com/nfnt/resize"
)

// ImageResize 图片缩放
func ImageResize(content []byte, minWidth uint) (outputImgContent []byte, err error) {
	outputImgContent = content

	img, _, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		return
	}

	imgWidth := img.Bounds().Max.X // image.DecodeConfig(bytes.NewReader(content))
	if imgWidth >= int(minWidth) {
		return content, nil // 图片已经超过最少尺寸，不做缩放处理
	}

	width := minWidth
	buf := bytes.Buffer{}
	m := resize.Resize(width, 0, img, resize.NearestNeighbor)
	if err = jpeg.Encode(&buf, m, nil); err != nil {
		return
	}
	outputImgContent = buf.Bytes()
	return
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
