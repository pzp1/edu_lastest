// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FieldCropMaterialDao is the data access object for the table field_crop_material.
type FieldCropMaterialDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  FieldCropMaterialColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// FieldCropMaterialColumns defines and stores column names for the table field_crop_material.
type FieldCropMaterialColumns struct {
	Id            string //
	FieldCropId   string // 种植记录ID
	MaterialId    string // 物料ID
	OperationType string // 操作类型(施肥/打药)
	Amount        string // 使用数量
	UseTime       string // 使用时间
	Remark        string // 备注
	CreatedAt     string //
}

// fieldCropMaterialColumns holds the columns for the table field_crop_material.
var fieldCropMaterialColumns = FieldCropMaterialColumns{
	Id:            "id",
	FieldCropId:   "field_crop_id",
	MaterialId:    "material_id",
	OperationType: "operation_type",
	Amount:        "amount",
	UseTime:       "use_time",
	Remark:        "remark",
	CreatedAt:     "created_at",
}

// NewFieldCropMaterialDao creates and returns a new DAO object for table data access.
func NewFieldCropMaterialDao(handlers ...gdb.ModelHandler) *FieldCropMaterialDao {
	return &FieldCropMaterialDao{
		group:    "default",
		table:    "field_crop_material",
		columns:  fieldCropMaterialColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FieldCropMaterialDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FieldCropMaterialDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FieldCropMaterialDao) Columns() FieldCropMaterialColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FieldCropMaterialDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FieldCropMaterialDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FieldCropMaterialDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
