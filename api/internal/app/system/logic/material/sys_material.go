package material

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysMaterial(New())
}

func New() *sSysMaterial {
	return &sSysMaterial{}
}

type sSysMaterial struct{}

/*
物料列表
*/
func (s *sSysMaterial) List(ctx context.Context, req *system.MaterialListReq) *system.MaterialListResp {

	resp := &system.MaterialListResp{}
	var err error

	g.Try(ctx, func(ctx context.Context) {

		model := dao.Material.Ctx(ctx)

		if req.Name != "" {
			model = model.Where("name like ?", "%"+req.Name+"%")
		}

		if req.Type != 0 {
			model = model.Where("type=?", req.Type)
		}

		if req.Category != 0 {
			model = model.Where("category=?", req.Category)
		}

		resp.Total, err = model.Count()
		liberr.ErrIsNil(ctx, err, "获取物料列表失败")

		if req.PageNum == 0 {
			req.PageNum = 1
		}

		resp.CurrentPage = req.PageNum

		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}

		err = model.Page(resp.CurrentPage, req.PageSize).
			Order("id desc").
			Scan(&resp.List)

		liberr.ErrIsNil(ctx, err, "查询物料列表失败")

	})

	return resp
}

/*
删除物料
*/
func (s *sSysMaterial) Delete(ctx context.Context, req *system.DeleteMaterialReq) *system.DeleteMaterialResp {

	resp := &system.DeleteMaterialResp{}

	g.Try(ctx, func(ctx context.Context) {

		_, err := dao.Material.Ctx(ctx).Delete("id=?", req.Id)

		liberr.ErrIsNil(ctx, err, "删除物料失败")

	})

	return resp
}

/*
新增或修改物料
*/
func (s *sSysMaterial) AddOrUpdate(ctx context.Context, req *system.AddOrUpdateMaterialReq) *system.AddOrUpdateMaterialResp {

	resp := &system.AddOrUpdateMaterialResp{}

	g.Try(ctx, func(ctx context.Context) {

		db := dao.Material.Ctx(ctx)

		if req.Id != 0 {

			_, err := db.WherePri(req.Id).Update(req)

			liberr.ErrIsNil(ctx, err, "修改物料失败")

		} else {

			_, err := db.Insert(req)

			liberr.ErrIsNil(ctx, err, "新增物料失败")

		}

	})

	return resp
}
