// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ResearchWrittenWorksDao is the data access object for the table research_written_works.
type ResearchWrittenWorksDao struct {
	table    string                      // table is the underlying table name of the DAO.
	group    string                      // group is the database configuration group name of the current DAO.
	columns  ResearchWrittenWorksColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler          // handlers for customized model modification.
}

// ResearchWrittenWorksColumns defines and stores column names for the table research_written_works.
type ResearchWrittenWorksColumns struct {
	Id        string // categoryid
	ClassName string // category name
}

// researchWrittenWorksColumns holds the columns for the table research_written_works.
var researchWrittenWorksColumns = ResearchWrittenWorksColumns{
	Id:        "id",
	ClassName: "class_name",
}

// NewResearchWrittenWorksDao creates and returns a new DAO object for table data access.
func NewResearchWrittenWorksDao(handlers ...gdb.ModelHandler) *ResearchWrittenWorksDao {
	return &ResearchWrittenWorksDao{
		group:    "default",
		table:    "research_written_works",
		columns:  researchWrittenWorksColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ResearchWrittenWorksDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ResearchWrittenWorksDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ResearchWrittenWorksDao) Columns() ResearchWrittenWorksColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ResearchWrittenWorksDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ResearchWrittenWorksDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ResearchWrittenWorksDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
