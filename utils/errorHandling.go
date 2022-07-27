package utils

import (
	"log"
)

// High-level function for critical errors
func PanicOnErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

// High-level function for non-critical errors
func LogOnErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
