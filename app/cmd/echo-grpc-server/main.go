package main

import (
	"app/api/echo"
	"app/idl/echo/echov1"
	"app/pkg/server"
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
)

var sigs chan os.Signal
var serveAddr = ":9000"

func init() {
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
}

func main() {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.RunGRPCServer(ctx, &server.GRPCOptions{
			ServeAddr: serveAddr,
			Register: func(s *grpc.Server) {
				echov1.RegisterEchoAPIServer(s, &echo.API{})
			},
		})
	}()
	log.Printf("Running echo HTTP server at " + serveAddr)

	<-sigs
	log.Println("Shutting down servers...")
}
