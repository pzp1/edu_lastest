package model

type PestListResp struct {
	Id          int64  `json:"id"`          // 主键ID
	FieldId     int64  `json:"fieldId"`     // 地块ID
	FieldName   string `json:"fieldName"`   // 地块名称
	CropId      int64  `json:"cropId"`      // 作物ID
	CropName    string `json:"cropName"`    // 作物名称
	Description string `json:"description"` // 虫灾描述
	Status      int    `json:"status"`      // 0未解决 1已解决
	CreatedAt   string `json:"createdAt"`   // 上报时间
	SloveAnswer string `json:"sloveAnswer"`
	SolveName   string `json:"solveName"`
}
