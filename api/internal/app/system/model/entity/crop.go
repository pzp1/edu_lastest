// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Crop is the golang structure for table crop.
type Crop struct {
	Id              int64       `json:"id"              orm:"id"                description:""`
	Name            string      `json:"name"            orm:"name"              description:"作物名称"`
	GrowthCycleDays int         `json:"growthCycleDays" orm:"growth_cycle_days" description:"生长周期（天）"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        description:""`
	Category        int         `json:"category"        orm:"category"          description:"种类 1粮食 2水果 3 蔬菜"`
}
