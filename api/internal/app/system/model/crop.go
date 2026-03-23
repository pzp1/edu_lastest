package model

type CropListResp struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	GrowthCycleDays int    `json:"growthCycleDays"`
	Category        int    `json:"category"`
	CreatedAt       string `json:"createdAt"`
}
