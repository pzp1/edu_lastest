// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FieldCrop is the golang structure of table field_crop for DAO operations like Where/Data.
type FieldCrop struct {
	g.Meta              `orm:"table:field_crop, do:true"`
	Id                  any         //
	FieldId             any         // 地块ID
	CropId              any         // 作物ID
	PlantDate           *gtime.Time // 种植日期
	ExpectedHarvestDate *gtime.Time // 预计收获日期
	HarvestDate         *gtime.Time // 实际收获日期
	CurrentStage        any         // 当前生长期
	Status              any         // 1种植中 2已收获
	Remark              any         // 备注
	CreatedAt           *gtime.Time //
	UpdatedAt           *gtime.Time //
	FieldName           any         // 地块名称
	CropName            any         // 作物名称
}
