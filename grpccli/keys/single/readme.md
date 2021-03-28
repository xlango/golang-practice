openssl自签单向证书

    OpenSSL> genrsa -des3 -out server.key 2048
    OpenSSL> req -new -key server.key -out server.csr
    Common Name：域名
    OpenSSL> rsa -in server.key -out server_no_pwd.key //去除密码
    OpenSSL> x509 -req -days 365 -in server.csr -signkey server_no_pwd.key -out server.crt
    所需文件：server_no_pwd.key、server.crt