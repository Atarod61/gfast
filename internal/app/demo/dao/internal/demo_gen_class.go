// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DemoGenClassDao is the data access object for the table demo_gen_class.
type DemoGenClassDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  DemoGenClassColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// DemoGenClassColumns defines and stores column names for the table demo_gen_class.
type DemoGenClassColumns struct {
	Id        string // 分类id
	ClassName string // 分类名
}

// demoGenClassColumns holds the columns for the table demo_gen_class.
var demoGenClassColumns = DemoGenClassColumns{
	Id:        "id",
	ClassName: "class_name",
}

// NewDemoGenClassDao creates and returns a new DAO object for table data access.
func NewDemoGenClassDao(handlers ...gdb.ModelHandler) *DemoGenClassDao {
	return &DemoGenClassDao{
		group:    "default",
		table:    "demo_gen_class",
		columns:  demoGenClassColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DemoGenClassDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DemoGenClassDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DemoGenClassDao) Columns() DemoGenClassColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DemoGenClassDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DemoGenClassDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DemoGenClassDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
