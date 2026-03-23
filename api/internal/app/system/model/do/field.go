// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Field is the golang structure of table field for DAO operations like Where/Data.
type Field struct {
	g.Meta    `orm:"table:field, do:true"`
	Id        any         //
	UserId    any         // 所属种植户
	Name      any         // 地块名称
	Area      any         // 面积（亩）
	Location  any         // 位置描述
	CreatedAt *gtime.Time //
	Status    any         // 状态 1使用中 0停用
	SortOrder any         // 排序
	UpdatedAt *gtime.Time // 更新时间
	Deleted   any         // 是否删除 0否 1是
}
