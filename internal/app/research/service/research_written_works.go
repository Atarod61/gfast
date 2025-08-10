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
	IResearchWrittenWorks interface {
		List(ctx context.Context, req *research.ResearchWrittenWorksSearchReq) (listRes *research.ResearchWrittenWorksSearchRes, err error)
		GetById(ctx context.Context, id uint) (res *model.ResearchWrittenWorksInfoRes, err error)
		Add(ctx context.Context, req *research.ResearchWrittenWorksAddReq) (err error)
		Edit(ctx context.Context, req *research.ResearchWrittenWorksEditReq) (err error)
		Delete(ctx context.Context, ids []uint) (err error)
	}
)

var (
	localResearchWrittenWorks IResearchWrittenWorks
)

func ResearchWrittenWorks() IResearchWrittenWorks {
	if localResearchWrittenWorks == nil {
		panic("implement not found for interface IResearchWrittenWorks, forgot register?")
	}
	return localResearchWrittenWorks
}

func RegisterResearchWrittenWorks(i IResearchWrittenWorks) {
	localResearchWrittenWorks = i
}
