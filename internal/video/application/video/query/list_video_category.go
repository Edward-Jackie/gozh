package query

type ListVideoCategory struct {
	Page     int `json:"page"`     // 页码
	PageSize int `json:"pageSize"` // 每页行数
}
