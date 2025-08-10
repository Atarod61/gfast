package demo


import (
"github.com/gogf/gf/v2/frame/g"
commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
"github.com/tiger1103/gfast/v3/internal/app/demo/model"
)
// DemoGenClassSearchReq 分页请求参数
type DemoGenClassSearchReq struct {
    g.Meta   `path:"/list" tags:"分类信息" method:"get" summary:"分类信息列表"`
    ClassName  string `p:"className"` //分类名
    commonApi.PageReq
    commonApi.Author
}
// DemoGenClassSearchRes 列表返回结果
type DemoGenClassSearchRes struct {
    g.Meta `mime:"application/json"`
    commonApi.ListRes
    List []*model.DemoGenClassListRes `json:"list"`
}
// DemoGenClassAddReq 添加操作请求参数
type DemoGenClassAddReq struct {
    g.Meta   `path:"/add" tags:"分类信息" method:"post" summary:"分类信息添加"`
    commonApi.Author
    ClassName  string   `p:"className" v:"required#分类名不能为空"`
}
// DemoGenClassAddRes 添加操作返回结果
type DemoGenClassAddRes  struct {
    commonApi.EmptyRes
}
// DemoGenClassEditReq 修改操作请求参数
type DemoGenClassEditReq struct {
    g.Meta   `path:"/edit" tags:"分类信息" method:"put" summary:"分类信息修改"`
    commonApi.Author
    Id    uint  `p:"id" v:"required#主键ID不能为空"`
    ClassName  string `p:"className" v:"required#分类名不能为空"`
}
// DemoGenClassEditRes 修改操作返回结果
type DemoGenClassEditRes  struct {
    commonApi.EmptyRes
}
// DemoGenClassGetReq 获取一条数据请求
type DemoGenClassGetReq struct {
    g.Meta `path:"/get" tags:"分类信息" method:"get" summary:"获取分类信息信息"`
    commonApi.Author
    Id    uint `p:"id" v:"required#主键必须"` //通过主键获取
}
// DemoGenClassGetRes 获取一条数据结果
type DemoGenClassGetRes struct {
    g.Meta `mime:"application/json"`
    *model.DemoGenClassInfoRes
}
// DemoGenClassDeleteReq 删除数据请求
type DemoGenClassDeleteReq struct {
    g.Meta `path:"/delete" tags:"分类信息" method:"delete" summary:"删除分类信息"`
    commonApi.Author
    Ids    []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}
// DemoGenClassDeleteRes 删除数据返回
type DemoGenClassDeleteRes struct {
    commonApi.EmptyRes
}