package query

type GetItemPlayCount struct {
	VideoId int64 `json:"videoId" form:"videoId"`
}
