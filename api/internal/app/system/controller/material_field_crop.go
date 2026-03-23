package controller

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

type cMaterialUse struct{}

var MaterialUse = cMaterialUse{}

/*
物料使用列表
*/
func (c *cMaterialUse) List(ctx context.Context, req *system.MaterialUseListReq) (res *system.MaterialUseListResp, err error) {

	res = service.SysFieldCropMaterial().List(ctx, req)

	return
}

/*
新增或修改物料使用
*/
func (c *cMaterialUse) AddOrUpdate(ctx context.Context, req *system.MaterialUseAddReq) (res *system.MaterialUseAddResp, err error) {

	res = service.SysFieldCropMaterial().AddOrUpdate(ctx, req)

	return
}

/*
修改物料使用
*/
func (c *cMaterialUse) Update(ctx context.Context, req *system.MaterialUseUpdateReq) (res *system.MaterialUseUpdateResp, err error) {

	res = service.SysFieldCropMaterial().Update(ctx, req)

	return
}

/*
删除物料使用
*/
func (c *cMaterialUse) Delete(ctx context.Context, req *system.MaterialUseDeleteReq) (res *system.MaterialUseDeleteReq, err error) {

	res = service.SysFieldCropMaterial().Delete(ctx, req)

	return
}
