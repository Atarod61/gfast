package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/demo16"
	"github.com/tiger1103/gfast/v3/internal/app/demo16/service"
	systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
)

type demoGen16Controller struct {
	systemController.BaseController
}

var DemoGen16 = new(demoGen16Controller)

// List 列表
func (c *demoGen16Controller) List(ctx context.Context, req *demo16.DemoGen16SearchReq) (res *demo16.DemoGen16SearchRes, err error) {
	res, err = service.DemoGen16().List(ctx, req)
	return
}

// Get 获取代码生成测试
func (c *demoGen16Controller) Get(ctx context.Context, req *demo16.DemoGen16GetReq) (res *demo16.DemoGen16GetRes, err error) {
	res = new(demo16.DemoGen16GetRes)
	res.DemoGen16InfoRes, err = service.DemoGen16().GetById(ctx, req.Id)
	return
}

// Add 添加代码生成测试
func (c *demoGen16Controller) Add(ctx context.Context, req *demo16.DemoGen16AddReq) (res *demo16.DemoGen16AddRes, err error) {
	err = service.DemoGen16().Add(ctx, req)
	return
}

// Edit 修改代码生成测试
func (c *demoGen16Controller) Edit(ctx context.Context, req *demo16.DemoGen16EditReq) (res *demo16.DemoGen16EditRes, err error) {
	err = service.DemoGen16().Edit(ctx, req)
	return
}

// Delete 删除代码生成测试
func (c *demoGen16Controller) Delete(ctx context.Context, req *demo16.DemoGen16DeleteReq) (res *demo16.DemoGen16DeleteRes, err error) {
	err = service.DemoGen16().Delete(ctx, req.Ids)
	return
}
