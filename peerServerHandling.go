package main

import (
	serverManagingService "github.com/MohammadArik/halalwedd/realtime-api/serverManaging"
)

//! Structure plan

// The list of all peer realtime servers
var peerServers *[]serverManagingService.ServerData

// The function to add new servers and execute necessary steps for that
func addNewServer(updateInfo *serverManagingService.DataUpdateReq) {
	//! Alogrithm to be added later
}

// The function to change a server and execute necessary steps for that
func changeServer(updateInfo *serverManagingService.DataUpdateReq) {
	//! Alogrithm to be added later
}

// The function to remove a server and execute necessary steps for that
func removeServer(updateInfo *serverManagingService.DataUpdateReq) {
	//! Alogrithm to be added later
}
