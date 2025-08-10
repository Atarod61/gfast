// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Research is the golang structure of table research for DAO operations like Where/Data.
type Research struct {
	g.Meta       `orm:"table:research, do:true"`
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
	Confirmed1By interface{} // confirmed1_by
	Confirmed1At *gtime.Time // confirmed1_at
	Confirmed2By interface{} // confirmed2_by
	Confirmed2At *gtime.Time // confirmed2_at
	Score        interface{} // score
}
