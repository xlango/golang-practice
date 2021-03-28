package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

//获取服务端双向证书cert信息
func GetServerCreds() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates:                []tls.Certificate{cert},
		ClientAuth:                  tls.RequireAndVerifyClientCert,
		ClientCAs:                   certPool,
	})

	return creds
}

//获取服务端双向证书cert信息
func GetClientCreds() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("cert/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates:                []tls.Certificate{cert},
		RootCAs:                     certPool,
		ServerName:                  "localhost",
	})

	return creds
}