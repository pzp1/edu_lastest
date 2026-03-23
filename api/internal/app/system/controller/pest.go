package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

type cPest struct{}

var Pest = cPest{}

/*
虫灾列表
*/
func (c *cPest) List(ctx context.Context, req *system.PestListReq) (res *system.PestListResp, err error) {
	res = service.SysPest().List(ctx, req)
	return
}

/*
新增或修改虫灾
*/
func (c *cPest) Add(ctx context.Context, req *system.AddPestReq) (res *system.AddOrUpdatePestResp, err error) {
	res = service.SysPest().Add(ctx, req)
	return
}

/*
删除虫灾
*/
func (c *cPest) Delete(ctx context.Context, req *system.DeletePestReq) (res *system.DeletePestResp, err error) {
	res = service.SysPest().Delete(ctx, req)
	return
}

/*
修改虫灾状态
*/
func (c *cPest) UpdateStatus(ctx context.Context, req *system.UpdatePestStatusReq) (res *system.UpdatePestStatusResp, err error) {
	res = service.SysPest().UpdateStatus(ctx, req)
	return
}
