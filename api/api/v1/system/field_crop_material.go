package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
)

type MaterialUseListReq struct {
	g.Meta      `path:"/materialUse/list" method:"get" summary:"物料使用列表"`
	FieldCropId int64 `json:"fieldCropId" dc:"种植记录ID"`
}

type MaterialUseListResp struct {
	List []MaterialUseItem `json:"list"`

	commonApi.ListRes
}

type MaterialUseItem struct {
	Id            int64   `json:"id"`
	FieldCropId   int64   `json:"fieldCropId"`
	MaterialId    int64   `json:"materialId"`
	MaterialName  string  `json:"materialName"`
	OperationType string  `json:"operationType"`
	Amount        float64 `json:"amount"`
	UseTime       string  `json:"useTime"`
	Remark        string  `json:"remark"`
}

type MaterialUseAddReq struct {
	g.Meta        `path:"/materialUse/add" method:"post" summary:"新增物料使用"`
	FieldCropId   int64  `json:"fieldCropId"`
	MaterialId    int64  `json:"materialId"`
	OperationType string `json:"operationType"`
	UseTime       string `json:"useTime"`
	Remark        string `json:"remark"`
	Amount        int64  `json:"amount"`
}

type MaterialUseAddResp struct{}

type MaterialUseUpdateReq struct {
	g.Meta        `path:"/materialUse/update" method:"post" summary:"修改物料使用"`
	Id            int64  `json:"id"`
	MaterialId    int64  `json:"materialId"`
	OperationType string `json:"operationType"`
	UseTime       string `json:"useTime"`
	Remark        string `json:"remark"`
	Amount        int64  `json:"amount"`
}

type MaterialUseUpdateResp struct{}

type MaterialUseDeleteReq struct {
	g.Meta `path:"/materialUse/delete" method:"post" summary:"删除物料使用"`
	Id     int64 `json:"id"`
}

type MaterialUseDeleteResp struct{}
