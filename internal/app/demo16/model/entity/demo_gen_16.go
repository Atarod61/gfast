package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DemoGen16 is the golang structure for table demo_gen16.
type DemoGen16 struct {
	Id           uint        `json:"id"           orm:"id"            description:""`
	Type         string      `json:"type"         orm:"type"          description:"type"`
	Subject      string      `json:"subject"      orm:"subject"       description:"subject"`
	Group        string      `json:"group"        orm:"group"         description:"group"`
	Publisher    string      `json:"publisher"    orm:"publisher"     description:"publisher"`
	Date         *gtime.Time `json:"date"         orm:"date"          description:"date"`
	File         string      `json:"file"         orm:"file"          description:"file"`
	Image        string      `json:"image"        orm:"image"         description:"image"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"creation date"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"modification date"`
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:"deletion date"`
	CreatedBy    uint64      `json:"createdBy"    orm:"created_by"    description:"creator"`
	UpdatedBy    uint64      `json:"updatedBy"    orm:"updated_by"    description:"modifier"`
	Confirmed1By string      `json:"confirmed1By" orm:"confirmed1_by" description:"confirmed1_by"`
	Confirmed1At *gtime.Time `json:"confirmed1At" orm:"confirmed1_at" description:"confirmed1_at"`
	Confirmed2By string      `json:"confirmed2By" orm:"confirmed2_by" description:"confirmed2_by"`
	Confirmed2At *gtime.Time `json:"confirmed2At" orm:"confirmed2_at" description:"confirmed2_at"`
	Score        string      `json:"score"        orm:"score"         description:"score"`
}
