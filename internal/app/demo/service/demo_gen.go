// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/demo"
	"github.com/tiger1103/gfast/v3/internal/app/demo/model"
)

type (
	IDemoGen interface {
		List(ctx context.Context, req *demo.DemoGenSearchReq) (listRes *demo.DemoGenSearchRes, err error)
		GetById(ctx context.Context, id uint) (res *model.DemoGenInfoRes, err error)
		Add(ctx context.Context, req *demo.DemoGenAddReq) (err error)
		Edit(ctx context.Context, req *demo.DemoGenEditReq) (err error)
		Delete(ctx context.Context, ids []uint) (err error)
	}
)

var (
	localDemoGen IDemoGen
)

func DemoGen() IDemoGen {
	if localDemoGen == nil {
		panic("implement not found for interface IDemoGen, forgot register?")
	}
	return localDemoGen
}

func RegisterDemoGen(i IDemoGen) {
	localDemoGen = i
}
