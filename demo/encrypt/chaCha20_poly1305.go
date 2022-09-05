package encrypt

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/pbkdf2"
)

func Poly1305En() {
	// - 一个 256 位的对称密钥(32byte)；用PBKDF2生产
	// - 一个 96 位的随机数；(12byte)
	// - 代加密数据明文，也就是私密信息；
	// - 关联数据，也就是公开信息。
	password := []byte("hello world")
	password = pbkdf2.Key(password, []byte("随便写1"), 4096, 32, sha256.New)
	fmt.Println(len(password))
	aead, err := chacha20poly1305.New(password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	nonce := make([]byte, aead.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	plaintext := []byte("this is a demo")
	ciphertext := aead.Seal(nil, nonce, plaintext, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ciphertext)
}
