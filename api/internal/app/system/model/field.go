package model

import "github.com/gogf/gf/v2/util/gmeta"

type FieldListResp struct {
	gmeta.Meta `orm:"table:field"`
	Id         int64      `json:"id" orm:"id"`
	UserId     int64      `json:"userId" orm:"user_id"`       // 所属种植户
	Name       string     `json:"name" orm:"name"`            // 地块名称
	Area       float64    `json:"area" orm:"area"`            // 面积（亩）
	Location   string     `json:"location" orm:"location"`    // 位置描述
	Status     int        `json:"status" orm:"status"`        // 状态 1使用中 2停用
	SortOrder  int        `json:"sortOrder" orm:"sort_order"` // 排序
	Deleted    int        `json:"deleted" orm:"deleted"`      // 是否删除 0否 1是
	CreatedAt  string     `json:"createdAt" orm:"created_at"` // 创建时间
	UpdatedAt  string     `json:"updatedAt" orm:"updated_at"` // 更新时间
	UserInfo   *FieldUser `orm:"with:id=user_id" json:"userInfo"`
}

type FieldUser struct {
	gmeta.Meta   `orm:"table:sys_user"`
	Id           int64  `json:"id"`
	UserNickName string `json:"userNickName" orm:"user_nickname"`
}
