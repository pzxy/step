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
MIICXgIBAAKBgQDRYOYlMyLpdQ3aw2chT0oY5RVUVQHj2GdLwYR5Bwz0mZMi+zYN
eNfWnGJ5ILT9SwQc4ubv9nfZHmnc0MCA53tMLeraH1Sr2C84Izjn/XHNBotqVhOb
CHcPloaFqEL8dEb77j1YE8xQtD7SpydLBb8DKu6hDXJQoDuW0CVO5dLLeQIDAQAB
AoGBAJSxagAVFsAQ3uvzRTREqOyRE9q2HAeBUapdHgUNWsgCoJmBKdvba/z/Rneb
IK2ZLqyM/1B/CpHopWmp3mws+EPSm91c7+MRRS+Gt34A00m0+ytys3YYSlJeM2Iq
1dgbtFNOlceJRn+M4HCw31A1j7YtkuLmQgFDmWWCr8j7AuUBAkEA41ett/E4/XTS
Psd91cB4RPleF9gPwQZ2t0OAMG2z4jCKdLYiO3kVEBWmm02qd0KBxx1cHuIOv9aU
oaEFleajsQJBAOvFhjtsHgVPljKN4X17IIt2sosL2eRrUCozYx3pkyXN4ybzNzbd
r0kHtb9OcFtxaDO0fRARKBIgKMh6H5qNfkkCQQCO+ZVHGW29+LpAwTViOKlqt2l1
lWxS+L2wg8MaseNvXRe2TtU4ke2tL/CXu0B5wFnd35kP0xtzin/vJlVc8LABAkEA
y0Omv1It65w6vGjvr1PYsgPqV9Am0TPnkApveFloQB5pqKnkv0uOFCMugLkqStvR
30nKzHBOIJpTLsuCtaco6QJAW0KMzlpE8ulfh/baogJzbiXiE73MYNaEWP5YOWZj
F3M93s4spkehPT2eV5Ix4B11HFI+NI8GaJbRVc6SjrP1Ng==
-----END RSA PRIVATE KEY-----`

	pubKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDRYOYlMyLpdQ3aw2chT0oY5RVU
VQHj2GdLwYR5Bwz0mZMi+zYNeNfWnGJ5ILT9SwQc4ubv9nfZHmnc0MCA53tMLera
H1Sr2C84Izjn/XHNBotqVhObCHcPloaFqEL8dEb77j1YE8xQtD7SpydLBb8DKu6h
DXJQoDuW0CVO5dLLeQIDAQAB
-----END PUBLIC KEY-----`
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
	cipherText, err := GcmEncrypt(plaintext, key)
	if err != nil {
		return nil, err
	}
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
