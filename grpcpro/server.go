package main

import (
	"google.golang.org/grpc"
	"grpcpro/helper"
	"grpcpro/services"
	"net"
)

func main() {
	//creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_pwd.key")
	//if err != nil {
	//	panic(err)
	//}

	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))

	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	rpcServer.ServeHTTP(writer,request)
	//})
	//httpServer := &http.Server{
	//	Addr:              ":8081",
	//	Handler:           mux,
	//	TLSConfig:         nil,
	//	ReadTimeout:       0,
	//	ReadHeaderTimeout: 0,
	//	WriteTimeout:      0,
	//	IdleTimeout:       0,
	//	MaxHeaderBytes:    0,
	//	TLSNextProto:      nil,
	//	ConnState:         nil,
	//	ErrorLog:          nil,
	//	BaseContext:       nil,
	//	ConnContext:       nil,
	//}
	//httpServer.ListenAndServeTLS("keys/server.crt", "keys/server_no_pwd.key")
}
