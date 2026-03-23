package model

import "github.com/gogf/gf/v2/os/gtime"

type MaterialListResp struct {
	Id              int64       `json:"id"`
	Name            string      `json:"name"`
	Type            int         `json:"type"`
	Category        int         `json:"category"`
	UsageMethod     string      `json:"usageMethod"`
	NutrientContent string      `json:"nutrientContent"`
	CreatedAt       *gtime.Time `json:"createdAt"`
}
