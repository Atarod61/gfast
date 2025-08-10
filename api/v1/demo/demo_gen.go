package demo


import (
    "github.com/gogf/gf/v2/container/garray"
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/os/gtime"
    commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
    comModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
    "github.com/tiger1103/gfast/v3/internal/app/demo/model"
)
// DemoGenSearchReq 分页请求参数
type DemoGenSearchReq struct {
    g.Meta   `path:"/list" tags:"代码生成测试" method:"get" summary:"代码生成测试列表"`
    DemoName  string `p:"demoName"` //姓名
    DemoAge  string `p:"demoAge" v:"demoAge@integer#年龄需为整数"` //年龄
    Classes  string `p:"classes"` //班级
    ClassesTwo  string `p:"classesTwo"` //班级2
    DemoBorn  string `p:"demoBorn" v:"demoBorn@datetime#出生年月需为YYYY-MM-DD hh:mm:ss格式"` //出生年月
    DemoGender  string `p:"demoGender" v:"demoGender@integer#性别需为整数"` //性别
    DemoStatus  string `p:"demoStatus" v:"demoStatus@integer#状态需为整数"` //状态
    DemoCate  string `p:"demoCate"` //分类
    commonApi.PageReq
    commonApi.Author
}
// DemoGenSearchRes 列表返回结果
type DemoGenSearchRes struct {
    g.Meta `mime:"application/json"`
    commonApi.ListRes
    List []*model.DemoGenListRes `json:"list"`
}
// DemoGenAddReq 添加操作请求参数
type DemoGenAddReq struct {
    g.Meta   `path:"/add" tags:"代码生成测试" method:"post" summary:"代码生成测试添加"`
    commonApi.Author
    DemoName  string   `p:"demoName" v:"required#姓名不能为空"`
    DemoAge  uint   `p:"demoAge" `
    Classes  string   `p:"classes" `
    ClassesTwo  string   `p:"classesTwo" `
    DemoBorn  *gtime.Time   `p:"demoBorn" `
    DemoGender  uint   `p:"demoGender" `
    DemoStatus  int   `p:"demoStatus" v:"required#状态不能为空"`
    DemoCate  garray.StrArray   `p:"demoCate" `
    DemoThumb  string   `p:"demoThumb" `
    DemoPhoto  []*comModel.UpFile   `p:"demoPhoto" `
    DemoInfo  string   `p:"demoInfo" `
    DemoFile  []*comModel.UpFile   `p:"demoFile" `
    CreatedBy       uint64
}
// DemoGenAddRes 添加操作返回结果
type DemoGenAddRes  struct {
    commonApi.EmptyRes
}
// DemoGenEditReq 修改操作请求参数
type DemoGenEditReq struct {
    g.Meta   `path:"/edit" tags:"代码生成测试" method:"put" summary:"代码生成测试修改"`
    commonApi.Author
    Id    uint  `p:"id" v:"required#主键ID不能为空"`
    DemoName  string `p:"demoName" v:"required#姓名不能为空"`
    DemoAge  uint `p:"demoAge" `
    Classes  string `p:"classes" `
    ClassesTwo  string `p:"classesTwo" `
    DemoBorn  *gtime.Time `p:"demoBorn" `
    DemoGender  uint `p:"demoGender" `
    DemoStatus  int `p:"demoStatus" v:"required#状态不能为空"`
    DemoCate  garray.StrArray `p:"demoCate" `
    DemoThumb  string `p:"demoThumb" `
    DemoPhoto  []*comModel.UpFile `p:"demoPhoto" `
    DemoInfo  string `p:"demoInfo" `
    DemoFile  []*comModel.UpFile `p:"demoFile" `
    UpdatedBy       uint64
}
// DemoGenEditRes 修改操作返回结果
type DemoGenEditRes  struct {
    commonApi.EmptyRes
}
// DemoGenGetReq 获取一条数据请求
type DemoGenGetReq struct {
    g.Meta `path:"/get" tags:"代码生成测试" method:"get" summary:"获取代码生成测试信息"`
    commonApi.Author
    Id    uint `p:"id" v:"required#主键必须"` //通过主键获取
}
// DemoGenGetRes 获取一条数据结果
type DemoGenGetRes struct {
    g.Meta `mime:"application/json"`
    *model.DemoGenInfoRes
}
// DemoGenDeleteReq 删除数据请求
type DemoGenDeleteReq struct {
    g.Meta `path:"/delete" tags:"代码生成测试" method:"delete" summary:"删除代码生成测试"`
    commonApi.Author
    Ids    []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}
// DemoGenDeleteRes 删除数据返回
type DemoGenDeleteRes struct {
    commonApi.EmptyRes
}