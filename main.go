package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/MohammadArik/halalwedd/realtime-api/aesEncryption"
	serverManagingService "github.com/MohammadArik/halalwedd/realtime-api/serverManagingService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ** Global variables
// The struct for executing aes tasks
var aesHandler *aesEncryption.AES

// Server credentials
var internalID int
var publicIP string
var privateIP string
var processID int

func main() {
	//* Log initial data
	log.Println("Server Initialization Started")
	log.Println("Server PID:", os.Getpid())
	//* Load AES keys and initiate AES class
	aesHandler = aesEncryption.InitiateAES("./keys/serverManaging.key")

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

	//* Parse managing-server IP from the ip_info.json file
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
	processID = os.Getpid()

	//* Calling the manager server to publish the server
	// 1. Connect to the managing server
	log.Println("Connecting to Managing Server...")
	serverHandlingAPI_conn, err := grpc.Dial(addressFileData["managingServerAdress"]+":6491", grpc.WithTransportCredentials(insecure.NewCredentials()))

	serverHandlingAPI_client := serverManagingService.NewServerHandlingClient(serverHandlingAPI_conn)
	// 2. Make the request
	res, err := serverHandlingAPI_client.PublishServer(context.Background(), &serverManagingService.PublishServerReq{
		Type:      serverManagingService.SERVER_TYPE_REALTIME,
		Public_IP: publicIP,
	})
	panicOnErr(err)
	log.Println("Server published to managing-server!")
	log.Println("Response Message:", res.Message)
	log.Println("Internal ID:", res.InternalID)
	log.Println("Server Data:", res.ServerData)
	internalID = int(res.InternalID)
	privateIP = res.ServerData[res.InternalID].Private_IP

	//* Initialize the Peer-Servers Manager

	//* Initializing Realtime API

	//* Starting Websocket Server

	//* The listener for system interrupt or termination
	sysInterruptChan := make(chan os.Signal)

	//? In case of confusion,
	//? syscall.SIGINT refers to interrupt(ctrl+C)
	//? syscall.SIGTERM refers to termination(systemctl stop)
	signal.Notify(sysInterruptChan, syscall.SIGINT, syscall.SIGTERM)

	//* The main-thread blocking select to listen for errors
	select {
	case edgeServerConnectivityAPI_error := <-edgeServerConnectivityAPI_ErrorChan:
		serverHandlingAPI_client.NotifyExit(context.Background(), &serverManagingService.NotifyExitReq{
			Type:       serverManagingService.EXIT_TYPE_ERROR,
			SType:      serverManagingService.SERVER_TYPE_REALTIME,
			InternalID: res.InternalID,
			PublicIP:   publicIP,
		})
		log.Println("Verification server error:", edgeServerConnectivityAPI_error.Error())

	case sysInterrupt := <-sysInterruptChan:
		serverHandlingAPI_client.NotifyExit(context.Background(), &serverManagingService.NotifyExitReq{
			Type:       serverManagingService.EXIT_TYPE_ERROR,
			SType:      serverManagingService.SERVER_TYPE_REALTIME,
			InternalID: res.InternalID,
			PublicIP:   publicIP,
		})
		log.Println("Closing server due to system interruption. Signal:", sysInterrupt)
	}

}
