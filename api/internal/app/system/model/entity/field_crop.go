// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FieldCrop is the golang structure for table field_crop.
type FieldCrop struct {
	Id                  int64       `json:"id"                  orm:"id"                    description:""`
	FieldId             int64       `json:"fieldId"             orm:"field_id"              description:"地块ID"`
	CropId              int64       `json:"cropId"              orm:"crop_id"               description:"作物ID"`
	PlantDate           *gtime.Time `json:"plantDate"           orm:"plant_date"            description:"种植日期"`
	ExpectedHarvestDate *gtime.Time `json:"expectedHarvestDate" orm:"expected_harvest_date" description:"预计收获日期"`
	HarvestDate         *gtime.Time `json:"harvestDate"         orm:"harvest_date"          description:"实际收获日期"`
	CurrentStage        string      `json:"currentStage"        orm:"current_stage"         description:"当前生长期"`
	Status              int         `json:"status"              orm:"status"                description:"1种植中 2已收获"`
	Remark              string      `json:"remark"              orm:"remark"                description:"备注"`
	CreatedAt           *gtime.Time `json:"createdAt"           orm:"created_at"            description:""`
	UpdatedAt           *gtime.Time `json:"updatedAt"           orm:"updated_at"            description:""`
	FieldName           string      `json:"fieldName"           orm:"field_name"            description:"地块名称"`
	CropName            string      `json:"cropName"            orm:"crop_name"             description:"作物名称"`
}
