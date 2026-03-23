package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
)

type FieldCropListReq struct {
	g.Meta `path:"/fieldCrop/list" tags:"种植记录" method:"get" summary:"种植记录列表"`

	FieldId      int64 `json:"fieldId" form:"fieldId"`
	CropId       int64 `json:"cropId" form:"cropId"`
	Status       int   `json:"status" form:"status"`
	UserID       int64 `json:"userId" form:"userId" validate:"required"` // 用户ID
	ExpectedFlag int64 `json:"expectedFlag"`
	commonApi.PageReq
	commonApi.Author
}

type FieldCropListResp struct {
	g.Meta `mime:"application/json"`

	commonApi.ListRes

	List []*model.FieldCropListResp `json:"list"`
}

type AddOrUpdateFieldCropReq struct {
	g.Meta `path:"/fieldCrop/addOrUpdate" tags:"种植记录" method:"post" summary:"新增或修改种植记录"`

	Id                  int64  `json:"id"`
	FieldId             int64  `json:"fieldId"`
	CropId              int64  `json:"cropId"`
	PlantDate           string `json:"plantDate"`
	ExpectedHarvestDate string `json:"expectedHarvestDate"`
	HarvestDate         string `json:"harvestDate"`
	CurrentStage        string `json:"currentStage"`
	Status              int    `json:"status"`
	Remark              string `json:"remark"`
	FieldName           string `json:"fieldName" dc:"地块名称"`
	CropName            string `json:"cropName" dc:"作物名称"`
}
type AddOrUpdateFieldCropResp struct{}

type DeleteFieldCropReq struct {
	g.Meta `path:"/fieldCrop/delete" tags:"种植记录" method:"post" summary:"删除种植记录"`

	Id int64 `json:"id" v:"required#ID不能为空"`
}

type DeleteFieldCropResp struct {
	g.Meta `mime:"application/json"`
}

type FieldCropDetailReq struct {
	g.Meta `path:"/fieldCrop/detail" tags:"种植记录" method:"get" summary:"种植记录详情"`

	Id int64 `json:"id"`
}

type FieldCropDetailResp struct {
	g.Meta `mime:"application/json"`

	model.FieldCrop
}
