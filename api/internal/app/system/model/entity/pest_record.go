// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PestRecord is the golang structure for table pest_record.
type PestRecord struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键ID"`
	FieldId     int64       `json:"fieldId"     orm:"field_id"     description:"地块ID"`
	FieldName   string      `json:"fieldName"   orm:"field_name"   description:"地块名称"`
	CropId      int64       `json:"cropId"      orm:"crop_id"      description:"作物ID"`
	CropName    string      `json:"cropName"    orm:"crop_name"    description:"作物名称"`
	Description string      `json:"description" orm:"description"  description:"虫灾描述"`
	Status      int         `json:"status"      orm:"status"       description:"状态 0未解决 1已解决"`
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建人"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	SloveAnswer string      `json:"sloveAnswer" orm:"slove_answer" description:"解决办法"`
	SolveBy     int         `json:"solveBy"     orm:"solve_by"     description:"解决人"`
}
