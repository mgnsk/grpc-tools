package main

import (
	"app/api/echo"
	"app/idl/echo/echov1"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
)

var sigs chan os.Signal

func init() {
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
}

func main() {
	// Defer based quit mechanism.
	defer glog.Flush()

	// Start grpc server
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		glog.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	echov1.RegisterEchoAPIServer(s, &echo.API{})
	defer s.Stop()

	go func() {
		if err := s.Serve(lis); err != nil {
			glog.Fatalf("server crashed: %v", err)
		}
	}()

	fmt.Println("Running gRPC server at localhost:9000")

	// Setup http gateway
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// TODO ServeMux opts
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Connect to the gRPC server.
	err = echov1.RegisterEchoAPIHandlerFromEndpoint(ctx, mux, "localhost:9000", opts)
	if err != nil {
		glog.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(cors.Default().Handler)
	router.Handle("/*", mux)

	srv := &http.Server{
		Addr:         ":8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	defer srv.Shutdown(ctx)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				glog.Fatal(err)
			}
			log.Println("http server closed")
		}
	}()

	fmt.Println("Running HTTP server at localhost:8000")
	<-sigs
}
