package query

// GetVideoInfo 获取视频
type GetVideoInfo struct {
	Id int64 `json:"id,string" form:"id" validate:"required"`
}
