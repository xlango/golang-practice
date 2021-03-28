package services

import (
	"context"
	"fmt"
)

type ProdService struct {

}

func (this *ProdService) GetProdStocks(cxt context.Context,req *QuerySize) (*ProdResponseList, error) {
	list := []*ProdResponse{
		&ProdResponse{ProdStock:1},
		&ProdResponse{ProdStock:2},
		&ProdResponse{ProdStock:3},
		&ProdResponse{ProdStock:4},
		&ProdResponse{ProdStock:5},
	}
	return &ProdResponseList{ProdList:list[:req.Size]}, nil
}

func (this *ProdService) GetProdStock(cxt context.Context,req *ProdRequest) (*ProdResponse, error) {
	fmt.Println(req.ProdId)
	fmt.Println(req.ProdArea)
	if req.ProdArea == ProdArea_A {
		return &ProdResponse{ProdStock:0}, nil
	}
	if req.ProdArea == ProdArea_B {
		return &ProdResponse{ProdStock:1}, nil
	}
	if req.ProdArea == ProdArea_C {
		return &ProdResponse{ProdStock:2}, nil
	}
	return &ProdResponse{ProdStock:20},nil
}
