// DO NOT EDIT. This file is generated by app/cmd/tools/generator/main.go

package main

import (
	"app/idl/echo/echov1"
	"app/pkg/server"
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
)

var sigs chan os.Signal

func init() {
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
}

func main() {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// HTTP endpoint
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.RunGatewayServer(ctx, &server.GatewayOptions{
			ServeAddr: os.Getenv("HTTP_BIND_ADDR"),
			GRPCAddr:  os.Getenv("GRPC_DIAL_ADDR"),
			DialOpts:  []grpc.DialOption{grpc.WithInsecure()},
			Register:  echov1.RegisterEchoAPIHandlerFromEndpoint,
		})
	}()

	<-sigs
}