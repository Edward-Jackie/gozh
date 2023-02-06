package query

type ListSystemEventPlay struct {
	VideoId     int64    `json:"videoId,string"`
	VideoItemId int64    `json:"videoItemId,string"`
	PlayTime    []string `json:"playTime"`
	Ip          string   `json:"ip"`
	Page        int      `json:"page"`     // 页码
	PageSize    int      `json:"pageSize"` // 每页行数
}
