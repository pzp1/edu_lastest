package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var Crop = cropController{}

type cropController struct {
	BaseController
}

// List 作物列表
func (c *cropController) List(ctx context.Context, req *system.CropListReq) (res *system.CropListResp, err error) {
	res = service.SysCrop().List(ctx, req)
	return
}

// Delete 删除作物
func (c *cropController) Delete(ctx context.Context, req *system.DeleteCropReq) (res *system.DeleteCropResp, err error) {
	res = service.SysCrop().Delete(ctx, req)
	return
}

// AddOrUpdate 新增或修改作物
func (c *cropController) AddOrUpdate(ctx context.Context, req *system.AddOrUpdateCropReq) (res *system.AddOrUpdateCropResp, err error) {
	res = service.SysCrop().AddOrUpdate(ctx, req)
	return
}
