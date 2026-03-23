package sysfield

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysField(New())
}

func New() *sSysField {
	return &sSysField{}
}

type sSysField struct {
}

func (s *sSysField) List(ctx context.Context, req *system.FieldListReq) *system.FieldListResp {
	resp := &system.FieldListResp{}
	var err error
	g.Try(ctx, func(ctx context.Context) {
		model := dao.Field.Ctx(ctx).WithAll()
		if req.UserID == 0 {
			liberr.ErrIsNil(ctx, gerror.New("用户id不能为空"), "")
		}
		userModel := dao.SysUser.Ctx(ctx)
		userData := entity.SysUser{}
		err = userModel.Where("id=?", req.UserID).Scan(&userData)
		liberr.ErrIsNil(ctx, err, "获取用户数据失败")
		if userData.IsAdmin == 1 {
		} else {
			if req.UserID != 0 {
				model = model.Where("user_id = ?", req.UserID)
			}
		}

		if req.Status != 0 {
			model = model.Where("status=?", req.Status)
		}
		if req.Name != "" {
			model = model.Where("name like ?", "%"+req.Name+"%")
		}
		resp.Total, err = model.Count()
		liberr.ErrIsNil(ctx, err, "获取地块列表数据失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		resp.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		err = model.Page(resp.CurrentPage, req.PageSize).Order("sort_order desc").Scan(&resp.List)

	})
	return resp
}

func (s *sSysField) Delete(ctx context.Context, req *system.DeleteFieldReq) *system.DeleteFieldResp {
	resp := &system.DeleteFieldResp{}
	g.Try(ctx, func(ctx context.Context) {
		_, err := dao.Field.Ctx(ctx).Delete("id=?", req.FieldId)
		liberr.ErrIsNil(ctx, err, "删除地块数据失败")
	})
	return resp
}

func (s *sSysField) AddOrUpdate(ctx context.Context, req *system.AddOrUpdateFieldReq) *system.AddOrUpdateFieldResp {
	resp := &system.AddOrUpdateFieldResp{}
	g.Try(ctx, func(ctx context.Context) {
		db := dao.Field.Ctx(ctx)
		//fieldData := &entity.Field{}
		//copier.Copy(fieldData, req)
		if req.Id != 0 {
			_, err := db.WherePri(req.Id).Update(req)
			liberr.ErrIsNil(ctx, err, "修改地块数据失败")
		} else {
			_, err := db.Insert(req)
			liberr.ErrIsNil(ctx, err, "添加地块数据失败")
		}
	})
	return resp
}
