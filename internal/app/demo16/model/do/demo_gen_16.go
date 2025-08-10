// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DemoGen16 is the golang structure of table demo_gen16 for DAO operations like Where/Data.
type DemoGen16 struct {
	g.Meta       `orm:"table:demo_gen16, do:true"`
	Id           interface{} //
	Type         interface{} // type
	Subject      interface{} // subject
	Group        interface{} // group
	Publisher    interface{} // publisher
	Date         *gtime.Time // date
	File         interface{} // file
	Image        interface{} // image
	CreatedAt    *gtime.Time // creation date
	UpdatedAt    *gtime.Time // modification date
	DeletedAt    *gtime.Time // deletion date
	CreatedBy    interface{} // creator
	UpdatedBy    interface{} // modifier
	Confirmed1By interface{} // Confirmed1_by
	Confirmed1At *gtime.Time // Confirmed1_at
	Confirmed2By interface{} // Confirmed2_by
	Confirmed2At *gtime.Time // Confirmed2_at
	Score        interface{} // score
}
