// DO NOT EDIT. This file is generated.

package main

import (
	"app/api/raspi"
	"app/idl/raspi/raspiv1"
	"app/pkg/server"
	"context"
	"fmt"
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

	// gRPC endpoint
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.RunGRPCServer(ctx, server.GRPCOptions{
			ServeAddr: os.Getenv("GRPC_BIND_ADDR"),
			Register: func(s *grpc.Server) {
				raspiv1.RegisterRaspiAPIServer(s, &raspi.API{})
			},
		})
	}()

	fmt.Println("Running raspi gRPC server at", os.Getenv("GRPC_BIND_ADDR"))

	<-sigs
}
