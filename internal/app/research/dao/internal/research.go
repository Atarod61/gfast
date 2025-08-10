// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ResearchDao is the data access object for the table research.
type ResearchDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ResearchColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ResearchColumns defines and stores column names for the table research.
type ResearchColumns struct {
	Id           string //
	Type         string // type
	Subject      string // subject
	Group        string // group
	Publisher    string // publisher
	Date         string // date
	File         string // file
	Image        string // image
	CreatedAt    string // creation date
	UpdatedAt    string // modification date
	DeletedAt    string // deletion date
	CreatedBy    string // creator
	UpdatedBy    string // modifier
	Confirmed1By string // confirmed1_by
	Confirmed1At string // confirmed1_at
	Confirmed2By string // confirmed2_by
	Confirmed2At string // confirmed2_at
	Score        string // score
}

// researchColumns holds the columns for the table research.
var researchColumns = ResearchColumns{
	Id:           "id",
	Type:         "type",
	Subject:      "subject",
	Group:        "group",
	Publisher:    "publisher",
	Date:         "date",
	File:         "file",
	Image:        "image",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	CreatedBy:    "created_by",
	UpdatedBy:    "updated_by",
	Confirmed1By: "confirmed1_by",
	Confirmed1At: "confirmed1_at",
	Confirmed2By: "confirmed2_by",
	Confirmed2At: "confirmed2_at",
	Score:        "score",
}

// NewResearchDao creates and returns a new DAO object for table data access.
func NewResearchDao(handlers ...gdb.ModelHandler) *ResearchDao {
	return &ResearchDao{
		group:    "default",
		table:    "research",
		columns:  researchColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ResearchDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ResearchDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ResearchDao) Columns() ResearchColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ResearchDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ResearchDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ResearchDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
