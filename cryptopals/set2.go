package main

import (
	"crypto/aes"
	"log"
)

func abs(v int) int {
	if v < 0 {
		return -1 * v
	}
	return v
}

func PKCSPadding(text []byte, bytes int) []byte {

	sizeText := len(text) % bytes

	pad := abs(bytes - sizeText)

	for i := 0; i < pad; i++ {
		text = append(text, byte(pad))
	}
	return text
}

func aesCBCEncrypt(key string, plainText []byte) []byte {
	iv := make([]byte, aes.BlockSize)
	for i := 0; i < aes.BlockSize; i++ {
		iv[i] = 0
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	if len(plainText)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	ciphertext := make([]byte, len(plainText))
	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		block.Encrypt(ciphertext[i:i+aes.BlockSize], xor(iv, plainText[i:i+aes.BlockSize]))
		iv = ciphertext[i : i+aes.BlockSize]
		log.Println(string(iv))
	}

	return ciphertext
}

func aesCBCDecrypt(key string, ciphertext []byte) []byte {
	iv := make([]byte, aes.BlockSize)
	for i := 0; i < aes.BlockSize; i++ {
		iv[i] = 0
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		panic("plain is not a multiple of the block size")
	}

	plainText := make([]byte, len(ciphertext))
	for i := 0; i < len(plainText); i += aes.BlockSize {
		block.Decrypt(plainText[i:i+aes.BlockSize], xor(iv, ciphertext[i:i+aes.BlockSize]))
		iv = plainText[i : i+aes.BlockSize]
	}

	return plainText
}

func xor(a, b []byte) []byte {


	res := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] ^ b[i]
	}

	return res
}
