package query

type ListSettingDictValues struct {
	DictId   int64  `json:"dictId,string"` // 字典ID
	Code     string `json:"code"`          // 编码
	Page     int    `json:"page"`          // 页码
	PageSize int    `json:"pageSize"`      // 每页行数
}
