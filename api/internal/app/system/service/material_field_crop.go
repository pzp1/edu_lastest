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
	ISysFieldCropMaterial interface {
		/*
			物料使用记录列表
		*/
		List(ctx context.Context, req *system.MaterialUseListReq) *system.MaterialUseListResp
		/*
			新增或修改物料使用记录
		*/
		AddOrUpdate(ctx context.Context, req *system.MaterialUseAddReq) *system.MaterialUseAddResp
		/*
			删除物料使用记录
		*/
		Delete(ctx context.Context, req *system.MaterialUseDeleteReq) *system.MaterialUseDeleteReq
		/*
			修改物料使用记录
		*/
		Update(ctx context.Context, req *system.MaterialUseUpdateReq) *system.MaterialUseUpdateResp
	}
)

var (
	localSysFieldCropMaterial ISysFieldCropMaterial
)

func SysFieldCropMaterial() ISysFieldCropMaterial {
	if localSysFieldCropMaterial == nil {
		panic("implement not found for interface ISysFieldCropMaterial, forgot register?")
	}
	return localSysFieldCropMaterial
}

func RegisterSysFieldCropMaterial(i ISysFieldCropMaterial) {
	localSysFieldCropMaterial = i
}
