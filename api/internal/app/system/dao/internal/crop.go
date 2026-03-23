// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CropDao is the data access object for the table crop.
type CropDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CropColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CropColumns defines and stores column names for the table crop.
type CropColumns struct {
	Id              string //
	Name            string // 作物名称
	GrowthCycleDays string // 生长周期（天）
	CreatedAt       string //
	Category        string // 种类 1粮食 2水果 3 蔬菜
}

// cropColumns holds the columns for the table crop.
var cropColumns = CropColumns{
	Id:              "id",
	Name:            "name",
	GrowthCycleDays: "growth_cycle_days",
	CreatedAt:       "created_at",
	Category:        "category",
}

// NewCropDao creates and returns a new DAO object for table data access.
func NewCropDao(handlers ...gdb.ModelHandler) *CropDao {
	return &CropDao{
		group:    "default",
		table:    "crop",
		columns:  cropColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CropDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CropDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CropDao) Columns() CropColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CropDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CropDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CropDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
