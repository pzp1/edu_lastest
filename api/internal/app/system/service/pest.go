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
	ISysPest interface {
		/*
			虫灾列表
		*/
		List(ctx context.Context, req *system.PestListReq) *system.PestListResp
		/*
			新增或修改虫灾
		*/
		Add(ctx context.Context, req *system.AddPestReq) *system.AddOrUpdatePestResp
		/*
			删除虫灾
		*/
		Delete(ctx context.Context, req *system.DeletePestReq) *system.DeletePestResp
		/*
			修改虫灾状态
		*/
		UpdateStatus(ctx context.Context, req *system.UpdatePestStatusReq) *system.UpdatePestStatusResp
	}
)

var (
	localSysPest ISysPest
)

func SysPest() ISysPest {
	if localSysPest == nil {
		panic("implement not found for interface ISysPest, forgot register?")
	}
	return localSysPest
}

func RegisterSysPest(i ISysPest) {
	localSysPest = i
}
