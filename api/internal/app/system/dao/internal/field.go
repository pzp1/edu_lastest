// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FieldDao is the data access object for the table field.
type FieldDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  FieldColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// FieldColumns defines and stores column names for the table field.
type FieldColumns struct {
	Id        string //
	UserId    string // 所属种植户
	Name      string // 地块名称
	Area      string // 面积（亩）
	Location  string // 位置描述
	CreatedAt string //
	Status    string // 状态 1使用中 0停用
	SortOrder string // 排序
	UpdatedAt string // 更新时间
	Deleted   string // 是否删除 0否 1是
}

// fieldColumns holds the columns for the table field.
var fieldColumns = FieldColumns{
	Id:        "id",
	UserId:    "user_id",
	Name:      "name",
	Area:      "area",
	Location:  "location",
	CreatedAt: "created_at",
	Status:    "status",
	SortOrder: "sort_order",
	UpdatedAt: "updated_at",
	Deleted:   "deleted",
}

// NewFieldDao creates and returns a new DAO object for table data access.
func NewFieldDao(handlers ...gdb.ModelHandler) *FieldDao {
	return &FieldDao{
		group:    "default",
		table:    "field",
		columns:  fieldColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FieldDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FieldDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FieldDao) Columns() FieldColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FieldDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FieldDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FieldDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
