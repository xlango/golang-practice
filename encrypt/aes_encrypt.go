package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AES加密
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//AES解密
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func main() {
	text := "test contenttest contenttest contenttest contenttest contenttest contenttest contenttest " +
		"contenttest contenttest contenttest contenttest contenttest contenttest contenttest " +
		"contenttest contenttest contenttest contenttest contenttest contenttest contenttest " +
		"contenttest contenttest contenttest contenttest contenttest contenttest contenttest contenttest content" +
		"test contenttest contenttest content"
	AesKey := []byte("0f90023fc9ae101e") //秘钥长度为16的倍数
	fmt.Printf("明文: %s\n秘钥: %s\n", text, string(AesKey))
	encrypted, err := AesEncrypt([]byte(text), AesKey)
	if err != nil {
		panic(err)
	}
	base64EncodeStr := base64.StdEncoding.EncodeToString(encrypted)
	fmt.Printf("加密后: %s\n", base64EncodeStr)
	base64DecodeStr, err := base64.StdEncoding.DecodeString(base64EncodeStr)
	origin, err := AesDecrypt(base64DecodeStr, AesKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密后明文: %s\n", string(origin))
}
