// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Field is the golang structure for table field.
type Field struct {
	Id        int64       `json:"id"        orm:"id"         description:""`
	UserId    int64       `json:"userId"    orm:"user_id"    description:"所属种植户"`
	Name      string      `json:"name"      orm:"name"       description:"地块名称"`
	Area      float64     `json:"area"      orm:"area"       description:"面积（亩）"`
	Location  string      `json:"location"  orm:"location"   description:"位置描述"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	Status    int         `json:"status"    orm:"status"     description:"状态 1使用中 0停用"`
	SortOrder int         `json:"sortOrder" orm:"sort_order" description:"排序"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	Deleted   int         `json:"deleted"   orm:"deleted"    description:"是否删除 0否 1是"`
}