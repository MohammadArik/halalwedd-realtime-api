package main

import (
	"encoding/hex"
	"fmt"
)

func decryptCipher(challengeText string) string {
	var cipher, err = hex.DecodeString(challengeText)
	panicOnErr(err)
	var iv = []byte{}
	var ciphertext []byte

	for i := 0; i < len(cipher); i++ {
		if i < 12 {
			ciphertext = append(ciphertext, cipher[i])
		} else if i >= 12 && i < 24 {
			iv = append(iv, cipher[i])
		} else {
			ciphertext = append(ciphertext, cipher[i])
		}
	}

	plaintext, err := aesClass.Open(nil, iv, ciphertext, nil)
	panicOnErr(err)

	return (string(plaintext) + "-" + fmt.Sprint(processID))
}
