package main

import (
	"context"

	serverConnectionService "github.com/MohammadArik/halalwedd/realtime-api/serverConnection"
)

type serverConnectionHandler struct {
	serverConnectionService.UnimplementedServerCheckingServer
}

func (serverConnectionHandler) PingServer(ctx context.Context, _ *serverConnectionService.CheckReq) (*serverConnectionService.CheckRes, error) {
	res := &serverConnectionService.CheckRes{
		Pong: "PONG",
	}
	return res, nil
}

// VerifyServer(context.Context, *AuthReq) (*AuthRes, error)
// DataUpdate(context.Context, *ServerInfoUpdate) (*DataUpdateRes, error)
