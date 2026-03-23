// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PestRecordDao is the data access object for the table pest_record.
type PestRecordDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PestRecordColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PestRecordColumns defines and stores column names for the table pest_record.
type PestRecordColumns struct {
	Id          string // 主键ID
	FieldId     string // 地块ID
	FieldName   string // 地块名称
	CropId      string // 作物ID
	CropName    string // 作物名称
	Description string // 虫灾描述
	Status      string // 状态 0未解决 1已解决
	CreatedBy   string // 创建人
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	SloveAnswer string // 解决办法
	SolveBy     string // 解决人
}

// pestRecordColumns holds the columns for the table pest_record.
var pestRecordColumns = PestRecordColumns{
	Id:          "id",
	FieldId:     "field_id",
	FieldName:   "field_name",
	CropId:      "crop_id",
	CropName:    "crop_name",
	Description: "description",
	Status:      "status",
	CreatedBy:   "created_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	SloveAnswer: "slove_answer",
	SolveBy:     "solve_by",
}

// NewPestRecordDao creates and returns a new DAO object for table data access.
func NewPestRecordDao(handlers ...gdb.ModelHandler) *PestRecordDao {
	return &PestRecordDao{
		group:    "default",
		table:    "pest_record",
		columns:  pestRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PestRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PestRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PestRecordDao) Columns() PestRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PestRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PestRecordDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PestRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
