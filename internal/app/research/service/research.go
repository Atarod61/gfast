// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/research"
	"github.com/tiger1103/gfast/v3/internal/app/research/model"
)

type (
	IResearch interface {
		List(ctx context.Context, req *research.ResearchSearchReq) (listRes *research.ResearchSearchRes, err error)
		GetById(ctx context.Context, id uint) (res *model.ResearchInfoRes, err error)
		Add(ctx context.Context, req *research.ResearchAddReq) (err error)
		Edit(ctx context.Context, req *research.ResearchEditReq) (err error)
		Delete(ctx context.Context, ids []uint) (err error)
	}
)

var (
	localResearch IResearch
)

func Research() IResearch {
	if localResearch == nil {
		panic("implement not found for interface IResearch, forgot register?")
	}
	return localResearch
}

func RegisterResearch(i IResearch) {
	localResearch = i
}
