// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FieldCropMaterial is the golang structure for table field_crop_material.
type FieldCropMaterial struct {
	Id            int64       `json:"id"            orm:"id"             description:""`
	FieldCropId   int64       `json:"fieldCropId"   orm:"field_crop_id"  description:"种植记录ID"`
	MaterialId    int64       `json:"materialId"    orm:"material_id"    description:"物料ID"`
	OperationType string      `json:"operationType" orm:"operation_type" description:"操作类型(施肥/打药)"`
	Amount        float64     `json:"amount"        orm:"amount"         description:"使用数量"`
	UseTime       *gtime.Time `json:"useTime"       orm:"use_time"       description:"使用时间"`
	Remark        string      `json:"remark"        orm:"remark"         description:"备注"`
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:""`
}
