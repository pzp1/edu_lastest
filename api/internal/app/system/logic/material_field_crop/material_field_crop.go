package field_crop_material

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() { service.RegisterSysFieldCropMaterial(NewSysFieldCropMaterial()) }

type sSysFieldCropMaterial struct{}

func NewSysFieldCropMaterial() *sSysFieldCropMaterial {
	return &sSysFieldCropMaterial{}
}

/*
物料使用记录列表
*/
func (s *sSysFieldCropMaterial) List(ctx context.Context, req *system.MaterialUseListReq) *system.MaterialUseListResp {
	resp := &system.MaterialUseListResp{}
	var err error

	g.Try(ctx, func(ctx context.Context) {

		model := dao.FieldCropMaterial.Ctx(ctx).
			LeftJoin("material m", "m.id = field_crop_material.material_id").
			Fields("field_crop_material.*,m.name materialName")
		countModel := dao.FieldCropMaterial.Ctx(ctx)
		if req.FieldCropId != 0 {
			model = model.Where("field_crop_id=?", req.FieldCropId)
			countModel = countModel.Where("field_crop_id=?", req.FieldCropId)
		}

		resp.Total, err = countModel.Count()
		liberr.ErrIsNil(ctx, err, "获取物料使用记录数量失败")
		err = model.Page(resp.CurrentPage, 1000).
			Order("id desc").
			Scan(&resp.List)

		liberr.ErrIsNil(ctx, err, "查询物料使用记录失败")

	})

	return resp
}

/*
新增或修改物料使用记录
*/
func (s *sSysFieldCropMaterial) AddOrUpdate(ctx context.Context, req *system.MaterialUseAddReq) *system.MaterialUseAddResp {

	resp := &system.MaterialUseAddResp{}

	g.Try(ctx, func(ctx context.Context) {

		db := dao.FieldCropMaterial.Ctx(ctx)
		_, err := db.Insert(req)
		liberr.ErrIsNil(ctx, err, "新增物料使用记录失败")

	})

	return resp
}

/*
删除物料使用记录
*/
func (s *sSysFieldCropMaterial) Delete(ctx context.Context, req *system.MaterialUseDeleteReq) *system.MaterialUseDeleteReq {

	resp := &system.MaterialUseDeleteReq{}

	g.Try(ctx, func(ctx context.Context) {

		_, err := dao.FieldCropMaterial.Ctx(ctx).
			Delete("id=?", req.Id)

		liberr.ErrIsNil(ctx, err, "删除物料使用记录失败")

	})

	return resp
}

/*
修改物料使用记录
*/
func (s *sSysFieldCropMaterial) Update(ctx context.Context, req *system.MaterialUseUpdateReq) *system.MaterialUseUpdateResp {

	resp := &system.MaterialUseUpdateResp{}

	g.Try(ctx, func(ctx context.Context) {

		db := dao.FieldCropMaterial.Ctx(ctx)

		// 更新记录
		_, err := db.WherePri(req.Id).Update(req)

		liberr.ErrIsNil(ctx, err, "修改物料使用记录失败")

	})

	return resp
}
