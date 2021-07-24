package cryptool

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

var (
	// mac地址
	mac string
	// 输入文件路径
	inputFile string
	// 输出文件路径
	outputFile string
	// opt表示操作类型，当opt为true时，表示解密文件，当opt为false是，表示加密文件，默认是false
	opt bool
)

// 对字符串进行sha256加密，取前16个字符
func encryptKey(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	bs := fmt.Sprintf("%x", h.Sum(nil))
	return []byte(bs)[:16]
}

// 文件读取
func read(path string) (content []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

// 写入文件
func write(path string, content []byte) (err error) {
	return ioutil.WriteFile(path, content, 0644)
}

// 使用aes gcm加密信息
func GcmEncrypt(plaintext []byte, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// 使用aes gcm解密信息
func GcmDecrypt(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}

// 加密配置信息
func encryptConfig(src []byte, mac string) (crypted []byte, err error) {
	return GcmEncrypt(src, encryptKey(mac))
}

// 解密配置信息
func decryptConfig(crypted []byte, mac string) (origData []byte, err error) {
	return GcmDecrypt(crypted, encryptKey(mac))
}

func Encrypt(macAddr string, opt bool, input, output string) error {
	mac = macAddr
	inputFile = input
	outputFile = output
	opt = opt
	var (
		err      error
		origData []byte
		optData  []byte
	)
	origData, err = read(inputFile)
	if err != nil {
		err = errors.Wrap(err, "failed to read input file")
		return err
	}
	if opt {
		// opt为true 解密
		optData, err = decryptConfig(origData, mac)
	} else {
		// opt为false 加密
		optData, err = encryptConfig(origData, mac)
	}
	if err != nil {
		err = errors.Wrap(err, "failed to operator input file content")
		return err
	}
	err = write(outputFile, optData)
	if err != nil {
		err = errors.Wrap(err, "failed to write operator content to output file")
		return err
	}
	return nil
}
