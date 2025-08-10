// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DemoGenClass is the golang structure of table demo_gen_class for DAO operations like Where/Data.
type DemoGenClass struct {
	g.Meta    `orm:"table:demo_gen_class, do:true"`
	Id        interface{} // 分类id
	ClassName interface{} // 分类名
}
