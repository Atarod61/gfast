package model


import (
"github.com/gogf/gf/v2/util/gmeta"
)


// DemoGenClassInfoRes is the golang structure for table demo_gen_class.
type DemoGenClassInfoRes struct {
    gmeta.Meta   `orm:"table:demo_gen_class"`
    Id       uint         `orm:"id,primary" json:"id"`    // 分类id
    ClassName    string         `orm:"class_name" json:"className"`    // 分类名
}




type DemoGenClassListRes struct{
    Id  uint   `json:"id"`
    ClassName  string   `json:"className"`
}