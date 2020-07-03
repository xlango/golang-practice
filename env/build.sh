CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
go build -o envtest -a envtest.go