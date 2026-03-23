package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

/*
*
作物列表
*/
type CropListReq struct {
	g.Meta `path:"/crop/list" tags:"作物管理" method:"get" summary:"作物列表"`

	Name     string `json:"name" form:"name"`         // 作物名称
	Category int    `json:"category" form:"category"` // 分类 1粮食 2水果 3蔬菜

	commonApi.PageReq
	commonApi.Author
}

type CropListResp struct {
	g.Meta `mime:"application/json"`

	commonApi.ListRes

	List []*model.CropListResp `json:"list"`
}

/*
*
删除作物
*/
type DeleteCropReq struct {
	g.Meta `path:"/crop/delete" tags:"作物管理" method:"delete" summary:"删除作物"`

	Id int64 `json:"id" dc:"作物ID"`
}

type DeleteCropResp struct{}

/*
*
新增 / 修改作物
*/
type AddOrUpdateCropReq struct {
	g.Meta `path:"/crop/addOrUp" tags:"作物管理" method:"post" summary:"新增或修改作物"`

	Id              int64  `json:"id" dc:"作物ID"`
	Name            string `json:"name" v:"required#请输入作物名称" dc:"作物名称"`
	GrowthCycleDays int    `json:"growthCycleDays" dc:"生长周期（天）"`
	Category        int    `json:"category" dc:"分类 1粮食 2水果 3蔬菜"`
}

type AddOrUpdateCropResp struct{}
