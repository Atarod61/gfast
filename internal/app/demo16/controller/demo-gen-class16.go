package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/demo16"
	"github.com/tiger1103/gfast/v3/internal/app/demo16/service"
	systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
)

type demoGenClass16Controller struct {
	systemController.BaseController
}

var DemoGenClass16 = new(demoGenClass16Controller)

// List 列表
func (c *demoGenClass16Controller) List(ctx context.Context, req *demo16.DemoGenClass16SearchReq) (res *demo16.DemoGenClass16SearchRes, err error) {
	res, err = service.DemoGenClass16().List(ctx, req)
	return
}

// Get 获取分类信息
func (c *demoGenClass16Controller) Get(ctx context.Context, req *demo16.DemoGenClass16GetReq) (res *demo16.DemoGenClass16GetRes, err error) {
	res = new(demo16.DemoGenClass16GetRes)
	res.DemoGenClass16InfoRes, err = service.DemoGenClass16().GetById(ctx, req.Id)
	return
}

// Add 添加分类信息
func (c *demoGenClass16Controller) Add(ctx context.Context, req *demo16.DemoGenClass16AddReq) (res *demo16.DemoGenClass16AddRes, err error) {
	err = service.DemoGenClass16().Add(ctx, req)
	return
}

// Edit 修改分类信息
func (c *demoGenClass16Controller) Edit(ctx context.Context, req *demo16.DemoGenClass16EditReq) (res *demo16.DemoGenClass16EditRes, err error) {
	err = service.DemoGenClass16().Edit(ctx, req)
	return
}

// Delete 删除分类信息
func (c *demoGenClass16Controller) Delete(ctx context.Context, req *demo16.DemoGenClass16DeleteReq) (res *demo16.DemoGenClass16DeleteRes, err error) {
	err = service.DemoGenClass16().Delete(ctx, req.Ids)
	return
}
