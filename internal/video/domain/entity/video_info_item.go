package entity

// VideoInfoItem 视频详情
type VideoInfoItem struct {
	// ID
	Id int64 `json:"id,string"`
	// 视频信息ID
	VideoId int64 `json:"videoId,string"`
	// 视频标题
	Title string `json:"title"`
	// etag
	Etag string `json:"etag"`
	// 视频地址
	Url string `json:"url"`
	// 缩略图
	Poster string `json:"poster"`
}

// Update 更新数据
func (videoInfoItem *VideoInfoItem) Update(options map[string]interface{}) {
	// ID
	if id, ok := options["Id"]; ok {
		videoInfoItem.Id = id.(int64)
	}
	// 视频信息ID
	if videoId, ok := options["VideoId"]; ok {
		videoInfoItem.VideoId = videoId.(int64)
	}
	// 视频标题
	if title, ok := options["Title"]; ok {
		videoInfoItem.Title = title.(string)
	}
	// etag
	if etag, ok := options["Etag"]; ok {
		videoInfoItem.Etag = etag.(string)
	}
	// 视频地址
	if url, ok := options["Url"]; ok {
		videoInfoItem.Url = url.(string)
	}
	// 缩略图
	if poster, ok := options["Poster"]; ok {
		videoInfoItem.Poster = poster.(string)
	}
}
