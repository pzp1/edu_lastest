// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FieldCropDao is the data access object for the table field_crop.
type FieldCropDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  FieldCropColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// FieldCropColumns defines and stores column names for the table field_crop.
type FieldCropColumns struct {
	Id                  string //
	FieldId             string // 地块ID
	CropId              string // 作物ID
	PlantDate           string // 种植日期
	ExpectedHarvestDate string // 预计收获日期
	HarvestDate         string // 实际收获日期
	CurrentStage        string // 当前生长期
	Status              string // 1种植中 2已收获
	Remark              string // 备注
	CreatedAt           string //
	UpdatedAt           string //
	FieldName           string // 地块名称
	CropName            string // 作物名称
}

// fieldCropColumns holds the columns for the table field_crop.
var fieldCropColumns = FieldCropColumns{
	Id:                  "id",
	FieldId:             "field_id",
	CropId:              "crop_id",
	PlantDate:           "plant_date",
	ExpectedHarvestDate: "expected_harvest_date",
	HarvestDate:         "harvest_date",
	CurrentStage:        "current_stage",
	Status:              "status",
	Remark:              "remark",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	FieldName:           "field_name",
	CropName:            "crop_name",
}

// NewFieldCropDao creates and returns a new DAO object for table data access.
func NewFieldCropDao(handlers ...gdb.ModelHandler) *FieldCropDao {
	return &FieldCropDao{
		group:    "default",
		table:    "field_crop",
		columns:  fieldCropColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FieldCropDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FieldCropDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FieldCropDao) Columns() FieldCropColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FieldCropDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FieldCropDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FieldCropDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
