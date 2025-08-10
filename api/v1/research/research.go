package research

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	comModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
	"github.com/tiger1103/gfast/v3/internal/app/research/model"
)

// DemoGenSearchReq 分页请求参数
type ResearchSearchReq struct {
	g.Meta       `path:"/list" tags:"Code Generation Test" method:"get" summary:"Code Generation Test List"`
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
type ResearchSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.ResearchListRes `json:"list"`
}

// DemoGenAddReq 添加操作请求参数
type ResearchAddReq struct {
	g.Meta `path:"/add" tags:"Code Generation Test" method:"post" summary:"Code Generation Test Add"`
	commonApi.Author
	Type         garray.StrArray    `p:"type" v:"required#Type cannot be empty"`
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
type ResearchAddRes struct {
	commonApi.EmptyRes
}

// DemoGenEditReq 修改操作请求参数
type ResearchEditReq struct {
	g.Meta `path:"/edit" tags:"Code Generation Test" method:"put" summary:"Code Generation Test Modification"`
	commonApi.Author
	Id           uint               `p:"id" v:"required#Primary keyIDcannot be empty"`
	Type         garray.StrArray    `p:"type" v:"required#Type cannot be empty"`
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
type ResearchEditRes struct {
	commonApi.EmptyRes
}

// DemoGenGetReq 获取一条数据请求
type ResearchGetReq struct {
	g.Meta `path:"/get" tags:"Code Generation Test" method:"get" summary:"Get Code Generation Test Information"`
	commonApi.Author
	Id uint `p:"id" v:"required#Primary Key Required"` //通过主键获取
}

// DemoGenGetRes 获取一条数据结果
type ResearchGetRes struct {
	g.Meta `mime:"application/json"`
	*model.ResearchInfoRes
}

// DemoGenDeleteReq 删除数据请求
type ResearchDeleteReq struct {
	g.Meta `path:"/delete" tags:"Code Generation Test" method:"delete" summary:"Delete Code Generation Test"`
	commonApi.Author
	Ids []uint `p:"ids" v:"required#Primary key required"` //通过主键删除
}

// DemoGenDeleteRes 删除数据返回
type ResearchDeleteRes struct {
	commonApi.EmptyRes
}
