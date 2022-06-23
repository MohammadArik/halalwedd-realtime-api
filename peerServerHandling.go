package main

import (
	newServerHandlingService "github.com/MohammadArik/halalwedd/realtime-api/newServerHandling"
	serverConnectionService "github.com/MohammadArik/halalwedd/realtime-api/serverConnection"
)

//! Structure plan

// The list of all peer realtime servers
var peerServers *[]newServerHandlingService.StoredServerData

// The function to add new servers and execute necessary steps for that
func addNewServer(updateInfo *serverConnectionService.ServerInfoUpdate) {
	//! Alogrithm to be added later
}

// The function to change a server and execute necessary steps for that
func changeServer(updateInfo *serverConnectionService.ServerInfoUpdate) {
	//! Alogrithm to be added later
}

// The function to remove a server and execute necessary steps for that
func removeServer(updateInfo *serverConnectionService.ServerInfoUpdate) {
	//! Alogrithm to be added later
}
