package gorsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io"
	"io/ioutil"
)

var RSA = &RSASecurity{}

type RSASecurity struct {
	pubStr string          // 公钥字符串
	priStr string          // 私钥字符串
	pubkey *rsa.PublicKey  // 公钥
	prikey *rsa.PrivateKey // 私钥
}

// 设置公钥
func (rsas *RSASecurity) SetPublicKey(pubStr string) (err error) {
	rsas.pubStr = pubStr
	rsas.pubkey, err = rsas.GetPublickey()
	return err
}

// 设置私钥
func (rsas *RSASecurity) SetPrivateKey(priStr string) (err error) {
	rsas.priStr = priStr
	rsas.prikey, err = rsas.GetPrivatekey()
	return err
}

// *rsa.PublicKey
func (rsas *RSASecurity) GetPrivatekey() (*rsa.PrivateKey, error) {
	return getPriKey([]byte(rsas.priStr))
}

// *rsa.PrivateKey
func (rsas *RSASecurity) GetPublickey() (*rsa.PublicKey, error) {
	return getPubKey([]byte(rsas.pubStr))
}

// 公钥解密
func (rsas *RSASecurity) PubKeyDECRYPT(input []byte) ([]byte, error) {
	if rsas.pubkey == nil {
		return []byte(""), errors.New(`Please set the public key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := pubKeyIO(rsas.pubkey, bytes.NewReader(input), output, false)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(output)
}

// RsaDecrypt 公钥解密
func RsaDecrypt(data, pubKey []byte) ([]byte, error) {
	key, err := getPubKeyObj(pubKey)
	if err != nil {
		return nil, err
	}
	return decryptData(key, bytes.NewReader(data))
}

// 私钥加密
func (rsas *RSASecurity) PriKeyENCTYPT(input []byte) ([]byte, error) {
	if rsas.prikey == nil {
		return []byte(""), errors.New(`Please set the private key in advance`)
	}
	output := bytes.NewBuffer(nil)
	err := priKeyIO(rsas.prikey, bytes.NewReader(input), output, true)
	if err != nil {
		return []byte(""), err
	}
	return ioutil.ReadAll(output)
}

// RsaEncrypt 私钥加密
func RsaEncrypt(data, priKey []byte) ([]byte, error) {
	key, err := getPriKeyObj(priKey)
	if err != nil {
		return nil, err
	}
	return encryptData(key, bytes.NewReader(data))
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
	return pri2.(*rsa.PrivateKey), nil
}
