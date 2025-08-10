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
	IDemoGen16 interface {
		List(ctx context.Context, req *demo16.DemoGen16SearchReq) (listRes *demo16.DemoGen16SearchRes, err error)
		GetById(ctx context.Context, id uint) (res *model.DemoGen16InfoRes, err error)
		Add(ctx context.Context, req *demo16.DemoGen16AddReq) (err error)
		Edit(ctx context.Context, req *demo16.DemoGen16EditReq) (err error)
		Delete(ctx context.Context, ids []uint) (err error)
	}
)

var (
	localDemoGen16 IDemoGen16
)

func DemoGen16() IDemoGen16 {
	if localDemoGen16 == nil {
		panic("implement not found for interface IDemoGen16, forgot register?")
	}
	return localDemoGen16
}

func RegisterDemoGen16(i IDemoGen16) {
	localDemoGen16 = i
}
