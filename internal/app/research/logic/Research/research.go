package Research

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/research"
	"github.com/tiger1103/gfast/v3/internal/app/research/dao"
	"github.com/tiger1103/gfast/v3/internal/app/research/model"
	"github.com/tiger1103/gfast/v3/internal/app/research/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/research/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	systemService "github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/libUtils"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterResearch(New())
}

func New() *sResearch {
	return &sResearch{}
}

type sResearch struct{}

func (s *sResearch) List(ctx context.Context, req *research.ResearchSearchReq) (listRes *research.ResearchSearchRes, err error) {
	listRes = new(research.ResearchSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Research.Ctx(ctx).WithAll()
		if req.Type != "" {
			m = m.Where(dao.Research.Columns().Type+" like ?", "%"+req.Type+"%")
		}
		if req.Subject != "" {
			m = m.Where(dao.Research.Columns().Subject+" = ?", req.Subject)
		}
		if req.Group != "" {
			m = m.Where(dao.Research.Columns().Group+" like ?", "%"+req.Group+"%")
		}
		if req.Publisher != "" {
			m = m.Where(dao.Research.Columns().Publisher+" like ?", "%"+req.Publisher+"%")

		}
		if req.Date != "" {
			m = m.Where(dao.Research.Columns().Date+" = ?", gconv.Time(req.Date))
		}
		if req.Confirmed1By != "" {
			m = m.Where(dao.Research.Columns().Confirmed1By+" = ?", req.Confirmed1By)
		}
		if req.Confirmed1At != "" {
			m = m.Where(dao.Research.Columns().Confirmed1At+" = ?", gconv.Time(req.Confirmed1At))
		}
		if req.Confirmed2By != "" {
			m = m.Where(dao.Research.Columns().Confirmed2By+" = ?", req.Confirmed2By)
		}
		if req.Confirmed2At != "" {
			m = m.Where(dao.Research.Columns().Confirmed2At+" = ?", gconv.Time(req.Confirmed2At))
		}
		if req.Score != "" {
			m = m.Where(dao.Research.Columns().Score+" like ?", "%"+req.Score+"%")
		}
		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "Failed to obtain total number of rows")
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
		var res []*model.ResearchInfoRes
		err = m.Fields(research.ResearchSearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "Failed to obtain data")
		listRes.List = make([]*model.ResearchListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.ResearchListRes{
				Id:           v.Id,
				Type:         v.Type,
				LinkedType:   v.LinkedType,
				Subject:      v.Subject,
				Group:        v.Group,
				LinkedGroup:  v.LinkedGroup,
				Publisher:    v.Publisher,
				Date:         v.Date,
				CreatedAt:    v.CreatedAt,
				CreatedBy:    v.CreatedBy,
				Confirmed1By: v.Confirmed1By,
				Confirmed1At: v.Confirmed1At,
				Confirmed2By: v.Confirmed2By,
				Confirmed2At: v.Confirmed2At,
				Score:        v.Score,
			}
		}
	})
	return
}

func (s *sResearch) GetById(ctx context.Context, id uint) (res *model.ResearchInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Research.Ctx(ctx).WithAll().Where(dao.Research.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "Failed to obtain information")
	})
	return
}

func (s *sResearch) Add(ctx context.Context, req *research.ResearchAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//demoCate := ""
		//req.DemoCate.FilterEmpty()
		//if !req.DemoCate.IsEmpty(){
		//demoCate=req.DemoCate.Join(",")
		//}
		for _, obj := range req.File {
			obj.Url, err = libUtils.GetFilesPath(ctx, obj.Url)
			liberr.ErrIsNil(ctx, err)
		}
		for _, obj := range req.Image {
			obj.Url, err = libUtils.GetFilesPath(ctx, obj.Url)
			liberr.ErrIsNil(ctx, err)
		}
		_, err = dao.Research.Ctx(ctx).Insert(do.Research{
			Type:         req.Type,
			Subject:      req.Subject,
			Group:        req.Group,
			Publisher:    req.Publisher,
			Date:         req.Date,
			File:         req.File,
			Image:        req.Image,
			Confirmed1By: req.Confirmed1By,
			Confirmed1At: req.Confirmed1At,
			Confirmed2By: req.Confirmed2By,
			Confirmed2At: req.Confirmed2At,
			Score:        req.Score,
			CreatedBy:    systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "Add failed")
	})
	return
}

func (s *sResearch) Edit(ctx context.Context, req *research.ResearchEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		//demoCate := ""
		//req.DemoCate.FilterEmpty()
		//if !req.DemoCate.IsEmpty(){
		//demoCate=req.DemoCate.Join(",")
		//}
		for _, obj := range req.File {
			obj.Url, err = libUtils.GetFilesPath(ctx, obj.Url)
			liberr.ErrIsNil(ctx, err)
		}
		for _, obj := range req.Image {
			obj.Url, err = libUtils.GetFilesPath(ctx, obj.Url)
			liberr.ErrIsNil(ctx, err)
		}
		_, err = dao.Research.Ctx(ctx).WherePri(req.Id).Update(do.Research{
			Type:         req.Type,
			Subject:      req.Subject,
			Group:        req.Group,
			Publisher:    req.Publisher,
			Date:         req.Date,
			File:         req.File,
			Image:        req.Image,
			Confirmed1By: req.Confirmed1By,
			Confirmed1At: req.Confirmed1At,
			Confirmed2By: req.Confirmed2By,
			Confirmed2At: req.Confirmed2At,
			Score:        req.Score,
			UpdatedBy:    systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "Modification failed")
	})
	return
}

func (s *sResearch) Delete(ctx context.Context, ids []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Research.Ctx(ctx).Delete(dao.Research.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "Deletion failed")
	})
	return
}
