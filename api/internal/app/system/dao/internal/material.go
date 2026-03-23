// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MaterialDao is the data access object for the table material.
type MaterialDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  MaterialColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// MaterialColumns defines and stores column names for the table material.
type MaterialColumns struct {
	Id              string //
	Name            string // 物料名称
	Type            string // 物料类型 1农药 2化肥
	Category        string // 分类
	UsageMethod     string // 使用方式
	NutrientContent string // 养分含量
	CreatedAt       string //
}

// materialColumns holds the columns for the table material.
var materialColumns = MaterialColumns{
	Id:              "id",
	Name:            "name",
	Type:            "type",
	Category:        "category",
	UsageMethod:     "usage_method",
	NutrientContent: "nutrient_content",
	CreatedAt:       "created_at",
}

// NewMaterialDao creates and returns a new DAO object for table data access.
func NewMaterialDao(handlers ...gdb.ModelHandler) *MaterialDao {
	return &MaterialDao{
		group:    "default",
		table:    "material",
		columns:  materialColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MaterialDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MaterialDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MaterialDao) Columns() MaterialColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MaterialDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MaterialDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MaterialDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
