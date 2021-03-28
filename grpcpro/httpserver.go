package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"grpcpro/helper"
	"grpcpro/services"
	"net/http"
)

func main() {
	gwmux := runtime.NewServeMux()
	opt :=[]grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}
	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:8081", opt)
	if err != nil {
		panic(err)
	}

	httpServer := &http.Server{
		Addr:    ":8082",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}
