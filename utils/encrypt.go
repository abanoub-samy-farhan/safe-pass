package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"encoding/hex"
)

func EncryptData(data string) string {
	defer func () {
		str := recover()
		if str != nil {
			fmt.Println(str)
		}
	}()

	key := []byte("2*EUH$@^9t$yGk6gUr8nzcKsBzf%zHbZ")
	plaintext := []byte(data)
	
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic("Error happend while generating a nonce" + err.Error())
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	encrypted := hex.EncodeToString(ciphertext)
	return encrypted
}

func DecryptData(encrypted string) string {
	defer func () {
		str := recover()
		if str != nil {
			fmt.Println(str)
		}
	}()

	chipertext, _ := hex.DecodeString(encrypted);

	key := []byte("2*EUH$@^9t$yGk6gUr8nzcKsBzf%zHbZ")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err)
	}

	nonce, ciphertext := chipertext[:gcm.NonceSize()], chipertext[gcm.NonceSize():]

	plainText, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	return string(plainText)
}