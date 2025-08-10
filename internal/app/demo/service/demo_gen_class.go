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
	IDemoGenClass interface {
		List(ctx context.Context, req *demo.DemoGenClassSearchReq) (listRes *demo.DemoGenClassSearchRes, err error)
		GetById(ctx context.Context, id uint) (res *model.DemoGenClassInfoRes, err error)
		Add(ctx context.Context, req *demo.DemoGenClassAddReq) (err error)
		Edit(ctx context.Context, req *demo.DemoGenClassEditReq) (err error)
		Delete(ctx context.Context, ids []uint) (err error)
	}
)

var (
	localDemoGenClass IDemoGenClass
)

func DemoGenClass() IDemoGenClass {
	if localDemoGenClass == nil {
		panic("implement not found for interface IDemoGenClass, forgot register?")
	}
	return localDemoGenClass
}

func RegisterDemoGenClass(i IDemoGenClass) {
	localDemoGenClass = i
}
