// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DemoGenDao is the data access object for the table demo_gen.
type DemoGenDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DemoGenColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DemoGenColumns defines and stores column names for the table demo_gen.
type DemoGenColumns struct {
	Id         string //
	DemoName   string // 姓名
	DemoAge    string // 年龄
	Classes    string // 班级
	DemoBorn   string // 出生年月
	DemoGender string // 性别
	CreatedAt  string // 创建日期
	UpdatedAt  string // 修改日期
	DeletedAt  string // 删除日期
	CreatedBy  string // 创建人
	UpdatedBy  string // 修改人
	DemoStatus string // 状态
	DemoCate   string // 分类
	DemoThumb  string // 头像
	DemoPhoto  string // 相册
	DemoInfo   string // 个人描述
	DemoFile   string // 相关附件
	ClassesTwo string // 班级二
}

// demoGenColumns holds the columns for the table demo_gen.
var demoGenColumns = DemoGenColumns{
	Id:         "id",
	DemoName:   "demo_name",
	DemoAge:    "demo_age",
	Classes:    "classes",
	DemoBorn:   "demo_born",
	DemoGender: "demo_gender",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	CreatedBy:  "created_by",
	UpdatedBy:  "updated_by",
	DemoStatus: "demo_status",
	DemoCate:   "demo_cate",
	DemoThumb:  "demo_thumb",
	DemoPhoto:  "demo_photo",
	DemoInfo:   "demo_info",
	DemoFile:   "demo_file",
	ClassesTwo: "classes_two",
}

// NewDemoGenDao creates and returns a new DAO object for table data access.
func NewDemoGenDao(handlers ...gdb.ModelHandler) *DemoGenDao {
	return &DemoGenDao{
		group:    "default",
		table:    "demo_gen",
		columns:  demoGenColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DemoGenDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DemoGenDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DemoGenDao) Columns() DemoGenColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DemoGenDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DemoGenDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *DemoGenDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
