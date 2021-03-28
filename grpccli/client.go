package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpccli/helper"
	"grpccli/services"
)

func main() {
	//creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "xyl")
	//if err != nil {
	//	panic(err)
	//}

	conn, err := grpc.Dial(":8081",grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	prodServiceClient := services.NewProdServiceClient(conn)

	resp, err := prodServiceClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 1,ProdArea:1})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.ProdStock)

	resps, err := prodServiceClient.GetProdStocks(context.Background(),&services.QuerySize{Size:3})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v",resps.ProdList)
}
