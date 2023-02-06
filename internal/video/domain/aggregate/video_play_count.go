package aggregate

type VideoPlayCount struct {
	Id    int64 `json:"id,string"`
	Count int64 `json:"count"`
}
