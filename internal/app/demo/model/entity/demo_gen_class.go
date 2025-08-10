// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// DemoGenClass is the golang structure for table demo_gen_class.
type DemoGenClass struct {
	Id        uint   `json:"id"        orm:"id"         description:"分类id"`
	ClassName string `json:"className" orm:"class_name" description:"分类名"`
}
