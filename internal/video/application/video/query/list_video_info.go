package query

type ListVideoInfo struct {
	Name       string `json:"name"`              // 视频名称
	CategoryId int64  `json:"categoryId,string"` // 视频类目
	Page       int    `json:"page"`              // 页码
	PageSize   int    `json:"pageSize"`          // 每页行数
}
