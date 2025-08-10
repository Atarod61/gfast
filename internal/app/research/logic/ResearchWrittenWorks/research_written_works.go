package ResearchWrittenWorks

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/research"
	"github.com/tiger1103/gfast/v3/internal/app/research/dao"
	"github.com/tiger1103/gfast/v3/internal/app/research/model"
	"github.com/tiger1103/gfast/v3/internal/app/research/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/research/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterResearchWrittenWorks(New())
}

func New() *sResearchWrittenWorks {
	return &sResearchWrittenWorks{}
}

type sResearchWrittenWorks struct{}

func (s *sResearchWrittenWorks) List(ctx context.Context, req *research.ResearchWrittenWorksSearchReq) (listRes *research.ResearchWrittenWorksSearchRes, err error) {
	listRes = new(research.ResearchWrittenWorksSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.ResearchWrittenWorks.Ctx(ctx).WithAll()
		if req.ClassName != "" {
			m = m.Where(dao.ResearchWrittenWorks.Columns().ClassName+" like ?", "%"+req.ClassName+"%")
		}
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "Failed to obtain the total number of rows")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "id asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.ResearchWrittenWorksInfoRes
		err = m.Fields(research.ResearchWrittenWorksSearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "Failed to obtain data")
		listRes.List = make([]*model.ResearchWrittenWorksListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.ResearchWrittenWorksListRes{
				Id:        v.Id,
				ClassName: v.ClassName,
			}
		}
	})
	return
}

func (s *sResearchWrittenWorks) GetById(ctx context.Context, id uint) (res *model.ResearchWrittenWorksInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.ResearchWrittenWorks.Ctx(ctx).WithAll().Where(dao.ResearchWrittenWorks.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "Failed to get information")
	})
	return
}

func (s *sResearchWrittenWorks) Add(ctx context.Context, req *research.ResearchWrittenWorksAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.ResearchWrittenWorks.Ctx(ctx).Insert(do.ResearchWrittenWorks{
			ClassName: req.ClassName,
		})
		liberr.ErrIsNil(ctx, err, "Add failed")
	})
	return
}

func (s *sResearchWrittenWorks) Edit(ctx context.Context, req *research.ResearchWrittenWorksEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.ResearchWrittenWorks.Ctx(ctx).WherePri(req.Id).Update(do.ResearchWrittenWorks{
			ClassName: req.ClassName,
		})
		liberr.ErrIsNil(ctx, err, "Modification failed")
	})
	return
}

func (s *sResearchWrittenWorks) Delete(ctx context.Context, ids []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.ResearchWrittenWorks.Ctx(ctx).Delete(dao.ResearchWrittenWorks.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "Deletion failed")
	})
	return
}
