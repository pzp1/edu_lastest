// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FieldCropMaterial is the golang structure of table field_crop_material for DAO operations like Where/Data.
type FieldCropMaterial struct {
	g.Meta        `orm:"table:field_crop_material, do:true"`
	Id            any         //
	FieldCropId   any         // 种植记录ID
	MaterialId    any         // 物料ID
	OperationType any         // 操作类型(施肥/打药)
	Amount        any         // 使用数量
	UseTime       *gtime.Time // 使用时间
	Remark        any         // 备注
	CreatedAt     *gtime.Time //
}
