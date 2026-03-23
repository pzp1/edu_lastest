package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

// ===================== 列表 =====================

type PestListReq struct {
	g.Meta `path:"/pest/list" tags:"虫灾管理" method:"get" summary:"虫灾列表"`

	FieldId   int64 `json:"fieldId" form:"fieldId"`
	CropId    int64 `json:"cropId" form:"cropId"`
	Status    int   `json:"status" form:"status"` // 0未解决 1已解决
	CreatedBy int   `json:"createdBy"`            // 0未解决 1已解决`
	commonApi.PageReq
	commonApi.Author
}

type PestListResp struct {
	g.Meta `mime:"application/json"`

	commonApi.ListRes

	List []*model.PestListResp `json:"list"`
}

// ===================== 删除 =====================

type DeletePestReq struct {
	g.Meta `path:"/pest/delete" tags:"虫灾管理" method:"post" summary:"删除虫灾"`

	Id int64 `json:"id" v:"required#ID不能为空"`
}

type DeletePestResp struct {
	g.Meta `mime:"application/json"`
}

// ===================== 新增 / 修改 =====================

type AddPestReq struct {
	g.Meta `path:"/pest/add" tags:"虫灾管理" method:"post" summary:"新增或修改虫灾"`

	FieldId     int64  `json:"fieldId" v:"required#地块不能为空"`
	CropId      int64  `json:"cropId" v:"required#作物不能为空"`
	Description string `json:"description" v:"required#虫灾描述不能为空"`
	Status      int    `json:"status"`    // 0未解决 1已解决
	CreatedBy   int    `json:"createdBy"` // 0未解决 1已解决`
	FieldName   string `json:"fieldName"`
	CropName    string `json:"cropName"`
}

type AddOrUpdatePestResp struct {
	g.Meta `mime:"application/json"`
}

// ===================== 修改状态 =====================

type UpdatePestStatusReq struct {
	g.Meta      `path:"/pest/updateStatus" tags:"虫灾管理" method:"post" summary:"修改虫灾状态"`
	Id          int64  `json:"id" v:"required#ID不能为空"`
	Status      int    `json:"status" v:"required#状态不能为空"`
	SloveAnswer string `json:"sloveAnswer"`
	SolveBY     string `json:"solveBY"`
}

type UpdatePestStatusResp struct {
	g.Meta `mime:"application/json"`
}
