package model


import (
    "github.com/gogf/gf/v2/os/gtime"
    "github.com/gogf/gf/v2/util/gmeta"
)


// DemoGenInfoRes is the golang structure for table demo_gen.
type DemoGenInfoRes struct {
    gmeta.Meta   `orm:"table:demo_gen"`
    Id       uint         `orm:"id,primary" json:"id"`    // ID
    DemoName    string         `orm:"demo_name" json:"demoName"`    // 姓名
    DemoAge    uint         `orm:"demo_age" json:"demoAge"`    // 年龄
    Classes    string         `orm:"classes" json:"classes"`    // 班级
    LinkedClasses  *LinkedDemoGenDemoGenClass  `orm:"with:id=classes" json:"linkedClasses"`
    ClassesTwo    string         `orm:"classes_two" json:"classesTwo"`    // 班级2
    LinkedClassesTwo  *LinkedDemoGenDemoGenClass  `orm:"with:id=classes_two" json:"linkedClassesTwo"`
    DemoBorn    *gtime.Time         `orm:"demo_born" json:"demoBorn"`    // 出生年月
    DemoGender    uint         `orm:"demo_gender" json:"demoGender"`    // 性别
    CreatedAt    *gtime.Time         `orm:"created_at" json:"createdAt"`    // 创建日期
    UpdatedAt    *gtime.Time         `orm:"updated_at" json:"updatedAt"`    // 修改日期
    DeletedAt    *gtime.Time         `orm:"deleted_at" json:"deletedAt"`    // 删除日期
    CreatedBy    uint64         `orm:"created_by" json:"createdBy"`    // 创建人
    UpdatedBy    uint64         `orm:"updated_by" json:"updatedBy"`    // 修改人
    DemoStatus    int         `orm:"demo_status" json:"demoStatus"`    // 状态
    DemoCate    string         `orm:"demo_cate" json:"demoCate"`    // 分类
    DemoThumb    string         `orm:"demo_thumb" json:"demoThumb"`    // 头像
    DemoPhoto    string         `orm:"demo_photo" json:"demoPhoto"`    // 相册
    DemoInfo    string         `orm:"demo_info" json:"demoInfo"`    // 个人描述
    DemoFile    string         `orm:"demo_file" json:"demoFile"`    // 相关附件
}


type LinkedDemoGenDemoGenClass struct {
    gmeta.Meta   `orm:"table:demo_gen_class"`
    Id    uint         `orm:"id" json:"id"`    // 分类id
    ClassName    string         `orm:"class_name" json:"className"`    // 分类名
}


type DemoGenListRes struct{
    Id  uint   `json:"id"`
    DemoName  string   `json:"demoName"`
    DemoAge  uint   `json:"demoAge"`
    Classes  string   `json:"classes"`
    LinkedClasses  *LinkedDemoGenDemoGenClass  `orm:"with:id=classes" json:"linkedClasses"`
    ClassesTwo  string   `json:"classesTwo"`
    LinkedClassesTwo  *LinkedDemoGenDemoGenClass  `orm:"with:id=classes_two" json:"linkedClassesTwo"`
    DemoBorn  *gtime.Time   `json:"demoBorn"`
    DemoGender  uint   `json:"demoGender"`
    CreatedAt  *gtime.Time   `json:"createdAt"`
    CreatedBy  uint64   `json:"createdBy"`
    DemoStatus  int   `json:"demoStatus"`
    DemoCate  string   `json:"demoCate"`
    DemoThumb  string   `json:"demoThumb"`
}