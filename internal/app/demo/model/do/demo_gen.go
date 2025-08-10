// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DemoGen is the golang structure of table demo_gen for DAO operations like Where/Data.
type DemoGen struct {
	g.Meta     `orm:"table:demo_gen, do:true"`
	Id         interface{} //
	DemoName   interface{} // 姓名
	DemoAge    interface{} // 年龄
	Classes    interface{} // 班级
	DemoBorn   *gtime.Time // 出生年月
	DemoGender interface{} // 性别
	CreatedAt  *gtime.Time // 创建日期
	UpdatedAt  *gtime.Time // 修改日期
	DeletedAt  *gtime.Time // 删除日期
	CreatedBy  interface{} // 创建人
	UpdatedBy  interface{} // 修改人
	DemoStatus interface{} // 状态
	DemoCate   interface{} // 分类
	DemoThumb  interface{} // 头像
	DemoPhoto  interface{} // 相册
	DemoInfo   interface{} // 个人描述
	DemoFile   interface{} // 相关附件
	ClassesTwo interface{} // 班级二
}
