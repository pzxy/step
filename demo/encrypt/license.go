package encrypt

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"math/rand"
)

var (
	prvKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA39O034WXb231C+qOI22qONcQaIiaEJoyoo03KrpUcGgUzvkl
ddUK2u1hxzqe1dzb82gJ4n0l8NpTDtFkxek/N1XgkkBHVOyvx9oAaOAEP413g/vF
32MMkVqFQiDLd5XkP7mkiLXMsY6uOKKEV920r91sVV9yqPpBwZLcG5/gKA3ybqtb
Gu9414lH/Q7cH6IYfLsw0psUmyjh4XhN9ZSGGCOogqVUJkp49PQy+/B3mEWZMBM4
uAznwPYnGSlWLqxlZWfjMnPNSea1oDGEmIMLln1hf3gH50azQySyI+G54eOcbwUn
QSAx6Y8gAlnj8IYo4m1XE2osnglmcsmQ6DzkRwIDAQABAoIBAFmSjJjDCpZiR6WM
KXPOKEf08TEUMauhOdBJeSFVTgNp86HcnPwp1N1GEFI9iGEHsNIY+ZO1B6D6jjCo
5Y46SL0pVeTRanp2+eWdvXo65BhW0dR0xnweVGI6+oc7PpHN33kNHuZC8va/pcMu
bo4HQI78qqUDvIwZEoWX8xPvCDKemjuH6WHTC7Au953oeQYDm8BMCzafftk0LX/o
h7bcR31+JzI5ukVs4FNDw6OpIbj8gHM77QAQvIH+l9NFG5xHBDXaFsdVQEpOKzl2
NCWG8PcXmxOa6ThQ3YP4nghazHjvUstlrIYCVlD3z+PeYcuE3GnFBM01mqfWy9yc
Xox0sQECgYEA/68gdQsXd9uCiJu6yo1w/WNMFcpoC3qOE1u/qy+Gk69IYelOALl0
bb5aAlaanVvY7wIo1Mayk3QNtwx4SA7bHDcSdX/397nJckP+iqdlKtScNfxpL4RO
tODGEzmTUfS1Zmho32dPdHjtqbBP6U/y5ni2O+vlt5LsQtdDS1PqHMcCgYEA4BqA
2IQrnuoFUw34Gg7uLWvqeBOVXNmN+9duAeZffH/vU+jBvpbEg/5jYiEvL4kiKkFR
EDktJm05tX/9CRL5zHoB3mcPilR69AHjBMJ1ZkHefmtUS8cz8m2/42/x4j5LvZGh
gOgWVRA2AWW5YGLNOHNdVsWcfWx5h3y3fCAMfIECgYAeMlSm/hRbd3VHJP0+hs3d
XyjROgJnuKWm66MLWKWQM5dc2Oz5cO2rOFvxAsFn7D0GnT6tQno8p+wmvjZGbFq8
TlTDw8VJYIvGDvm07mgoUsvQJGtXzUxuWE2MX7RibYAvBA2+tLjRjKak78nASEzr
oz4Cfa23rIZyrnGoJ9TyaQKBgGURDtKWqOV+PPB7+EwoN5ocdR26Fz5Mjw/o1B/f
OWj4eYKqEdZWQEIW09NJi/IwMkxEML7USlxuHfyQwBj7idSBoZ4fdI66EslAoJVg
7Xk0c5cn6FhUekT24fc8YSXbWmUf7GoqQQgTy4dNzRB+/nZa8NlVPYIDDbvgrw83
lOoBAoGAATHVaXCFYHX+p5AZcNfVHLKvCh8+Z6GgJQI630vSHIT11HnRpUKiQfwO
0FNNDi/7XgtUIkwftzoidS+izCP0K4W67x7SdfMrRqInFIMGmSG67DI6bm1Nl6Rl
1XCggnZO+kGi5jVNRjvXg7xD2qLMCKTh7kKqrY6H7eYE2/DbQC4=
-----END RSA PRIVATE KEY-----`

	pubKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA39O034WXb231C+qOI22q
ONcQaIiaEJoyoo03KrpUcGgUzvklddUK2u1hxzqe1dzb82gJ4n0l8NpTDtFkxek/
N1XgkkBHVOyvx9oAaOAEP413g/vF32MMkVqFQiDLd5XkP7mkiLXMsY6uOKKEV920
r91sVV9yqPpBwZLcG5/gKA3ybqtbGu9414lH/Q7cH6IYfLsw0psUmyjh4XhN9ZSG
GCOogqVUJkp49PQy+/B3mEWZMBM4uAznwPYnGSlWLqxlZWfjMnPNSea1oDGEmIML
ln1hf3gH50azQySyI+G54eOcbwUnQSAx6Y8gAlnj8IYo4m1XE2osnglmcsmQ6Dzk
RwIDAQAB
-----END PUBLIC KEY-----`
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
	sign, err := RsaSignWithSha256(dst, []byte(prvKey))
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
