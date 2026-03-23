// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PestRecord is the golang structure of table pest_record for DAO operations like Where/Data.
type PestRecord struct {
	g.Meta      `orm:"table:pest_record, do:true"`
	Id          any         // 主键ID
	FieldId     any         // 地块ID
	FieldName   any         // 地块名称
	CropId      any         // 作物ID
	CropName    any         // 作物名称
	Description any         // 虫灾描述
	Status      any         // 状态 0未解决 1已解决
	CreatedBy   any         // 创建人
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	SloveAnswer any         // 解决办法
	SolveBy     any         // 解决人
}
