package syscrop

import (
	"context"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterSysCrop(New())
}

func New() *sSysCrop {
	return &sSysCrop{}
}

type sSysCrop struct{}

/*
作物列表
*/
func (s *sSysCrop) List(ctx context.Context, req *system.CropListReq) *system.CropListResp {

	resp := &system.CropListResp{}
	var err error

	g.Try(ctx, func(ctx context.Context) {

		model := dao.Crop.Ctx(ctx)

		if req.Name != "" {
			model = model.Where("name like ?", "%"+req.Name+"%")
		}

		if req.Category != 0 {
			model = model.Where("category=?", req.Category)
		}

		resp.Total, err = model.Count()
		liberr.ErrIsNil(ctx, err, "获取作物列表失败")

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

		liberr.ErrIsNil(ctx, err, "查询作物列表失败")
	})

	return resp
}

/*
删除作物
*/
func (s *sSysCrop) Delete(ctx context.Context, req *system.DeleteCropReq) *system.DeleteCropResp {

	resp := &system.DeleteCropResp{}

	g.Try(ctx, func(ctx context.Context) {

		_, err := dao.Crop.Ctx(ctx).Delete("id=?", req.Id)

		liberr.ErrIsNil(ctx, err, "删除作物失败")
	})

	return resp
}

/*
新增或修改作物
*/
func (s *sSysCrop) AddOrUpdate(ctx context.Context, req *system.AddOrUpdateCropReq) *system.AddOrUpdateCropResp {

	resp := &system.AddOrUpdateCropResp{}

	g.Try(ctx, func(ctx context.Context) {

		db := dao.Crop.Ctx(ctx)

		if req.Id != 0 {

			_, err := db.WherePri(req.Id).Update(req)

			liberr.ErrIsNil(ctx, err, "修改作物失败")

		} else {

			_, err := db.Insert(req)

			liberr.ErrIsNil(ctx, err, "新增作物失败")

		}

	})

	return resp
}
