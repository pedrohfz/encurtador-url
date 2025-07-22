package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"log"
)

var secretKey = []byte("f3a1c7e8b2d4f6a9e0c5b8d2a7e3f9c1")

func Encrypt(originalUrl string) string {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	plainText := []byte(originalUrl)
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	iv := cipherText[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		log.Fatal(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return hex.EncodeToString(cipherText)
}

func Decrypt(encryptedUrl string) string {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		log.Fatal(err)
	}

	cipherText, err := hex.DecodeString(encryptedUrl)
	if err != nil {
		log.Fatal(err)
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}
