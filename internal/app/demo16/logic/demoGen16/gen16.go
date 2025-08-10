package demoGen16


import (
    "context"
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/util/gconv"
    "github.com/tiger1103/gfast/v3/api/v1/demo16"
    "github.com/tiger1103/gfast/v3/internal/app/demo16/dao"
    "github.com/tiger1103/gfast/v3/internal/app/demo16/model"
    "github.com/tiger1103/gfast/v3/internal/app/demo16/model/do"
    "github.com/tiger1103/gfast/v3/internal/app/system/consts"
    systemService "github.com/tiger1103/gfast/v3/internal/app/system/service"
    "github.com/tiger1103/gfast/v3/library/libUtils"
    "github.com/tiger1103/gfast/v3/library/liberr"
	"github.com/tiger1103/gfast/v3/internal/app/demo16/service"
)

func init() {
    service.RegisterDemoGen16(New())
}

func New() *sDemoGen16 {
    return &sDemoGen16{}
}


type sDemoGen16 struct{}


func (s *sDemoGen16)List(ctx context.Context, req *demo16.DemoGen16SearchReq) (listRes *demo16.DemoGen16SearchRes, err error){
    listRes = new(demo16.DemoGen16SearchRes)
    err = g.Try(ctx, func(ctx context.Context) {
        m := dao.DemoGen16.Ctx(ctx).WithAll()
        if req.Type != "" {
            m = m.Where(dao.DemoGen16.Columns().Type+" like ?", "%"+req.Type+"%")
        }
		if req.Subject != "" {
            m = m.Where(dao.DemoGen16.Columns().Subject+" = ?", req.Subject)
        }
		if req.Group != "" {
            m = m.Where(dao.DemoGen16.Columns().Group+" like ?", "%"+req.Group+"%")
        }
		if req.Publisher != "" {
            m = m.Where(dao.DemoGen16.Columns().Publisher+" like ?", "%"+req.Publisher+"%")

        }
        if req.Date != "" {
            m = m.Where(dao.DemoGen16.Columns().Date+" = ?", gconv.Time(req.Date))
        }
		if req.Confirmed1By != "" {
            m = m.Where(dao.DemoGen16.Columns().Confirmed1By+" = ?", req.Confirmed1By)
        }
        if req.Confirmed1At != "" {
            m = m.Where(dao.DemoGen16.Columns().Confirmed1At+" = ?", gconv.Time(req.Confirmed1At))
        }
		if req.Confirmed2By != "" {
            m = m.Where(dao.DemoGen16.Columns().Confirmed2By+" = ?", req.Confirmed2By)
        }
        if req.Confirmed2At != "" {
            m = m.Where(dao.DemoGen16.Columns().Confirmed2At+" = ?", gconv.Time(req.Confirmed2At))
        }
		if req.Score != "" {
            m = m.Where(dao.DemoGen16.Columns().Score+" like ?", "%"+req.Score+"%")
        }
        listRes.Total, err = m.Count()
        liberr.ErrIsNil(ctx, err, "获取总行数失败")
        if req.PageNum == 0 {
            req.PageNum = 1
        }
        listRes.CurrentPage = req.PageNum
        if req.PageSize == 0 {
            req.PageSize = consts.PageSize
        }
        order:= "id asc"
        if req.OrderBy!=""{
            order = req.OrderBy
        }
        var res []*model.DemoGen16InfoRes
        err = m.Fields(demo16.DemoGen16SearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
        liberr.ErrIsNil(ctx, err, "获取数据失败")
        listRes.List = make([]*model.DemoGen16ListRes,len(res))
        for k,v:=range res{
            listRes.List[k] = &model.DemoGen16ListRes{
                Id : v.Id,
                Type : v.Type,
				LinkedType:v.LinkedType,
                Subject : v.Subject,
                Group : v.Group,
                LinkedGroup:v.LinkedGroup,
                Publisher : v.Publisher,
                Date : v.Date,
                CreatedAt : v.CreatedAt,
                CreatedBy : v.CreatedBy,
                Confirmed1By : v.Confirmed1By,
                Confirmed1At : v.Confirmed1At,
                Confirmed2By : v.Confirmed2By,
				Confirmed2At : v.Confirmed2At,
				Score : v.Score,
            }
        }
    })
    return
}


func (s *sDemoGen16)GetById(ctx context.Context, id uint) (res *model.DemoGen16InfoRes,err error){
    err =g.Try(ctx, func(ctx context.Context){
        err = dao.DemoGen16.Ctx(ctx).WithAll().Where(dao.DemoGen16.Columns().Id,  id).Scan(&res)
        liberr.ErrIsNil(ctx,err,"获取信息失败")
    })
    return
}


func (s *sDemoGen16)Add(ctx context.Context, req *demo16.DemoGen16AddReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {
        //demoCate := ""
        //req.DemoCate.FilterEmpty()
        //if !req.DemoCate.IsEmpty(){
            //demoCate=req.DemoCate.Join(",")
        //}
        for _,obj:=range req.File{
            obj.Url,err = libUtils.GetFilesPath(ctx,obj.Url)
            liberr.ErrIsNil(ctx, err)
        }
        for _,obj:=range req.Image{
            obj.Url,err = libUtils.GetFilesPath(ctx,obj.Url)
            liberr.ErrIsNil(ctx, err)
        }
        _, err = dao.DemoGen16.Ctx(ctx).Insert(do.DemoGen16{
            Type:req.Type,
            Subject:req.Subject,
            Group:req.Group,
            Publisher:req.Publisher,
            Date:req.Date,
			File:req.File,
			Image:req.Image,
			Confirmed1By:req.Confirmed1By,
            Confirmed1At:req.Confirmed1At,
            Confirmed2By:req.Confirmed2By,
			Confirmed2At:req.Confirmed2At,
			Score:req.Score,
            CreatedBy:systemService.Context().GetUserId(ctx),
        })
        liberr.ErrIsNil(ctx, err, "添加失败")
    })
    return
}


func (s *sDemoGen16)Edit(ctx context.Context, req *demo16.DemoGen16EditReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {
        //demoCate := ""
        //req.DemoCate.FilterEmpty()
        //if !req.DemoCate.IsEmpty(){
            //demoCate=req.DemoCate.Join(",")
        //}
        for _,obj:=range req.File{
            obj.Url,err = libUtils.GetFilesPath(ctx,obj.Url)
            liberr.ErrIsNil(ctx, err)
        }
        for _,obj:=range req.Image{
            obj.Url,err = libUtils.GetFilesPath(ctx,obj.Url)
            liberr.ErrIsNil(ctx, err)
        }
        _, err = dao.DemoGen16.Ctx(ctx).WherePri(req.Id).Update(do.DemoGen16{
            Type:req.Type,
            Subject:req.Subject,
            Group:req.Group,
            Publisher:req.Publisher,
            Date:req.Date,
			File:req.File,
			Image:req.Image,
			Confirmed1By:req.Confirmed1By,
            Confirmed1At:req.Confirmed1At,
            Confirmed2By:req.Confirmed2By,
			Confirmed2At:req.Confirmed2At,
			Score:req.Score,
            UpdatedBy:systemService.Context().GetUserId(ctx),
        })
        liberr.ErrIsNil(ctx, err, "修改失败")
    })
    return
}


func (s *sDemoGen16)Delete(ctx context.Context, ids []uint) (err error){
    err = g.Try(ctx,func(ctx context.Context){
        _, err = dao.DemoGen16.Ctx(ctx).Delete(dao.DemoGen16.Columns().Id+" in (?)", ids)
        liberr.ErrIsNil(ctx,err,"删除失败")
    })
    return
}
