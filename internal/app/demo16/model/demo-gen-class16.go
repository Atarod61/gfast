package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
)

// DemoGenClassInfoRes is the golang structure for table demo_gen_class.
type DemoGenClass16InfoRes struct {
	gmeta.Meta `orm:"table:demo_gen_class16"`
	Id         uint   `orm:"id,primary" json:"id"`        // 分类id
	ClassName  string `orm:"class_name" json:"className"` // 分类名
}

type DemoGenClass16ListRes struct {
	Id        uint   `json:"id"`
	ClassName string `json:"className"`
}
