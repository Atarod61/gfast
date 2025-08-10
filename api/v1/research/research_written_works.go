package research

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/research/model"
)

// DemoGenClassSearchReq 分页请求参数
type ResearchWrittenWorksSearchReq struct {
	g.Meta    `path:"/list" tags:"Classification Information" method:"get" summary:"Classification Information List"`
	ClassName string `p:"className"` //分类名
	commonApi.PageReq
	commonApi.Author
}

// DemoGenClassSearchRes 列表返回结果
type ResearchWrittenWorksSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.ResearchWrittenWorksListRes `json:"list"`
}

// DemoGenClassAddReq 添加操作请求参数
type ResearchWrittenWorksAddReq struct {
	g.Meta `path:"/add" tags:"Classification Information" method:"post" summary:"Add classification information"`
	commonApi.Author
	ClassName string `p:"className" v:"required#Classification name cannot be empty"`
}

// DemoGenClassAddRes 添加操作返回结果
type ResearchWrittenWorksAddRes struct {
	commonApi.EmptyRes
}

// DemoGenClassEditReq 修改操作请求参数
type ResearchWrittenWorksEditReq struct {
	g.Meta `path:"/edit" tags:"Classification Information" method:"put" summary:"Edit classification information"`
	commonApi.Author
	Id        uint   `p:"id" v:"required#Primary keyIDcannot be empty"`
	ClassName string `p:"className" v:"required#The category name cannot be empty"`
}

// DemoGenClassEditRes 修改操作返回结果
type ResearchWrittenWorksEditRes struct {
	commonApi.EmptyRes
}

// DemoGenClassGetReq 获取一条数据请求
type ResearchWrittenWorksGetReq struct {
	g.Meta `path:"/get" tags:"category information" method:"get" summary:"Get category information"`
	commonApi.Author
	Id uint `p:"id" v:"required#Primary key required"` //通过主键获取
}

// DemoGenClassGetRes 获取一条数据结果
type ResearchWrittenWorksGetRes struct {
	g.Meta `mime:"application/json"`
	*model.ResearchWrittenWorksInfoRes
}

// DemoGenClassDeleteReq 删除数据请求
type ResearchWrittenWorksDeleteReq struct {
	g.Meta `path:"/delete" tags:"Classification Information" method:"delete" summary:"Delete Classification Information"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#Primary key required"` //通过主键删除
}

// DemoGenClassDeleteRes 删除数据返回
type ResearchWrittenWorksDeleteRes struct {
	commonApi.EmptyRes
}
