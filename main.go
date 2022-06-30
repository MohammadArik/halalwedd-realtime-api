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

	serverManagingService "github.com/MohammadArik/halalwedd/realtime-api/serverManagingService"
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
	edgeServerConnectivityAPI_listener, err := net.Listen("tcp", ":3500")
	// 2. Initialize the server-handler
	edgeServerConnectivityAPI_handler := &handler_edgeServerConnectivityAPI{}
	// 3. Registering the server
	edgeServerConnectivityAPI_opts := []grpc.ServerOption{}
	edgeServerConnectivityAPI := grpc.NewServer(edgeServerConnectivityAPI_opts...)

	serverManagingService.RegisterEdgeServerConnectivityServer(edgeServerConnectivityAPI, edgeServerConnectivityAPI_handler)
	// 4. Starting the server to listen
	edgeServerConnectivityAPI_ErrorChan := make(chan error)
	go func() {
		edgeServerConnectivityAPI_ErrorChan <- edgeServerConnectivityAPI.Serve(edgeServerConnectivityAPI_listener)
	}()
	log.Println("Verification Server Started")

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
	log.Println("Connecting to Managing Server...")
	serverHandlingAPI_conn, err := grpc.Dial(addressFileData["managingServerAdress"]+":6491", grpc.WithTransportCredentials(insecure.NewCredentials()))

	serverHandlingAPI_client := serverManagingService.NewServerHandlingClient(serverHandlingAPI_conn)
	// 2. Make the request
	res, err := serverHandlingAPI_client.PublishServer(context.Background(), &serverManagingService.PublishServerReq{
		Type:       serverManagingService.SERVER_TYPE_REALTIME,
		Public_IP:  publicIP,
		Private_IP: privateIP,
	})
	panicOnErr(err)
	log.Println("Server published to managing-server!")
	log.Println("Response Message:", res.Message)
	log.Println("Internal ID:", res.InternalID)
	log.Println("Server Data:", res.ServerData)
	internalID = int(res.InternalID)

	//* Initialize the Peer-Servers Manager

	//* Initializing Realtime API

	//* Starting Websocket Server

	//* The main-thread blocking select to listen for errors
	select {
	case edgeServerConnectivityAPI_error := <-edgeServerConnectivityAPI_ErrorChan:
		log.Println("Verification server error:", edgeServerConnectivityAPI_error.Error())
	}

}
