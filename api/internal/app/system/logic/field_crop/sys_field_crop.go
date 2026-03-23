package field_crop

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	"github.com/tiger1103/gfast/v3/internal/app/system/dao"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() { service.RegisterSysFieldCrop(NewSysFieldCrop()) }

type sSysFieldCrop struct{}

func NewSysFieldCrop() *sSysFieldCrop {
	return &sSysFieldCrop{}
}

/*
种植记录列表
*/
func (s *sSysFieldCrop) List(ctx context.Context, req *system.FieldCropListReq) *system.FieldCropListResp {

	resp := &system.FieldCropListResp{}
	var err error

	g.Try(ctx, func(ctx context.Context) {

		model := dao.FieldCrop.Ctx(ctx).
			LeftJoin("field f", "f.id = field_crop.field_id").
			Fields("field_crop.*")

		if req.FieldId != 0 {
			model = model.Where("field_crop.field_id = ?", req.FieldId)
		}

		if req.CropId != 0 {
			model = model.Where("field_crop.crop_id = ?", req.CropId)
		}

		if req.Status != 0 {
			model = model.Where("field_crop.status = ?", req.Status)
		}
		if req.ExpectedFlag == 1 {
			model = model.Where(`
		field_crop.status = 1
		AND field_crop.expected_harvest_date BETWEEN CURDATE() AND DATE_ADD(CURDATE(), INTERVAL 30 DAY)
	`)
		}
		// 查询用户
		userData := entity.SysUser{}
		err = dao.SysUser.Ctx(ctx).
			Where("id = ?", req.UserID).
			Scan(&userData)
		liberr.ErrIsNil(ctx, err, "获取用户数据失败")

		// 非管理员限制
		if userData.IsAdmin != 1 && req.UserID != 0 {
			model = model.Where("f.user_id = ?", req.UserID)
		}

		// ========================
		// 原生SQL统计数量
		// ========================

		sql := `
		SELECT COUNT(1)
		FROM field_crop fc
		LEFT JOIN field f ON f.id = fc.field_id
		WHERE 1=1
		`
		if req.ExpectedFlag == 1 {
			sql += `
	AND fc.expected_harvest_date BETWEEN CURDATE() AND DATE_ADD(CURDATE(), INTERVAL 30 DAY)
	`
		}
		args := []interface{}{}

		if req.FieldId != 0 {
			sql += " AND fc.field_id = ?"
			args = append(args, req.FieldId)
		}

		if req.CropId != 0 {
			sql += " AND fc.crop_id = ?"
			args = append(args, req.CropId)
		}

		if req.Status != 0 {
			sql += " AND fc.status = ?"
			args = append(args, req.Status)
		}

		if userData.IsAdmin != 1 && req.UserID != 0 {
			sql += " AND f.user_id = ?"
			args = append(args, req.UserID)
		}

		value, err := g.DB().GetValue(ctx, sql, args...)
		liberr.ErrIsNil(ctx, err, "获取种植记录数量失败")

		resp.Total = value.Int()

		// ========================
		// 分页参数
		// ========================

		if req.PageNum == 0 {
			req.PageNum = 1
		}

		resp.CurrentPage = req.PageNum

		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}

		// ========================
		// 查询列表
		// ========================

		err = model.
			Page(resp.CurrentPage, req.PageSize).
			Order("field_crop.id desc").
			Scan(&resp.List)

		liberr.ErrIsNil(ctx, err, "查询种植记录失败")

	})

	return resp
}

/*
新增或修改种植记录
*/
func (s *sSysFieldCrop) AddOrUpdate(ctx context.Context, req *system.AddOrUpdateFieldCropReq) *system.AddOrUpdateFieldCropResp {

	resp := &system.AddOrUpdateFieldCropResp{}

	g.Try(ctx, func(ctx context.Context) {

		db := dao.FieldCrop.Ctx(ctx)

		// 查询地块名称
		if req.FieldId != 0 {

			field := entity.Field{}

			err := dao.Field.Ctx(ctx).
				Where("id = ?", req.FieldId).
				Scan(&field)

			liberr.ErrIsNil(ctx, err, "查询地块失败")

			req.FieldName = field.Name
		}

		// 查询作物名称
		if req.CropId != 0 {

			crop := entity.Crop{}

			err := dao.Crop.Ctx(ctx).
				Where("id = ?", req.CropId).
				Scan(&crop)

			liberr.ErrIsNil(ctx, err, "查询作物失败")

			req.CropName = crop.Name
		}

		// 修改
		if req.Id != 0 {

			_, err := db.
				WherePri(req.Id).
				Update(req)

			liberr.ErrIsNil(ctx, err, "修改种植记录失败")

		} else {

			// 新增
			_, err := db.Insert(req)

			liberr.ErrIsNil(ctx, err, "新增种植记录失败")

		}

	})

	return resp
}

/*
删除种植记录
*/
func (s *sSysFieldCrop) Delete(ctx context.Context, req *system.DeleteFieldCropReq) *system.DeleteFieldCropResp {

	resp := &system.DeleteFieldCropResp{}

	g.Try(ctx, func(ctx context.Context) {

		_, err := dao.FieldCrop.Ctx(ctx).Delete("id=?", req.Id)

		liberr.ErrIsNil(ctx, err, "删除种植记录失败")

	})

	return resp
}

/*
种植记录详情
*/
func (s *sSysFieldCrop) Detail(ctx context.Context, req *system.FieldCropDetailReq) *system.FieldCropDetailResp {

	resp := &system.FieldCropDetailResp{}
	var err error

	g.Try(ctx, func(ctx context.Context) {

		err = dao.FieldCrop.Ctx(ctx).
			Where("id=?", req.Id).
			Scan(&resp)

		liberr.ErrIsNil(ctx, err, "查询种植记录详情失败")

	})

	return resp
}
