// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Material is the golang structure of table material for DAO operations like Where/Data.
type Material struct {
	g.Meta          `orm:"table:material, do:true"`
	Id              any         //
	Name            any         // 物料名称
	Type            any         // 物料类型 1农药 2化肥
	Category        any         // 分类
	UsageMethod     any         // 使用方式
	NutrientContent any         // 养分含量
	CreatedAt       *gtime.Time //
}
