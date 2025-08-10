package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// DemoGenInfoRes is the golang structure for table demo_gen.
type ResearchInfoRes struct {
	gmeta.Meta   `orm:"table:research"`
	Id           uint                                `orm:"id,primary" json:"id"` // ID
	Type         string                              `orm:"type" json:"type"`     // 姓名
	LinkedType   *LinkedResearchResearchWrittenWorks `orm:"with:id=type" json:"linkedType"`
	Subject      string                              `orm:"subject" json:"subject"`
	Group        string                              `orm:"group" json:"group"`
	LinkedGroup  *LinkedResearchResearchWrittenWorks `orm:"with:id=group" json:"linkedGroup"`
	Publisher    string                              `orm:"publisher" json:"publisher"`
	Date         *gtime.Time                         `orm:"date" json:"date"`                  // 出生年月
	File         string                              `orm:"file" json:"file"`                  // 相关附件
	Image        string                              `orm:"image" json:"image"`                // 相册
	CreatedAt    *gtime.Time                         `orm:"created_at" json:"createdAt"`       // 创建日期
	UpdatedAt    *gtime.Time                         `orm:"updated_at" json:"updatedAt"`       // 修改日期
	DeletedAt    *gtime.Time                         `orm:"deleted_at" json:"deletedAt"`       // 删除日期
	CreatedBy    uint64                              `orm:"created_by" json:"createdBy"`       // 创建人
	UpdatedBy    uint64                              `orm:"updated_by" json:"updatedBy"`       // 修改人
	Confirmed1By string                              `orm:"confirmed1by" json:"confirmed1By"`  // 分类
	Confirmed1At *gtime.Time                         `orm:"confirmed2_at" json:"confirmed1At"` // 出生年月
	Confirmed2By string                              `orm:"confirmed2by" json:"confirmed2By"`  // 头像
	Confirmed2At *gtime.Time                         `orm:"confirmed1_at" json:"confirmed2At"` // 出生年月
	Score        string                              `orm:"score" json:"score"`
}

type LinkedResearchResearchWrittenWorks struct {
	gmeta.Meta `orm:"table:research_written_works"`
	Id         uint   `orm:"id" json:"id"`                // 分类id
	ClassName  string `orm:"class_name" json:"className"` // 分类名
}

type ResearchListRes struct {
	Id           uint                                `json:"id"`
	Type         string                              `json:"type"`
	LinkedType   *LinkedResearchResearchWrittenWorks `orm:"with:id=type" json:"linkedType"`
	Subject      string                              `json:"subject"`
	Group        string                              `json:"group"`
	LinkedGroup  *LinkedResearchResearchWrittenWorks `orm:"with:id=group" json:"linkedGroup"`
	Publisher    string                              `json:"Publisher"`
	Date         *gtime.Time                         `json:"date"`
	CreatedAt    *gtime.Time                         `json:"createdAt"`
	CreatedBy    uint64                              `json:"createdBy"`
	Confirmed1By string                              `json:"confirmed1By"`
	Confirmed1At *gtime.Time                         `json:"confirmed1At"`
	Confirmed2By string                              `json:"confirmed2By"`
	Confirmed2At *gtime.Time                         `json:"confirmed2At"`
	Score        string                              `json:"Score"`
}
