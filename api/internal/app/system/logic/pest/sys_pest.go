package pest

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() { service.RegisterSysPest(NewSysPest()) }

type sSysPest struct{}

func NewSysPest() *sSysPest {
	return &sSysPest{}
}

/*
虫灾列表
*/
func (s *sSysPest) List(ctx context.Context, req *system.PestListReq) *system.PestListResp {

	resp := &system.PestListResp{}
	var err error

	g.Try(ctx, func(ctx context.Context) {

		model := dao.PestRecord.Ctx(ctx).
			LeftJoin("field f", "f.id = pest_record.field_id").
			LeftJoin("crop c", "c.id = pest_record.crop_id").
			LeftJoin("sys_user u", "u.id = pest_record.solve_by"). // ✅ 新增
			Fields(`
		pest_record.*,
		f.name AS fieldName,
		c.name AS cropName,
		u.user_nickname AS solveName   -- ✅ 映射到结构体
	`)

		countModel := dao.PestRecord.Ctx(ctx)

		if req.FieldId != 0 {
			model = model.Where("pest_record.field_id=?", req.FieldId)
			countModel = countModel.Where("field_id=?", req.FieldId)
		}

		if req.CropId != 0 {
			model = model.Where("pest_record.crop_id=?", req.CropId)
			countModel = countModel.Where("crop_id=?", req.CropId)
		}
		if req.CreatedBy != 0 {
			model = model.Where("pest_record.created_by=?", req.CreatedBy)
			countModel = countModel.Where("created_by=?", req.CreatedBy)
		} else {
			model = model.Where("pest_record.status=?", req.Status)
			countModel = countModel.Where("status=?", req.Status)
		}

		resp.Total, err = countModel.Count()
		liberr.ErrIsNil(ctx, err, "获取虫灾数量失败")
		resp.CurrentPage = req.PageNum
		err = model.Page(resp.CurrentPage, req.PageSize).
			Order("pest_record.id desc").
			Scan(&resp.List)

		liberr.ErrIsNil(ctx, err, "查询虫灾列表失败")

	})

	return resp
}

/*
新增或修改虫灾
*/
func (s *sSysPest) Add(ctx context.Context, req *system.AddPestReq) *system.AddOrUpdatePestResp {

	resp := &system.AddOrUpdatePestResp{}

	g.Try(ctx, func(ctx context.Context) {

		db := dao.PestRecord.Ctx(ctx)
		_, err := db.Insert(req)
		liberr.ErrIsNil(ctx, err, "新增虫灾失败")

	})

	return resp
}

/*
删除虫灾
*/
func (s *sSysPest) Delete(ctx context.Context, req *system.DeletePestReq) *system.DeletePestResp {

	resp := &system.DeletePestResp{}

	g.Try(ctx, func(ctx context.Context) {

		_, err := dao.PestRecord.Ctx(ctx).
			Delete("id=?", req.Id)

		liberr.ErrIsNil(ctx, err, "删除虫灾失败")

	})

	return resp
}

/*
修改虫灾状态
*/
func (s *sSysPest) UpdateStatus(ctx context.Context, req *system.UpdatePestStatusReq) *system.UpdatePestStatusResp {

	resp := &system.UpdatePestStatusResp{}

	g.Try(ctx, func(ctx context.Context) {

		_, err := dao.PestRecord.Ctx(ctx).
			WherePri(req.Id).
			Data(g.Map{
				"status":       req.Status,
				"slove_answer": req.SloveAnswer,
				"solve_by":     req.SolveBY,
			}).
			Update()

		liberr.ErrIsNil(ctx, err, "修改虫灾状态失败")

	})

	return resp
}
