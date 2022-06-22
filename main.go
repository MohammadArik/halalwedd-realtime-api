package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"os"
)

var aesClass cipher.AEAD

func main() {
	//* Log initial data
	log.Println("Server Initialization Started")
	log.Println("Server PID:", os.Getpid())
	//* Load AES keys and initiate AES class
	keyFileData, err := os.ReadFile("./keys/key.dat")
	panicOnErr(err)
	key, err := hex.DecodeString(string(keyFileData))
	panicOnErr(err)
	block, err := aes.NewCipher(key)
	panicOnErr(err)
	aesClass, err = cipher.NewGCM(block)
	panicOnErr(err)

	//* Initialize the verification server

	//* Calling the manager server to publish the server

}
