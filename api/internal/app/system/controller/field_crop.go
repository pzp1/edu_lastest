package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var FieldCrop = FieldCropController{}

type FieldCropController struct {
	BaseController
}

// List 种植记录列表
func (c *FieldCropController) List(ctx context.Context, req *system.FieldCropListReq) (res *system.FieldCropListResp, err error) {
	res = service.SysFieldCrop().List(ctx, req)
	return
}

// Delete 删除种植记录
func (c *FieldCropController) Delete(ctx context.Context, req *system.DeleteFieldCropReq) (res *system.DeleteFieldCropResp, err error) {
	res = service.SysFieldCrop().Delete(ctx, req)
	return
}

// AddOrUpdate 新增或修改种植记录
func (c *FieldCropController) AddOrUpdate(ctx context.Context, req *system.AddOrUpdateFieldCropReq) (res *system.AddOrUpdateFieldCropResp, err error) {
	res = service.SysFieldCrop().AddOrUpdate(ctx, req)
	return
}

// Detail 种植记录详情
func (c *FieldCropController) Detail(ctx context.Context, req *system.FieldCropDetailReq) (res *system.FieldCropDetailResp, err error) {
	res = service.SysFieldCrop().Detail(ctx, req)
	return
}
