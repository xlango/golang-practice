package main

import (
	"fmt"
	"os"
)

func main() {
	getenv := os.Getenv("TOPSEC_OLD_CLIENT_ID")
	fmt.Println(getenv)
}
