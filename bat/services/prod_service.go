package services

import (
"context"
"gomicro-grpc/bat/datamodels"
"strconv"
)

//测试方法
func newProd(id int32,pname string) *datamodels.ProdModel {
	return &datamodels.ProdModel{ProdID: id,ProdName:pname}
}

type ProdService struct {

}

func(*ProdService) GetProdsList(ctx context.Context,in *datamodels.ProdsRequest,res *datamodels.ProdListResponse) error {
	models := make([]*datamodels.ProdModel,0)
	var i int32
	for i = 0; i < in.Size; i ++ {
		models = append(models,newProd(100+i,"prodname"+strconv.Itoa(100+int(i))))
	}
	res.Data = models
	return nil
}