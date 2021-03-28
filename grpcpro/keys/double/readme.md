openssl自签双向证书

    #CA证书
    OpenSSL> genrsa -out ca.key 2048
    OpenSSL> req -new -x509 -days 3650 -key ca.key -out ca.pem
    
    #服务端证书
    OpenSSL> genrsa -out server.key 2048
    OpenSSL> req -new -key server.key -out server.csr
    Common Name ： localhost
    OpenSSL> x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem
    
    #客户端证书
    OpenSSL> ecparam -genkey -name secp384r1 -out client.key
    OpenSSL> req -new -key client.key -out client.csr
    OpenSSL> x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem