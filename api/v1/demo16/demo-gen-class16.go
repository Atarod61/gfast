package demo16

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/demo16/model"
)

// DemoGenClassSearchReq 分页请求参数
type DemoGenClass16SearchReq struct {
	g.Meta    `path:"/list" tags:"分类信息" method:"get" summary:"分类信息列表"`
	ClassName string `p:"className"` //分类名
	commonApi.PageReq
	commonApi.Author
}

// DemoGenClassSearchRes 列表返回结果
type DemoGenClass16SearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.DemoGenClass16ListRes `json:"list"`
}

// DemoGenClassAddReq 添加操作请求参数
type DemoGenClass16AddReq struct {
	g.Meta `path:"/add" tags:"分类信息" method:"post" summary:"分类信息添加"`
	commonApi.Author
	ClassName string `p:"className" v:"required#分类名不能为空"`
}

// DemoGenClassAddRes 添加操作返回结果
type DemoGenClass16AddRes struct {
	commonApi.EmptyRes
}

// DemoGenClassEditReq 修改操作请求参数
type DemoGenClass16EditReq struct {
	g.Meta `path:"/edit" tags:"分类信息" method:"put" summary:"分类信息修改"`
	commonApi.Author
	Id        uint   `p:"id" v:"required#主键ID不能为空"`
	ClassName string `p:"className" v:"required#分类名不能为空"`
}

// DemoGenClassEditRes 修改操作返回结果
type DemoGenClass16EditRes struct {
	commonApi.EmptyRes
}

// DemoGenClassGetReq 获取一条数据请求
type DemoGenClass16GetReq struct {
	g.Meta `path:"/get" tags:"分类信息" method:"get" summary:"获取分类信息信息"`
	commonApi.Author
	Id uint `p:"id" v:"required#主键必须"` //通过主键获取
}

// DemoGenClassGetRes 获取一条数据结果
type DemoGenClass16GetRes struct {
	g.Meta `mime:"application/json"`
	*model.DemoGenClass16InfoRes
}

// DemoGenClassDeleteReq 删除数据请求
type DemoGenClass16DeleteReq struct {
	g.Meta `path:"/delete" tags:"分类信息" method:"delete" summary:"删除分类信息"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}

// DemoGenClassDeleteRes 删除数据返回
type DemoGenClass16DeleteRes struct {
	commonApi.EmptyRes
}
