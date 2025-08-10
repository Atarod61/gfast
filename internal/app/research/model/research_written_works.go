package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
)

// DemoGenClassInfoRes is the golang structure for table demo_gen_class.
type ResearchWrittenWorksInfoRes struct {
	gmeta.Meta `orm:"table:research_written_works"`
	Id         uint   `orm:"id,primary" json:"id"`        // 分类id
	ClassName  string `orm:"class_name" json:"className"` // 分类名
}

type ResearchWrittenWorksListRes struct {
	Id        uint   `json:"id"`
	ClassName string `json:"className"`
}
