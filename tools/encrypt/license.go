package encrypt

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"math/rand"
)

var (
	prvKey, pubKey = GenRsaKey()
)

func GenLicense(plaintext string) (string, string, error) {
	// 生成密钥
	key := genKey(16)
	// 生成加密数据
	ciphertext, err := genCipher(key, []byte(plaintext))
	if err != nil {
		return "", "", err
	}
	// 生成数据摘要
	digest, err := genDigest(ciphertext)
	if err != nil {
		return "", "", err
	}
	// 生成签名 key+hash
	sign, err := genSign(key, digest)
	if err != nil {
		return "", "", err
	}
	// 生成license
	license, err := genLicense(key, ciphertext, sign)
	if err != nil {
		return "", "", err
	}
	// 还要生成存储信息
	return base64.StdEncoding.EncodeToString(license), base64.StdEncoding.EncodeToString(digest[:]), nil
}

func genKey(n int) []byte {
	key := make([]byte, n)
	if n == 0 {
		return key
	}
	_, _ = rand.Read(key)
	return key
}

func genCipher(key []byte, plaintext []byte) ([]byte, error) {
	auth := []byte("这是关联数据,可以为空")
	cipherText, err := GcmEncrypt(plaintext, key, auth)
	if err != nil {
		return nil, err
	}
	// auth 需要规定大小,或者
	return cipherText, nil
}

func genSign(key, digest []byte) ([]byte, error) {
	dst := make([]byte, len(key)+len(digest))
	copy(dst, append(key, digest[:]...))
	sign, err := RsaSignWithSha256(dst, prvKey)
	if err != nil {
		return nil, err
	}
	return sign, err
}

func genDigest(ciphertext []byte) ([]byte, error) {
	digest := sha256.Sum256(ciphertext)
	return digest[:], nil
}

func genLicense(key, cipherText, sign []byte) ([]byte, error) {
	licenseBuf := bytes.NewBuffer([]byte{})
	var err error
	if err = binary.Write(licenseBuf, binary.BigEndian, key); err != nil {
		return nil, err
	}
	if err = binary.Write(licenseBuf, binary.BigEndian, uint32(len(cipherText))); err != nil {
		return nil, err
	}
	if err = binary.Write(licenseBuf, binary.BigEndian, cipherText); err != nil {
		return nil, err
	}
	if err = binary.Write(licenseBuf, binary.BigEndian, sign); err != nil {
		return nil, err
	}
	return licenseBuf.Bytes(), nil
}
