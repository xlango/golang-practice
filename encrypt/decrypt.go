package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

var publickey = FileLoad("encrypt/public.pem")

var privatekey = FileLoad("encrypt/privatekey.pem")

func RSAEncrypt(orgidata []byte) ([]byte, error) {

	block, _ := pem.Decode(publickey)
	if block == nil {
		return nil, errors.New("public key is bad")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, orgidata) //加密

}

func RSADecrypt(cipertext []byte) ([]byte, error) {

	block, _ := pem.Decode(privatekey)
	if block == nil {
		return nil, errors.New("public key is bad")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipertext)

}

func FileLoad(filepath string) []byte {

	privatefile, err := os.Open(filepath)
	defer privatefile.Close()

	if err != nil {
		return nil
	}

	privateKey := make([]byte, 2048)
	num, err := privatefile.Read(privateKey)
	return privateKey[:num]

}

func main() {

	var data []byte

	var err error
	data, err = RSAEncrypt([]byte("QQ1329172872"))

	if err != nil {
		fmt.Println("错误", err)
	}

	fmt.Println("加密：", base64.StdEncoding.EncodeToString(data))
	origData, err := RSADecrypt(data) //解密

	if err != nil {
		fmt.Println("错误", err)
	}

	fmt.Println("解密:", string(origData))
	//pk := FileLoad("myprivatekey.pem")
	//fmt.Println(string(pk))
}
