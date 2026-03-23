package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type MaterialListReq struct {
	g.Meta `path:"/material/list" tags:"物料管理" method:"get" summary:"物料列表"`

	Name     string `json:"name" form:"name"`
	Type     int    `json:"type" form:"type"`         // 1农药 2化肥
	Category int    `json:"category" form:"category"` // 分类

	commonApi.PageReq
	commonApi.Author
}

type MaterialListResp struct {
	g.Meta `mime:"application/json"`

	commonApi.ListRes

	List []*model.MaterialListResp `json:"list"`
}

type DeleteMaterialReq struct {
	g.Meta `path:"/material/delete" tags:"物料管理" method:"post" summary:"删除物料"`

	Id int64 `json:"id" v:"required#ID不能为空"`
}

type DeleteMaterialResp struct {
	g.Meta `mime:"application/json"`
}

type AddOrUpdateMaterialReq struct {
	g.Meta `path:"/material/addOrUpdate" tags:"物料管理" method:"post" summary:"新增或修改物料"`

	Id              int64  `json:"id"`
	Name            string `json:"name"`
	Type            int    `json:"type"`
	Category        int    `json:"category"`
	UsageMethod     string `json:"usageMethod"`
	NutrientContent string `json:"nutrientContent"`
}

type AddOrUpdateMaterialResp struct {
	g.Meta `mime:"application/json"`
}
