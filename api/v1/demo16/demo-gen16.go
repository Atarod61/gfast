package demo16

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	comModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
	"github.com/tiger1103/gfast/v3/internal/app/demo16/model"
)

// DemoGenSearchReq 分页请求参数
type DemoGen16SearchReq struct {
	g.Meta       `path:"/list" tags:"代码生成测试" method:"get" summary:"代码生成测试列表"`
	Type         string `p:"type"`    //姓名
	Subject      string `p:"subject"` //班级
	Group        string `p:"group"`
	Publisher    string `p:"publisher"`
	Date         string `p:"date" v:"Date@datetime#The publication date should be inYYYY-MM-DD hh:mm:ssformat"` //出生年月
	Confirmed1By string `p:"confirmed1By"`                                                                      //分类
	Confirmed1At string `p:"confirmed1At" `
	Confirmed2By string `p:"confirmed2By"`
	Confirmed2At string `p:"confirmed2At" `
	Score        string `p:"score"`
	commonApi.PageReq
	commonApi.Author
}

// DemoGenSearchRes 列表返回结果
type DemoGen16SearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.DemoGen16ListRes `json:"list"`
}

// DemoGenAddReq 添加操作请求参数
type DemoGen16AddReq struct {
	g.Meta `path:"/add" tags:"代码生成测试" method:"post" summary:"代码生成测试添加"`
	commonApi.Author
	Type         garray.StrArray    `p:"type" v:"required#姓名不能为空"`
	Subject      string             `p:"subject" `
	Group        string             `p:"group" `
	Publisher    string             `p:"publisher" `
	Date         *gtime.Time        `p:"date" `
	File         []*comModel.UpFile `p:"File" `
	Image        []*comModel.UpFile `p:"image" `
	Confirmed1By string             `p:"confirmed1By" `
	Confirmed1At *gtime.Time        `p:"confirmed1At" `
	Confirmed2By string             `p:"confirmed2By" `
	Confirmed2At *gtime.Time        `p:"confirmed2At" `
	Score        string             `p:"score" `
	CreatedBy    uint64
}

// DemoGenAddRes 添加操作返回结果
type DemoGen16AddRes struct {
	commonApi.EmptyRes
}

// DemoGenEditReq 修改操作请求参数
type DemoGen16EditReq struct {
	g.Meta `path:"/edit" tags:"代码生成测试" method:"put" summary:"代码生成测试修改"`
	commonApi.Author
	Id           uint               `p:"id" v:"required#主键ID不能为空"`
	Type         garray.StrArray    `p:"type" v:"required#姓名不能为空"`
	Subject      string             `p:"subject" `
	Group        string             `p:"group" `
	Publisher    string             `p:"publisher" `
	Date         *gtime.Time        `p:"date" `
	File         []*comModel.UpFile `p:"file" `
	Image        []*comModel.UpFile `p:"image" `
	Confirmed1By string             `p:"confirmed1By" `
	Confirmed1At *gtime.Time        `p:"confirmed1At" `
	Confirmed2By string             `p:"confirmed2By" `
	Confirmed2At *gtime.Time        `p:"confirmed2At" `
	Score        string             `p:"Score" `
	UpdatedBy    uint64
}

// DemoGenEditRes 修改操作返回结果
type DemoGen16EditRes struct {
	commonApi.EmptyRes
}

// DemoGenGetReq 获取一条数据请求
type DemoGen16GetReq struct {
	g.Meta `path:"/get" tags:"代码生成测试" method:"get" summary:"获取代码生成测试信息"`
	commonApi.Author
	Id uint `p:"id" v:"required#主键必须"` //通过主键获取
}

// DemoGenGetRes 获取一条数据结果
type DemoGen16GetRes struct {
	g.Meta `mime:"application/json"`
	*model.DemoGen16InfoRes
}

// DemoGenDeleteReq 删除数据请求
type DemoGen16DeleteReq struct {
	g.Meta `path:"/delete" tags:"代码生成测试" method:"delete" summary:"删除代码生成测试"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#主键必须"` //通过主键删除
}

// DemoGenDeleteRes 删除数据返回
type DemoGen16DeleteRes struct {
	commonApi.EmptyRes
}
