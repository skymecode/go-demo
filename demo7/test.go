package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

type rsaSecurity struct {
}

var RsaSecurity rsaSecurity

func (rs *rsaSecurity) RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	p, _ := pem.Decode(origData)
	if p == nil {
		return nil, errors.New("public key error")
	}
	pub, err := x509.ParsePKIXPublicKey(p.Bytes)
	if err != nil {
		return nil, err
	}
	//类型断言
	res := pub.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, res, origData)
}

// 面向对象编程
func main() {
	//data, _ := RsaSecurity.RsaEncrypt([]byte(encrypt))
}
