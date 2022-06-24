package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"log"
	"net"
	"os"

	serverConnectionService "github.com/MohammadArik/halalwedd/realtime-api/serverConnection"
	serverHandlingService "github.com/MohammadArik/halalwedd/realtime-api/serverHandling"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ** Global variables
// The struct for executing aes tasks
var aesClass cipher.AEAD

// Server credentials
var internalID int
var publicIP string
var privateIP string

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

	//* Gather the instances Public and Private IP address and IP address of managing server
	// 1. Get the home directory
	homeDir, err := os.UserHomeDir()
	panicOnErr(err)
	// 2. Get the ip info file
	addressFile, err := os.ReadFile(homeDir + "/ip_info.json")
	panicOnErr(err)
	// 3. Extract the data
	addressFileData := map[string]string{}
	err = json.Unmarshal(addressFile, &addressFileData)
	panicOnErr(err)
	publicIP = addressFileData["publicAddress"]
	privateIP = addressFileData["privateAddress"]

	//* Calling the manager server to publish the server
	// 1. Connect to the managing server
	conn, err := grpc.Dial(addressFileData["managingServerAdress"]+":6491", grpc.WithTransportCredentials(insecure.NewCredentials()))

	managingServerClient := serverHandlingService.NewServerHandlingClient(conn)

	res, err := managingServerClient.PublishServer(context.Background(), &serverHandlingService.ServerInfo{
		Type: serverHandlingService.SERVER_TYPE_REALTIME,
	})

	//* Initialize the Peer-Servers Manager

	//* Initializing Realtime API

	//* Starting Websocket Server

	//* The main-thread blocking select to listen for errors
	select {
	case verError := <-verificationServerErrorChan:
		log.Println("Verification server error:", verError)

	}

}
