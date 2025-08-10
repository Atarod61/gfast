package demoGenClass16

import (
    "context"
    "github.com/gogf/gf/v2/frame/g"
    "github.com/tiger1103/gfast/v3/api/v1/demo16"
    "github.com/tiger1103/gfast/v3/internal/app/demo16/dao"
    "github.com/tiger1103/gfast/v3/internal/app/demo16/model"
    "github.com/tiger1103/gfast/v3/internal/app/demo16/model/do"
    "github.com/tiger1103/gfast/v3/internal/app/system/consts"
    "github.com/tiger1103/gfast/v3/library/liberr"
	"github.com/tiger1103/gfast/v3/internal/app/demo16/service"
)

func init() {
    service.RegisterDemoGenClass16(New())
}

func New() *sDemoGenClass16 {
    return &sDemoGenClass16{}
}


type sDemoGenClass16 struct{}


func (s *sDemoGenClass16)List(ctx context.Context, req *demo16.DemoGenClass16SearchReq) (listRes *demo16.DemoGenClass16SearchRes, err error){
    listRes = new(demo16.DemoGenClass16SearchRes)
    err = g.Try(ctx, func(ctx context.Context) {
        m := dao.DemoGenClass16.Ctx(ctx).WithAll()
        if req.ClassName != "" {
            m = m.Where(dao.DemoGenClass16.Columns().ClassName+" like ?", "%"+req.ClassName+"%")
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
        var res []*model.DemoGenClass16InfoRes
        err = m.Fields(demo16.DemoGenClass16SearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
        liberr.ErrIsNil(ctx, err, "获取数据失败")
        listRes.List = make([]*model.DemoGenClass16ListRes,len(res))
        for k,v:=range res{
            listRes.List[k] = &model.DemoGenClass16ListRes{
                Id : v.Id,
                ClassName : v.ClassName,
            }
        }
    })
    return
}


func (s *sDemoGenClass16)GetById(ctx context.Context, id uint) (res *model.DemoGenClass16InfoRes,err error){
    err =g.Try(ctx, func(ctx context.Context){
        err = dao.DemoGenClass16.Ctx(ctx).WithAll().Where(dao.DemoGenClass16.Columns().Id,  id).Scan(&res)
        liberr.ErrIsNil(ctx,err,"获取信息失败")
    })
    return
}


func (s *sDemoGenClass16)Add(ctx context.Context, req *demo16.DemoGenClass16AddReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {
        _, err = dao.DemoGenClass16.Ctx(ctx).Insert(do.DemoGenClass16{
            ClassName:req.ClassName,
        })
        liberr.ErrIsNil(ctx, err, "添加失败")
    })
    return
}


func (s *sDemoGenClass16)Edit(ctx context.Context, req *demo16.DemoGenClass16EditReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {
        _, err = dao.DemoGenClass16.Ctx(ctx).WherePri(req.Id).Update(do.DemoGenClass16{
            ClassName:req.ClassName,
        })
        liberr.ErrIsNil(ctx, err, "修改失败")
    })
    return
}


func (s *sDemoGenClass16)Delete(ctx context.Context, ids []uint) (err error){
    err = g.Try(ctx,func(ctx context.Context){
        _, err = dao.DemoGenClass16.Ctx(ctx).Delete(dao.DemoGenClass16.Columns().Id+" in (?)", ids)
        liberr.ErrIsNil(ctx,err,"删除失败")
    })
    return
}
