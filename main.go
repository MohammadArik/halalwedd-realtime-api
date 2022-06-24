package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"net"
	"os"

	serverConnectionService "github.com/MohammadArik/halalwedd/realtime-api/serverConnection"
	"google.golang.org/grpc"
)

// ** Global variables
// The struct for executing aes tasks
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
	// 1. Initializing the connection for the server
	verificationServerListener, err := net.Listen("tcp", ":3500")
	// 2. Initialize the server-handler
	verificationServerHandler := &managingServerConnectionHandler{}
	// 3. Registering the server
	verificationServerOpts := []grpc.ServerOption{}
	verificationServer := grpc.NewServer(verificationServerOpts...)

	serverConnectionService.RegisterServerCheckingServer(verificationServer, verificationServerHandler)
	// 4. Starting the server to listen
	verificationServerErrorChan := make(chan error)
	go func() {
		verificationServerErrorChan <- verificationServer.Serve(verificationServerListener)
	}()

	//* Calling the manager server to publish the server

	//* Initialize the Peer-Servers Manager

	//* Initializing Realtime API

	//* Starting Websocket Server

	//* The main-thread blocking select to listen for errors
	select {
		case verError := <- verificationServerErrorChan:
			log.Println("Verification server error:", verError)

	}

}
