package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/research"
	"github.com/tiger1103/gfast/v3/internal/app/research/service"
	systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
)

type ResearchWrittenWorksController struct {
	systemController.BaseController
}

var ResearchWrittenWorks = new(ResearchWrittenWorksController)

// List 列表
func (c *ResearchWrittenWorksController) List(ctx context.Context, req *research.ResearchWrittenWorksSearchReq) (res *research.ResearchWrittenWorksSearchRes, err error) {
	res, err = service.ResearchWrittenWorks().List(ctx, req)
	return
}

// Get 获取分类信息
func (c *ResearchWrittenWorksController) Get(ctx context.Context, req *research.ResearchWrittenWorksGetReq) (res *research.ResearchWrittenWorksGetRes, err error) {
	res = new(research.ResearchWrittenWorksGetRes)
	res.ResearchWrittenWorksInfoRes, err = service.ResearchWrittenWorks().GetById(ctx, req.Id)
	return
}

// Add 添加分类信息
func (c *ResearchWrittenWorksController) Add(ctx context.Context, req *research.ResearchWrittenWorksAddReq) (res *research.ResearchWrittenWorksAddRes, err error) {
	err = service.ResearchWrittenWorks().Add(ctx, req)
	return
}

// Edit 修改分类信息
func (c *ResearchWrittenWorksController) Edit(ctx context.Context, req *research.ResearchWrittenWorksEditReq) (res *research.ResearchWrittenWorksEditRes, err error) {
	err = service.ResearchWrittenWorks().Edit(ctx, req)
	return
}

// Delete 删除分类信息
func (c *ResearchWrittenWorksController) Delete(ctx context.Context, req *research.ResearchWrittenWorksDeleteReq) (res *research.ResearchWrittenWorksDeleteRes, err error) {
	err = service.ResearchWrittenWorks().Delete(ctx, req.Ids)
	return
}
