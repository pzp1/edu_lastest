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
	ISysFieldCrop interface {
		/*
			种植记录列表
		*/
		List(ctx context.Context, req *system.FieldCropListReq) *system.FieldCropListResp
		/*
			新增或修改种植记录
		*/
		AddOrUpdate(ctx context.Context, req *system.AddOrUpdateFieldCropReq) *system.AddOrUpdateFieldCropResp
		/*
			删除种植记录
		*/
		Delete(ctx context.Context, req *system.DeleteFieldCropReq) *system.DeleteFieldCropResp
		/*
			种植记录详情
		*/
		Detail(ctx context.Context, req *system.FieldCropDetailReq) *system.FieldCropDetailResp
	}
)

var (
	localSysFieldCrop ISysFieldCrop
)

func SysFieldCrop() ISysFieldCrop {
	if localSysFieldCrop == nil {
		panic("implement not found for interface ISysFieldCrop, forgot register?")
	}
	return localSysFieldCrop
}

func RegisterSysFieldCrop(i ISysFieldCrop) {
	localSysFieldCrop = i
}
