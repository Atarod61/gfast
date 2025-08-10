package demoGen


import (
    "context"
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/util/gconv"
    "github.com/tiger1103/gfast/v3/api/v1/demo"
    "github.com/tiger1103/gfast/v3/internal/app/demo/dao"
    "github.com/tiger1103/gfast/v3/internal/app/demo/model"
    "github.com/tiger1103/gfast/v3/internal/app/demo/model/do"
    "github.com/tiger1103/gfast/v3/internal/app/system/consts"
    systemService "github.com/tiger1103/gfast/v3/internal/app/system/service"
    "github.com/tiger1103/gfast/v3/library/libUtils"
    "github.com/tiger1103/gfast/v3/library/liberr"
    "github.com/tiger1103/gfast/v3/internal/app/demo/service"
)

func init() {
    service.RegisterDemoGen(New())
}


func New() *sDemoGen {
    return &sDemoGen{}
}


type sDemoGen struct{}


func (s *sDemoGen)List(ctx context.Context, req *demo.DemoGenSearchReq) (listRes *demo.DemoGenSearchRes, err error){
    listRes = new(demo.DemoGenSearchRes)
    err = g.Try(ctx, func(ctx context.Context) {
        m := dao.DemoGen.Ctx(ctx).WithAll()
        if req.DemoName != "" {
            m = m.Where(dao.DemoGen.Columns().DemoName+" like ?", "%"+req.DemoName+"%")
        }
        if req.DemoAge != "" {
            m = m.Where(dao.DemoGen.Columns().DemoAge+" = ?", gconv.Uint(req.DemoAge))
        }
        if req.Classes != "" {
            m = m.Where(dao.DemoGen.Columns().Classes+" = ?", req.Classes)
        }
        if req.ClassesTwo != "" {
            m = m.Where(dao.DemoGen.Columns().ClassesTwo+" = ?", req.ClassesTwo)
        }
        if req.DemoBorn != "" {
            m = m.Where(dao.DemoGen.Columns().DemoBorn+" = ?", gconv.Time(req.DemoBorn))
        }
        if req.DemoGender != "" {
            m = m.Where(dao.DemoGen.Columns().DemoGender+" = ?", gconv.Uint(req.DemoGender))
        }
        if req.DemoStatus != "" {
            m = m.Where(dao.DemoGen.Columns().DemoStatus+" = ?", gconv.Int(req.DemoStatus))
        }
        if req.DemoCate != "" {
            m = m.Where(dao.DemoGen.Columns().DemoCate+" = ?", req.DemoCate)
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
        var res []*model.DemoGenInfoRes
        err = m.Fields(demo.DemoGenSearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
        liberr.ErrIsNil(ctx, err, "获取数据失败")
        listRes.List = make([]*model.DemoGenListRes,len(res))
        for k,v:=range res{
            listRes.List[k] = &model.DemoGenListRes{
                Id : v.Id,
                DemoName : v.DemoName,
                DemoAge : v.DemoAge,
                Classes : v.Classes,
                LinkedClasses:v.LinkedClasses,
                ClassesTwo : v.ClassesTwo,
                LinkedClassesTwo:v.LinkedClassesTwo,
                DemoBorn : v.DemoBorn,
                DemoGender : v.DemoGender,
                CreatedAt : v.CreatedAt,
                CreatedBy : v.CreatedBy,
                DemoStatus : v.DemoStatus,
                DemoCate : v.DemoCate,
                DemoThumb : v.DemoThumb,
            }
        }
    })
    return
}


func (s *sDemoGen)GetById(ctx context.Context, id uint) (res *model.DemoGenInfoRes,err error){
    err =g.Try(ctx, func(ctx context.Context){
        err = dao.DemoGen.Ctx(ctx).WithAll().Where(dao.DemoGen.Columns().Id,  id).Scan(&res)
        liberr.ErrIsNil(ctx,err,"获取信息失败")
    })
    return
}


func (s *sDemoGen)Add(ctx context.Context, req *demo.DemoGenAddReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {
        demoCate := ""
        req.DemoCate.FilterEmpty()
        if !req.DemoCate.IsEmpty(){
            demoCate=req.DemoCate.Join(",")
        }
        for _,obj:=range req.DemoPhoto{
            obj.Url,err = libUtils.GetFilesPath(ctx,obj.Url)
            liberr.ErrIsNil(ctx, err)
        }
        for _,obj:=range req.DemoFile{
            obj.Url,err = libUtils.GetFilesPath(ctx,obj.Url)
            liberr.ErrIsNil(ctx, err)
        }
        _, err = dao.DemoGen.Ctx(ctx).Insert(do.DemoGen{
            DemoName:req.DemoName,
            DemoAge:req.DemoAge,
            Classes:req.Classes,
            ClassesTwo:req.ClassesTwo,
            DemoBorn:req.DemoBorn,
            DemoGender:req.DemoGender,
            DemoStatus:req.DemoStatus,
            DemoCate:demoCate,
            DemoThumb:req.DemoThumb,
            DemoPhoto:req.DemoPhoto,
            DemoInfo:req.DemoInfo,
            DemoFile:req.DemoFile,
            CreatedBy:systemService.Context().GetUserId(ctx),
        })
        liberr.ErrIsNil(ctx, err, "添加失败")
    })
    return
}


func (s *sDemoGen)Edit(ctx context.Context, req *demo.DemoGenEditReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {
        demoCate := ""
        req.DemoCate.FilterEmpty()
        if !req.DemoCate.IsEmpty(){
            demoCate=req.DemoCate.Join(",")
        }
        for _,obj:=range req.DemoPhoto{
            obj.Url,err = libUtils.GetFilesPath(ctx,obj.Url)
            liberr.ErrIsNil(ctx, err)
        }
        for _,obj:=range req.DemoFile{
            obj.Url,err = libUtils.GetFilesPath(ctx,obj.Url)
            liberr.ErrIsNil(ctx, err)
        }
        _, err = dao.DemoGen.Ctx(ctx).WherePri(req.Id).Update(do.DemoGen{
            DemoName:req.DemoName,
            DemoAge:req.DemoAge,
            Classes:req.Classes,
            ClassesTwo:req.ClassesTwo,
            DemoBorn:req.DemoBorn,
            DemoGender:req.DemoGender,
            DemoStatus:req.DemoStatus,
            DemoCate:demoCate,
            DemoThumb:req.DemoThumb,
            DemoPhoto:req.DemoPhoto,
            DemoInfo:req.DemoInfo,
            DemoFile:req.DemoFile,
            UpdatedBy:systemService.Context().GetUserId(ctx),
        })
        liberr.ErrIsNil(ctx, err, "修改失败")
    })
    return
}


func (s *sDemoGen)Delete(ctx context.Context, ids []uint) (err error){
    err = g.Try(ctx,func(ctx context.Context){
        _, err = dao.DemoGen.Ctx(ctx).Delete(dao.DemoGen.Columns().Id+" in (?)", ids)
        liberr.ErrIsNil(ctx,err,"删除失败")
    })
    return
}