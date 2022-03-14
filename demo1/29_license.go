package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

var (
	prvKey = `MIICXQIBAAKBgQC/T0ZZab4bUwP2rcReLRabR7K5M1aUFqr5PyzNkccaK0HiS40t
  X0w9+OQh30BYUeEMr/x/edwZt9+g3ArzSFVehqkmxdAGS/WbM+RSZ9FOxrTNY+0o
  zfWBm4aAuNLmUsvuH4h/Ud4bSaJ2WcAGcG4iGVbx2H4rM4Hua3JkMEmDMwIDAQAB
  AoGADfHQVeogQkudyVhHLPoqYK0sONWmJOs2ES7qrB1JHv4yMjWBl+IPY6EFt/ZW
  8Vkeh1c4l24hffpqIYaIgc0u87FvVjSc6RxWOqh6tYMQMmgeDq2czjJz/ki5XG7a
  jzftvxdYQ4cq9wtbhjcauydIXOr0Wtd8J+jGMyNg3IRTLIkCQQD4yLGBuCsJVgoj
  iJ267WCLBZyQwP6p3d6IzxMLvZoHQhD5axtFLypFSNW+VeQxlma+hVqJl1Cu5VRg
  k7jf1cJ9AkEAxNvQo8yGCeggwS+Z/+rD05ZjsxmGhDkJqOUh417R3nwJWxiuhEUk
  qRv30W+xfbrWBo94xTqYBDQ2fp/eACMbbwJAUjJU24wSxAnmXakkm/8T17rp6LWr
  Lkp83Oy12ToyqtU8MHwfzXLw32VaIYHVXEDZ/3al3DHfWXmxsBr+iQ3iLQJBAKf7
  u0pLGIYTov+3Ev6i4oAHcQQ/392snmWVQRm4q/XieklWAnW0WthDmXEKkrbrW/81
  JqWGnnnX5DihSttrGdkCQQD0AyilM6e3hUzTJwgbbtWn5PkzVRkIquJEzhIvkzF7
  ewp0ZJpCytX/jrjGjuyMhfk35D29sJRKLvYGbdUxYsTI`
	pubKey = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC/T0ZZab4bUwP2rcReLRabR7K5
  M1aUFqr5PyzNkccaK0HiS40tX0w9+OQh30BYUeEMr/x/edwZt9+g3ArzSFVehqkm
  xdAGS/WbM+RSZ9FOxrTNY+0ozfWBm4aAuNLmUsvuH4h/Ud4bSaJ2WcAGcG4iGVbx
  2H4rM4Hua3JkMEmDMwIDAQAB`
)

/**
	-----BEGIN RSA PRIVATE KEY-----
  MIICXQIBAAKBgQC/T0ZZab4bUwP2rcReLRabR7K5M1aUFqr5PyzNkccaK0HiS40t
  X0w9+OQh30BYUeEMr/x/edwZt9+g3ArzSFVehqkmxdAGS/WbM+RSZ9FOxrTNY+0o
  zfWBm4aAuNLmUsvuH4h/Ud4bSaJ2WcAGcG4iGVbx2H4rM4Hua3JkMEmDMwIDAQAB
  AoGADfHQVeogQkudyVhHLPoqYK0sONWmJOs2ES7qrB1JHv4yMjWBl+IPY6EFt/ZW
  8Vkeh1c4l24hffpqIYaIgc0u87FvVjSc6RxWOqh6tYMQMmgeDq2czjJz/ki5XG7a
  jzftvxdYQ4cq9wtbhjcauydIXOr0Wtd8J+jGMyNg3IRTLIkCQQD4yLGBuCsJVgoj
  iJ267WCLBZyQwP6p3d6IzxMLvZoHQhD5axtFLypFSNW+VeQxlma+hVqJl1Cu5VRg
  k7jf1cJ9AkEAxNvQo8yGCeggwS+Z/+rD05ZjsxmGhDkJqOUh417R3nwJWxiuhEUk
  qRv30W+xfbrWBo94xTqYBDQ2fp/eACMbbwJAUjJU24wSxAnmXakkm/8T17rp6LWr
  Lkp83Oy12ToyqtU8MHwfzXLw32VaIYHVXEDZ/3al3DHfWXmxsBr+iQ3iLQJBAKf7
  u0pLGIYTov+3Ev6i4oAHcQQ/392snmWVQRm4q/XieklWAnW0WthDmXEKkrbrW/81
  JqWGnnnX5DihSttrGdkCQQD0AyilM6e3hUzTJwgbbtWn5PkzVRkIquJEzhIvkzF7
  ewp0ZJpCytX/jrjGjuyMhfk35D29sJRKLvYGbdUxYsTI
  -----END RSA PRIVATE KEY-----

  -----BEGIN PUBLIC KEY-----
  MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC/T0ZZab4bUwP2rcReLRabR7K5
  M1aUFqr5PyzNkccaK0HiS40tX0w9+OQh30BYUeEMr/x/edwZt9+g3ArzSFVehqkm
  xdAGS/WbM+RSZ9FOxrTNY+0ozfWBm4aAuNLmUsvuH4h/Ud4bSaJ2WcAGcG4iGVbx
  2H4rM4Hua3JkMEmDMwIDAQAB
  -----END PUBLIC KEY-----
*/
func main() {
	plaintext := `{
    "appId": "cn.fdm.offlicensedemo",
    "issuedTime": 1595951714,
    "notBefore": 1538671712,
    "notAfter": 1640966400,
    "customerInfo": "亚马逊公司"
}`
	license, hash, err := GenLicense(plaintext)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(license)
	fmt.Println(hash)
}

func GenLicense(plaintext string) (string, string, error) {
	//生成密钥
	key := genKey(16)
	fmt.Println(key)
	//生成加密数据
	cipherText, err := GcmEncrypt([]byte(plaintext), key)
	if err != nil {
		return "", "", err
	}
	//生成数据摘要
	digest := sha256.Sum256(cipherText)
	fmt.Println(digest)
	//生成签名 key+hash
	dst := make([]byte, len(key)+len(digest))
	data := append(key, digest[:]...)
	copy(dst, data)
	sign := RsaSignWithSha256(dst, []byte(prvKey))
	//生成license
	licenseBuf := bytes.NewBuffer([]byte{})
	binary.Write(licenseBuf, binary.BigEndian, key)
	binary.Write(licenseBuf, binary.BigEndian, uint32(len(cipherText)))
	binary.Write(licenseBuf, binary.BigEndian, cipherText)
	binary.Write(licenseBuf, binary.BigEndian, sign)
	license := licenseBuf.Bytes()
	return hex.EncodeToString(license), hex.EncodeToString(digest[:]), nil
}

func verifyLicense(license string) {

}

func genKey(n int) []byte {
	key := make([]byte, n)
	if n == 0 {
		return key
	}
	_, _ = rand.Read(key)
	return key
}

//加密
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

//sign
