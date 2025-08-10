// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/demo16"
	"github.com/tiger1103/gfast/v3/internal/app/demo16/model"
)

type (
	IDemoGenClass16 interface {
		List(ctx context.Context, req *demo16.DemoGenClass16SearchReq) (listRes *demo16.DemoGenClass16SearchRes, err error)
		GetById(ctx context.Context, id uint) (res *model.DemoGenClass16InfoRes, err error)
		Add(ctx context.Context, req *demo16.DemoGenClass16AddReq) (err error)
		Edit(ctx context.Context, req *demo16.DemoGenClass16EditReq) (err error)
		Delete(ctx context.Context, ids []uint) (err error)
	}
)

var (
	localDemoGenClass16 IDemoGenClass16
)

func DemoGenClass16() IDemoGenClass16 {
	if localDemoGenClass16 == nil {
		panic("implement not found for interface IDemoGenClass16, forgot register?")
	}
	return localDemoGenClass16
}

func RegisterDemoGenClass16(i IDemoGenClass16) {
	localDemoGenClass16 = i
}
