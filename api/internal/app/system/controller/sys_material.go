package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var Material = materialController{}

type materialController struct {
	BaseController
}

// List 地块列表
func (c *materialController) List(ctx context.Context, req *system.MaterialListReq) (res *system.MaterialListResp, err error) {
	res = service.SysMaterial().List(ctx, req)
	return
}

func (c *materialController) Delete(ctx context.Context, req *system.DeleteMaterialReq) (res *system.DeleteMaterialResp, err error) {
	res = service.SysMaterial().Delete(ctx, req)
	return
}

func (c *materialController) AddOrUpdate(ctx context.Context, req *system.AddOrUpdateMaterialReq) (res *system.AddOrUpdateMaterialResp, err error) {
	res = service.SysMaterial().AddOrUpdate(ctx, req)
	return
}
