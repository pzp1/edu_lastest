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
	ISysMaterial interface {
		/*
			物料列表
		*/
		List(ctx context.Context, req *system.MaterialListReq) *system.MaterialListResp
		/*
			删除物料
		*/
		Delete(ctx context.Context, req *system.DeleteMaterialReq) *system.DeleteMaterialResp
		/*
			新增或修改物料
		*/
		AddOrUpdate(ctx context.Context, req *system.AddOrUpdateMaterialReq) *system.AddOrUpdateMaterialResp
	}
)

var (
	localSysMaterial ISysMaterial
)

func SysMaterial() ISysMaterial {
	if localSysMaterial == nil {
		panic("implement not found for interface ISysMaterial, forgot register?")
	}
	return localSysMaterial
}

func RegisterSysMaterial(i ISysMaterial) {
	localSysMaterial = i
}
