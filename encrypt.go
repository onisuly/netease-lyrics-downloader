package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"time"
	"unicode/utf8"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RSAEncrypt(text string, pubKey string, modulus string) string {
	text = reverse(text)
	iText := new(big.Int)
	iText.SetString(hex.EncodeToString([]byte(text)), 16)

	iKey := new(big.Int)
	iKey.SetString(pubKey, 16)

	iModule := new(big.Int)
	iModule.SetString(modulus, 16)

	rsPow := new(big.Int)
	rsMod := new(big.Int)

	rsPow.Exp(iText, iKey, nil)
	rsMod.Mod(rsPow, iModule)

	return fmt.Sprintf("%x", rsMod)
}

func AESEncrypt(text string, key string) string {
	plainText := []byte(text)
	iv := []byte("0102030405060708")

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	padding := block.BlockSize() - len(plainText)%block.BlockSize()
	temp := bytes.Repeat([]byte{byte(padding)}, padding)
	plainText = append(plainText, temp...)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)

	return base64.StdEncoding.EncodeToString(cipherText)
}

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}
