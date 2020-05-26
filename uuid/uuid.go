package main

import (
	"fmt"
	//"github.com/gofrs/uuid"
	"github.com/satori/go.uuid"
)

func main() {
	tmpUuid, _ := uuid.NewV4()
	tmpUuid1, _ := uuid.NewV4()
	tmpUuid2, _ := uuid.NewV4()
	tmpUuid3, _ := uuid.NewV4()
	tmpUuid4, _ := uuid.NewV4()
	//genuid1 := strings.Replace(strings.TrimSpace(tmpUuid.String()), "-", "", -1)
	//genuid := strings.Replace(strings.TrimSpace(string(genuid1)), "\n", "", -1)
	fmt.Println(tmpUuid.String())
	fmt.Println(tmpUuid1.String())
	fmt.Println(tmpUuid2.String())
	fmt.Println(tmpUuid3.String())
	fmt.Println(tmpUuid4.String())
}
