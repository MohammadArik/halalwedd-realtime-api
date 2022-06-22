package main

import (
	"context"
	"log"

	serverConnectionService "github.com/MohammadArik/halalwedd/realtime-api/serverConnection"
)

type serverConnectionHandler struct {
	serverConnectionService.UnimplementedServerCheckingServer
}

func (serverConnectionHandler) PingServer(context.Context, *serverConnectionService.CheckReq) (*serverConnectionService.CheckRes, error) {
	res := &serverConnectionService.CheckRes{
		Pong: "PONG",
	}
	return res, nil
}

func (serverConnectionHandler) VerifyServer(_ context.Context, req *serverConnectionService.AuthReq) (*serverConnectionService.AuthRes, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error recovered:", r)
		}
	}()

	res := &serverConnectionService.AuthRes{
		Result: decryptCipher(req.Challenge),
	}
	return res, nil
}

func (serverConnectionHandler) DataUpdate(_ context.Context, req *serverConnectionService.ServerInfoUpdate) (*serverConnectionService.DataUpdateRes, error) {
	return nil, nil
}
