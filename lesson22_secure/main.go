package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	orginalText := "123rasulQq"
	fmt.Println(orginalText)
	key := []byte("this's secret key.enough 32 bits")
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	keyString := hex.EncodeToString(key)

	fmt.Println("Encrypting....")
	cryptoText := encrypt(keyString, orginalText)
	fmt.Println(cryptoText)

	fmt.Println("Decrypting.....")
	// encrypt base64 crypto to original value
	text := decrypt(keyString, cryptoText)
	fmt.Println(text)

}

func encrypt(keyString string, StringToEncrypt string) (encryptedString string) {
	// convert key to bites

	key, _ := hex.DecodeString(keyString)
	plainText := []byte(StringToEncrypt)

	// Create a new Cipher Block from the key

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// The IV needs to be unique , but not secure. Therefore, it's common to
	//	// include it at the beginning of the ciphertext.

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err.Error())
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	// convert to base64
	return base64.URLEncoding.EncodeToString(cipherText)
}

func decrypt(keyString string, StringToDecrypt string) string {
	key, _ := hex.DecodeString(keyString)
	chipherText, _ := base64.URLEncoding.DecodeString(StringToDecrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	if len(chipherText) < aes.BlockSize {
		panic("cipher text to shoort")

	}
	iv := chipherText[:aes.BlockSize]
	chipherText = chipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(chipherText, chipherText)

	return fmt.Sprintf("%s", chipherText)

}
