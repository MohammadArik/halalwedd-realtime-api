package main

import "log"

// High-level function for critical errors
func panicOnErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
