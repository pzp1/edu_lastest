package model

import "github.com/gogf/gf/v2/os/gtime"

type FieldCrop struct {
	Id                  int64       `json:"id"                  orm:"id"                   description:"ID"`
	FieldId             int64       `json:"fieldId"             orm:"field_id"             description:"地块ID"`
	CropId              int64       `json:"cropId"              orm:"crop_id"              description:"作物ID"`
	FieldName           string      `json:"fieldName"           orm:"field_name"            description:"地块名称"`
	CropName            string      `json:"cropName"            orm:"crop_name"             description:"作物名称"`
	PlantDate           *gtime.Time `json:"plantDate"           orm:"plant_date"           description:"种植日期"`
	ExpectedHarvestDate *gtime.Time `json:"expectedHarvestDate" orm:"expected_harvest_date" description:"预计收获日期"`
	HarvestDate         *gtime.Time `json:"harvestDate"         orm:"harvest_date"         description:"实际收获日期"`
	CurrentStage        string      `json:"currentStage"        orm:"current_stage"        description:"当前生长期"`
	Status              int         `json:"status"              orm:"status"               description:"状态"`
	Remark              string      `json:"remark"              orm:"remark"               description:"备注"`
	CreatedAt           *gtime.Time `json:"createdAt"           orm:"created_at"`
	UpdatedAt           *gtime.Time `json:"updatedAt"           orm:"updated_at"`
}

type FieldCropListResp struct {
	Id      int64 `json:"id"`
	FieldId int64 `json:"fieldId"`
	CropId  int64 `json:"cropId"`

	FieldName string `json:"fieldName"` // 地块名称
	CropName  string `json:"cropName"`  // 作物名称

	PlantDate           *gtime.Time `json:"plantDate"`
	ExpectedHarvestDate *gtime.Time `json:"expectedHarvestDate"`
	HarvestDate         *gtime.Time `json:"harvestDate"`

	CurrentStage string `json:"currentStage"`

	Status int `json:"status"`

	Remark string `json:"remark"`

	CreatedAt *gtime.Time `json:"createdAt"`
}
