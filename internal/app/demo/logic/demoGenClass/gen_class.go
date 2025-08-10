package demoGenClass

import (
    "context"
    "github.com/gogf/gf/v2/frame/g"
    "github.com/tiger1103/gfast/v3/api/v1/demo"
    "github.com/tiger1103/gfast/v3/internal/app/demo/dao"
    "github.com/tiger1103/gfast/v3/internal/app/demo/model"
    "github.com/tiger1103/gfast/v3/internal/app/demo/model/do"
    "github.com/tiger1103/gfast/v3/internal/app/system/consts"
    "github.com/tiger1103/gfast/v3/library/liberr"
    "github.com/tiger1103/gfast/v3/internal/app/demo/service"
)


func init() {
    service.RegisterDemoGenClass(New())
}

func New() *sDemoGenClass {
    return &sDemoGenClass{}
}


type sDemoGenClass struct{}


func (s *sDemoGenClass)List(ctx context.Context, req *demo.DemoGenClassSearchReq) (listRes *demo.DemoGenClassSearchRes, err error){
    listRes = new(demo.DemoGenClassSearchRes)
    err = g.Try(ctx, func(ctx context.Context) {
        m := dao.DemoGenClass.Ctx(ctx).WithAll()
        if req.ClassName != "" {
            m = m.Where(dao.DemoGenClass.Columns().ClassName+" like ?", "%"+req.ClassName+"%")
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
        var res []*model.DemoGenClassInfoRes
        err = m.Fields(demo.DemoGenClassSearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
        liberr.ErrIsNil(ctx, err, "获取数据失败")
        listRes.List = make([]*model.DemoGenClassListRes,len(res))
        for k,v:=range res{
            listRes.List[k] = &model.DemoGenClassListRes{
                Id : v.Id,
                ClassName : v.ClassName,
            }
        }
    })
    return
}


func (s *sDemoGenClass)GetById(ctx context.Context, id uint) (res *model.DemoGenClassInfoRes,err error){
    err =g.Try(ctx, func(ctx context.Context){
        err = dao.DemoGenClass.Ctx(ctx).WithAll().Where(dao.DemoGenClass.Columns().Id,  id).Scan(&res)
        liberr.ErrIsNil(ctx,err,"获取信息失败")
    })
    return
}


func (s *sDemoGenClass)Add(ctx context.Context, req *demo.DemoGenClassAddReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {
        _, err = dao.DemoGenClass.Ctx(ctx).Insert(do.DemoGenClass{
            ClassName:req.ClassName,
        })
        liberr.ErrIsNil(ctx, err, "添加失败")
    })
    return
}


func (s *sDemoGenClass)Edit(ctx context.Context, req *demo.DemoGenClassEditReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {
        _, err = dao.DemoGenClass.Ctx(ctx).WherePri(req.Id).Update(do.DemoGenClass{
            ClassName:req.ClassName,
        })
        liberr.ErrIsNil(ctx, err, "修改失败")
    })
    return
}


func (s *sDemoGenClass)Delete(ctx context.Context, ids []uint) (err error){
    err = g.Try(ctx,func(ctx context.Context){
        _, err = dao.DemoGenClass.Ctx(ctx).Delete(dao.DemoGenClass.Columns().Id+" in (?)", ids)
        liberr.ErrIsNil(ctx,err,"删除失败")
    })
    return
}