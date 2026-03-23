// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Crop is the golang structure of table crop for DAO operations like Where/Data.
type Crop struct {
	g.Meta          `orm:"table:crop, do:true"`
	Id              any         //
	Name            any         // 作物名称
	GrowthCycleDays any         // 生长周期（天）
	CreatedAt       *gtime.Time //
	Category        any         // 种类 1粮食 2水果 3 蔬菜
}
