package main

import (
	"context"
	"log"

	serverManagingService "github.com/MohammadArik/halalwedd/realtime-api/serverManagingService"
)

type handler_edgeServerConnectivityAPI struct {
	serverManagingService.UnimplementedEdgeServerConnectivityServer
}

func (handler_edgeServerConnectivityAPI) PingServer(context.Context, *serverManagingService.PingReq) (*serverManagingService.PingRes, error) {
	res := &serverManagingService.PingRes{
		Pong: "PONG",
	}
	return res, nil
}

func (handler_edgeServerConnectivityAPI) VerifyServer(_ context.Context, req *serverManagingService.VerifyReq) (*serverManagingService.VerifyRes, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error recovered:", r)
		}
	}()

	res := &serverManagingService.VerifyRes{
		Result: decryptCipher(req.Challenge),
	}
	return res, nil
}

func (handler_edgeServerConnectivityAPI) DataUpdate(_ context.Context, req *serverManagingService.DataUpdateReq) (*serverManagingService.DataUpdateRes, error) {
	if req.Type == serverManagingService.UPDATE_TYPE_New {
		addNewServer(req)
	} else if req.Type == serverManagingService.UPDATE_TYPE_Change {
		changeServer(req)
	} else {
		removeServer(req)
	}

	return &serverManagingService.DataUpdateRes{
		Message: "Success",
	}, nil
}
