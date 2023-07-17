package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

// 填充
func Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// 填充的反向操作
func UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

func AESEncrypted(plainText string, keyStr string) (string, error) {
	h := sha256.New()
	h.Write([]byte(keyStr))
	key := h.Sum(nil)
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := c.BlockSize()
	encryptedData := Padding([]byte(plainText), blockSize)
	cryptedData := make([]byte, len(encryptedData))
	cbc := cipher.NewCBCEncrypter(c, key[:blockSize])
	cbc.CryptBlocks(cryptedData, encryptedData)
	return base64.URLEncoding.EncodeToString(cryptedData), nil
}

func AESDecrypted(plainText string, keyStr string) (string, error) {
	h := sha256.New()
	h.Write([]byte(keyStr))
	key := h.Sum(nil)
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := c.BlockSize()

	cipherText, _ := base64.URLEncoding.DecodeString(plainText)

	cryptedData := make([]byte, len(cipherText))
	cbc := cipher.NewCBCDecrypter(c, key[:blockSize])
	cbc.CryptBlocks(cryptedData, cipherText)

	cryptedData, err = UnPadding(cryptedData)
	if err != nil {
		return "", err
	}
	return string(cryptedData), nil
}
