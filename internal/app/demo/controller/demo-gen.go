package controller


import (
    "context"
    "github.com/tiger1103/gfast/v3/api/v1/demo"
    "github.com/tiger1103/gfast/v3/internal/app/demo/service"
    systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
)


type demoGenController struct {
    systemController.BaseController
}


var DemoGen = new(demoGenController)


// List 列表
func (c *demoGenController) List(ctx context.Context, req *demo.DemoGenSearchReq) (res *demo.DemoGenSearchRes, err error) {
    res, err = service.DemoGen().List(ctx, req)
    return
}


// Get 获取代码生成测试
func (c *demoGenController) Get(ctx context.Context, req *demo.DemoGenGetReq) (res *demo.DemoGenGetRes, err error) {
    res = new(demo.DemoGenGetRes)
    res.DemoGenInfoRes,err = service.DemoGen().GetById(ctx, req.Id)
    return
}


// Add 添加代码生成测试
func (c *demoGenController) Add(ctx context.Context, req *demo.DemoGenAddReq) (res *demo.DemoGenAddRes, err error) {
    err = service.DemoGen().Add(ctx, req)
    return
}


// Edit 修改代码生成测试
func (c *demoGenController) Edit(ctx context.Context, req *demo.DemoGenEditReq) (res *demo.DemoGenEditRes, err error) {
    err = service.DemoGen().Edit(ctx, req)
    return
}


// Delete 删除代码生成测试
func (c *demoGenController) Delete(ctx context.Context, req *demo.DemoGenDeleteReq) (res *demo.DemoGenDeleteRes, err error) {
    err = service.DemoGen().Delete(ctx, req.Ids)
    return
}