// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/system"
)

type (
	ISysCrop interface {
		/*
			作物列表
		*/
		List(ctx context.Context, req *system.CropListReq) *system.CropListResp
		/*
			删除作物
		*/
		Delete(ctx context.Context, req *system.DeleteCropReq) *system.DeleteCropResp
		/*
			新增或修改作物
		*/
		AddOrUpdate(ctx context.Context, req *system.AddOrUpdateCropReq) *system.AddOrUpdateCropResp
	}
)

var (
	localSysCrop ISysCrop
)

func SysCrop() ISysCrop {
	if localSysCrop == nil {
		panic("implement not found for interface ISysCrop, forgot register?")
	}
	return localSysCrop
}

func RegisterSysCrop(i ISysCrop) {
	localSysCrop = i
}
