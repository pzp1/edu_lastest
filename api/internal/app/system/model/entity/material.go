// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Material is the golang structure for table material.
type Material struct {
	Id              int64       `json:"id"              orm:"id"               description:""`
	Name            string      `json:"name"            orm:"name"             description:"物料名称"`
	Type            int         `json:"type"            orm:"type"             description:"物料类型 1农药 2化肥"`
	Category        int         `json:"category"        orm:"category"         description:"分类"`
	UsageMethod     string      `json:"usageMethod"     orm:"usage_method"     description:"使用方式"`
	NutrientContent string      `json:"nutrientContent" orm:"nutrient_content" description:"养分含量"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:""`
}
