package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
)

var Field = fieldController{}

type fieldController struct {
	BaseController
}

// List 地块列表
func (c *fieldController) List(ctx context.Context, req *system.FieldListReq) (res *system.FieldListResp, err error) {
	res = service.SysField().List(ctx, req)
	return
}

func (c *fieldController) Delete(ctx context.Context, req *system.DeleteFieldReq) (res *system.DeleteFieldResp, err error) {
	res = service.SysField().Delete(ctx, req)
	return
}

func (c *fieldController) AddOrUpdate(ctx context.Context, req *system.AddOrUpdateFieldReq) (res *system.AddOrUpdateFieldResp, err error) {
	res = service.SysField().AddOrUpdate(ctx, req)
	return
}
