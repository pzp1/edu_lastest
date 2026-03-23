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
	ISysField interface {
		List(ctx context.Context, req *system.FieldListReq) *system.FieldListResp
		Delete(ctx context.Context, req *system.DeleteFieldReq) *system.DeleteFieldResp
		AddOrUpdate(ctx context.Context, req *system.AddOrUpdateFieldReq) *system.AddOrUpdateFieldResp
	}
)

var (
	localSysField ISysField
)

func SysField() ISysField {
	if localSysField == nil {
		panic("implement not found for interface ISysField, forgot register?")
	}
	return localSysField
}

func RegisterSysField(i ISysField) {
	localSysField = i
}
