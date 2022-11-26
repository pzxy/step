package gorsa

import (
	"encoding/base64"
)

// 私钥加密
func PriKeyEncrypt(data, privateKey string) (string, error) {

	grsa := RSASecurity{}
	grsa.SetPrivateKey(privateKey)

	rsadata, err := grsa.PriKeyENCTYPT([]byte(data))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(rsadata), nil
}

// 公钥解密
func PublicDecrypt(data, publicKey string) (string, error) {

	databs, _ := base64.StdEncoding.DecodeString(data)

	grsa := RSASecurity{}
	grsa.SetPublicKey(publicKey)

	rsadata, err := grsa.PubKeyDECRYPT([]byte(databs))
	if err != nil {
		return "", err
	}

	return string(rsadata), nil

}
