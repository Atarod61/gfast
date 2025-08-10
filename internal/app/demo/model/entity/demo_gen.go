// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DemoGen is the golang structure for table demo_gen.
type DemoGen struct {
	Id         uint        `json:"id"         orm:"id"          description:""`
	DemoName   string      `json:"demoName"   orm:"demo_name"   description:"姓名"`
	DemoAge    uint        `json:"demoAge"    orm:"demo_age"    description:"年龄"`
	Classes    string      `json:"classes"    orm:"classes"     description:"班级"`
	DemoBorn   *gtime.Time `json:"demoBorn"   orm:"demo_born"   description:"出生年月"`
	DemoGender uint        `json:"demoGender" orm:"demo_gender" description:"性别"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建日期"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"修改日期"`
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"  description:"删除日期"`
	CreatedBy  uint64      `json:"createdBy"  orm:"created_by"  description:"创建人"`
	UpdatedBy  uint64      `json:"updatedBy"  orm:"updated_by"  description:"修改人"`
	DemoStatus int         `json:"demoStatus" orm:"demo_status" description:"状态"`
	DemoCate   string      `json:"demoCate"   orm:"demo_cate"   description:"分类"`
	DemoThumb  string      `json:"demoThumb"  orm:"demo_thumb"  description:"头像"`
	DemoPhoto  string      `json:"demoPhoto"  orm:"demo_photo"  description:"相册"`
	DemoInfo   string      `json:"demoInfo"   orm:"demo_info"   description:"个人描述"`
	DemoFile   string      `json:"demoFile"   orm:"demo_file"   description:"相关附件"`
	ClassesTwo string      `json:"classesTwo" orm:"classes_two" description:"班级二"`
}
