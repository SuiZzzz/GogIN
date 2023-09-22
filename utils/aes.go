package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

const (
	Key = "j34[d10(*e596*/2jiw45@#wbdk,(dmw"
)

// Encrypt aes加密 格式：向量+密码+填充
func Encrypt(key, src []byte) (data []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("aes加密失败")
		return nil, err
	}

	// 填充密码为16的倍数，末尾补" "
	diff := block.BlockSize() - len(src)%block.BlockSize()
	if diff != 0 {
		temp := make([]byte, diff)
		temp = bytes.ReplaceAll(temp, []byte{0}, []byte{32})
		src = append(src, temp...)
	}

	text := make([]byte, aes.BlockSize+len(src))
	iv := text[:aes.BlockSize]

	// 生成向量iv
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		log.Println("随机向量生成失败")
		return nil, err
	}
	bm := cipher.NewCBCEncrypter(block, iv)
	bm.CryptBlocks(text[aes.BlockSize:], src)

	return text, nil
}

// Decrypt aes解密
func Decrypt(key, src []byte) (data []byte, err error) {

	iv := src[:aes.BlockSize]
	text := src[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	bm := cipher.NewCBCDecrypter(block, iv)
	bm.CryptBlocks(text, text)

	// 去除尾部填充
	index := bytes.IndexByte(text, 32)
	if index > -1 {
		text = text[:index]
	}

	return text, nil
}
