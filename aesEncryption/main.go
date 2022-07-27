package aesEncryption

import (
	"crypto/aes"
	"crypto/cipher"
	cryptoRand "crypto/rand"
	"encoding/hex"
	"io"
	"os"

	"github.com/MohammadArik/halalwedd/realtime-api/utils"
)

type AES struct {
	aead *cipher.AEAD
}

func InitiateAES(keyFilePath string) *AES {
	//* Load AES keys and initiate AES Block
	keyFileData, err := os.ReadFile(keyFilePath)
	utils.PanicOnErr(err)
	key, err := hex.DecodeString(string(keyFileData))
	utils.PanicOnErr(err)
	block, err := aes.NewCipher(key)
	utils.PanicOnErr(err)
	aesGCM, err := cipher.NewGCM(block)
	newStruct := AES{
		aead: &aesGCM,
	}

	return &newStruct
}

func (class *AES) Encrypt(plaintext string) (result string) {
	// 1. Generating the Initializing Vector or Nonce
	var iv = make([]byte, 12)
	if _, err := io.ReadFull(cryptoRand.Reader, iv); err != nil {
		panic(err.Error())
	}

	// 2. Creating the ciphertext
	var ciphertext = (*class.aead).Seal(nil, iv, []byte(plaintext), nil)

	// 3. Encapsulating the iv with the ciphertext
	var cipher = []byte{}
	for i := 0; i < (len(ciphertext) + 12); i++ {
		if i < 12 {
			cipher = append(cipher, ciphertext[i])
		} else if i >= 12 && i < 24 {
			cipher = append(cipher, iv[i-12])
		} else {
			cipher = append(cipher, ciphertext[i-12])
		}
	}

	// 4. Convert to hex
	result = hex.EncodeToString(cipher)

	return
}

func (class *AES) Decrypt(cipher string) (plainText string) {
	cipherByte, err := hex.DecodeString(cipher)
	utils.PanicOnErr(err)
	var iv []byte
	var ciphertext []byte

	for i := 0; i < len(cipherByte); i++ {
		if i < 12 {
			ciphertext = append(ciphertext, cipherByte[i])
		} else if i >= 12 && i < 24 {
			iv = append(iv, cipherByte[i])
		} else {
			ciphertext = append(ciphertext, cipherByte[i])
		}
	}

	plainTextByte, err := (*class.aead).Open(nil, iv, ciphertext, nil)
	utils.PanicOnErr(err)

	return string(plainTextByte)
}
