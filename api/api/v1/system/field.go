package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type FieldListReq struct {
	g.Meta `path:"/field/list" tags:"分类信息" method:"get" summary:"分类信息列表"`
	Name   string `json:"name" form:"name"`                         // 地块名称
	Status int    `json:"status" form:"status"`                     // 状态 1使用中 0停用
	UserID int64  `json:"userId" form:"userId" validate:"required"` // 用户ID
	commonApi.PageReq
	commonApi.Author
}

type FieldListResp struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.FieldListResp `json:"list"`
}

type DeleteFieldReq struct {
	g.Meta  `path:"/field/delete" tags:"删除地块" method:"delete" summary:"删除地块"`
	FieldId int `json:"fieldId"`
}
type DeleteFieldResp struct{}

type AddOrUpdateFieldReq struct {
	g.Meta    `path:"/field/addOrUp" tags:"分类信息" method:"post" summary:"新增或删除地块"`
	Id        int64  `json:"id"        dc:"地块ID"`
	UserId    int64  `json:"userId"    v:"required#请选择耕种人" dc:"耕种人"`
	Name      string `json:"name"      v:"required#请输入地块名称" dc:"地块名称"`
	Area      int64  `json:"area"      dc:"面积"`
	Location  string `json:"location"  dc:"位置"`
	Status    int    `json:"status"    dc:"状态"`
	SortOrder int    `json:"sortOrder" dc:"排序"`
}

type AddOrUpdateFieldResp struct{}
