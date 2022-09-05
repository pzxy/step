package encrypt

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"math/big"
)

// RsaEncrypt 公钥加密
func RsaEncrypt(data, keyBytes []byte) ([]byte, error) {
	// 解密pem格式的公钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	// 加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// RsaEncrypt 私钥解密
func RsaDecrypt(ciphertext, keyBytes []byte) ([]byte, error) {
	// 获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("private key error")
	}
	// 解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GenRsaKey 生成密钥
func GenRsaKey() (priKey, pubKey []byte) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	priKey = pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubKey = pem.EncodeToMemory(block)
	return
}

// RsaSignWithSha256 签名
func RsaSignWithSha256(data []byte, keyBytes []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("private key error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// RsaVerySignWithSha256 验证签名
func RsaVerySignWithSha256(data, signData, keyBytes []byte) error {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return errors.New("public key error")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		return err
	}
	return nil
}

// RsaPriEncrypt 私钥加密
func RsaPriEncrypt(data, priKey []byte) ([]byte, error) {
	key, err := getPriKeyObj(priKey)
	if err != nil {
		return nil, err
	}
	return encryptData(key, bytes.NewReader(data))
}

// RsaPubDecrypt 公钥解密
func RsaPubDecrypt(data, pubKey []byte) ([]byte, error) {
	key, err := getPubKeyObj(pubKey)
	if err != nil {
		return nil, err
	}
	return decryptData(key, bytes.NewReader(data))
}

func getPriKeyObj(priKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(priKey)
	if block == nil {
		return nil, errors.New("get private key error")
	}
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return pri, nil
	}
	pri2, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := pri2.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("*rsa.PrivateKey type assert fail")
	}
	return key, nil
}

func getPubKeyObj(publickey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publickey)
	if block == nil {
		return nil, errors.New("get public key error")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("*rsa.PublicKey type assert fail")
	}
	return key, nil
}

func encryptData(pri *rsa.PrivateKey, r io.Reader) ([]byte, error) {
	var err error
	w := bytes.NewBuffer([]byte{})

	var b []byte
	size := 0
	k := ((pri.N.BitLen() + 7) / 8) - 11
	buf := make([]byte, k)
	for {
		size, err = r.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if size < k {
			b = buf[:size]
		} else {
			b = buf
		}
		b, err = priKeyEncrypt(rand.Reader, pri, b)
		if err != nil {
			return nil, err
		}
		if _, err = w.Write(b); err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}

func decryptData(pub *rsa.PublicKey, in io.Reader) ([]byte, error) {
	var err error
	w := bytes.NewBuffer([]byte{})

	k := (pub.N.BitLen() + 7) / 8
	buf := make([]byte, k)
	var b []byte
	size := 0
	for {
		size, err = in.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if size < k {
			b = buf[:size]
		} else {
			b = buf
		}
		b, err = pubKeyDecrypt(pub, b)
		if err != nil {
			return nil, err
		}
		if _, err = w.Write(b); err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}

func priKeyEncrypt(rand io.Reader, priKey *rsa.PrivateKey, hashed []byte) ([]byte, error) {
	tLen := len(hashed)
	k := (priKey.N.BitLen() + 7) / 8
	if k < tLen+11 {
		return nil, errors.New("data length error")
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k], hashed)
	m := new(big.Int).SetBytes(em)
	c, err := decrypt(rand, priKey, m)
	if err != nil {
		return nil, err
	}
	src := c.Bytes()
	numPaddingBytes := len(em) - len(src)
	for i := 0; i < numPaddingBytes; i++ {
		em[i] = 0
	}
	copy(em[numPaddingBytes:], src)
	return em, nil
}

func pubKeyDecrypt(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	k := (pub.N.BitLen() + 7) / 8
	if k != len(data) {
		return nil, errors.New("data length error")
	}
	m := new(big.Int).SetBytes(data)
	if m.Cmp(pub.N) > 0 {
		return nil, errors.New("public key error")
	}
	m.Exp(m, big.NewInt(int64(pub.E)), pub.N)
	d := leftPad(m.Bytes(), k)
	if d[0] != 0 {
		return nil, errors.New("data error")
	}
	if d[1] != 0 && d[1] != 1 {
		return nil, errors.New("private key error")
	}
	var i = 2
	for ; i < len(d); i++ {
		if d[i] == 0 {
			break
		}
	}
	i++
	if i == len(d) {
		return nil, nil
	}
	return d[i:], nil
}

func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}

func decrypt(random io.Reader, priv *rsa.PrivateKey, c *big.Int) (m *big.Int, err error) {
	if c.Cmp(priv.N) > 0 {
		err = errors.New("decryption error")
		return
	}
	var ir *big.Int
	if random != nil {
		var r *big.Int
		for {
			r, err = rand.Int(random, priv.N)
			if err != nil {
				return
			}
			if r.Cmp(big.NewInt(0)) == 0 {
				r = big.NewInt(1)
			}
			var ok bool
			ir, ok = modInverse(r, priv.N)
			if ok {
				break
			}
		}
		bigE := big.NewInt(int64(priv.E))
		rpowe := new(big.Int).Exp(r, bigE, priv.N)
		cCopy := new(big.Int).Set(c)
		cCopy.Mul(cCopy, rpowe)
		cCopy.Mod(cCopy, priv.N)
		c = cCopy
	}
	if priv.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, priv.D, priv.N)
	} else {
		m = new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
		m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, priv.Primes[0])
		}
		m.Mul(m, priv.Precomputed.Qinv)
		m.Mod(m, priv.Primes[0])
		m.Mul(m, priv.Primes[1])
		m.Add(m, m2)

		for i, values := range priv.Precomputed.CRTValues {
			prime := priv.Primes[2+i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}
	if ir != nil {
		m.Mul(m, ir)
		m.Mod(m, priv.N)
	}

	return
}

func modInverse(a, n *big.Int) (ia *big.Int, ok bool) {
	g := new(big.Int)
	x := new(big.Int)
	y := new(big.Int)
	g.GCD(x, y, a, n)
	if g.Cmp(big.NewInt(1)) != 0 {
		return
	}
	if x.Cmp(big.NewInt(1)) < 0 {
		x.Add(x, n)
	}
	return x, true
}
