package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/research"
	"github.com/tiger1103/gfast/v3/internal/app/research/service"
	systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
)

type ResearchController struct {
	systemController.BaseController
}

var Research = new(ResearchController)

// List 列表
func (c *ResearchController) List(ctx context.Context, req *research.ResearchSearchReq) (res *research.ResearchSearchRes, err error) {
	res, err = service.Research().List(ctx, req)
	return
}

// Get 获取代码生成测试
func (c *ResearchController) Get(ctx context.Context, req *research.ResearchGetReq) (res *research.ResearchGetRes, err error) {
	res = new(research.ResearchGetRes)
	res.ResearchInfoRes, err = service.Research().GetById(ctx, req.Id)
	return
}

// Add 添加代码生成测试
func (c *ResearchController) Add(ctx context.Context, req *research.ResearchAddReq) (res *research.ResearchAddRes, err error) {
	err = service.Research().Add(ctx, req)
	return
}

// Edit 修改代码生成测试
func (c *ResearchController) Edit(ctx context.Context, req *research.ResearchEditReq) (res *research.ResearchEditRes, err error) {
	err = service.Research().Edit(ctx, req)
	return
}

// Delete 删除代码生成测试
func (c *ResearchController) Delete(ctx context.Context, req *research.ResearchDeleteReq) (res *research.ResearchDeleteRes, err error) {
	err = service.Research().Delete(ctx, req.Ids)
	return
}
