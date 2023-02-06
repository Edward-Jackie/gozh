package dto

import (
	"github.com/gookit/goutil/strutil"
	"gozh/internal/video/domain/entity"
	"strings"
	"time"
)

type VideoInfo struct {
	// ID
	Id int64 `json:"id,string"`
	// 视频名称
	Name string `json:"name"`
	// 类目ID
	CategoryId int64 `json:"categoryId,string"`
	// 类目路径名称
	CategoryPath string `json:"categoryPath"`
	// 视频地址，存储为数组json
	Video []*entity.VideoInfoItem `json:"video"`
	// 备注
	Comment string `json:"comment"`
	// 播放总量
	PlayCount int64 `json:"playCount"`
	// 创建时间
	CreatedAt time.Time `json:"createdAt"`
	// 创建人ID
	CreatedBy int64 `json:"createdBy,string"`
	// 创建人名字
	CreatedName string `json:"createdName,omitempty"`
	// 类目
	Category *entity.VideoCategory `json:"category"`
}

func TransferVideoInfo(videInfo *entity.VideoInfo) *VideoInfo {
	return &VideoInfo{
		Id:          videInfo.Id,
		Name:        videInfo.Name,
		CategoryId:  videInfo.CategoryId,
		Comment:     videInfo.Comment,
		CreatedAt:   videInfo.CreatedAt,
		CreatedBy:   videInfo.CreatedBy,
		CreatedName: "",
	}
}

func TransferCategoryPath(data map[int64]*entity.VideoCategory, category *entity.VideoCategory) string {
	if category.ParentId > 0 {
		names := make([]string, 0)
		pathArr := strings.Split(category.Path, ",")
		for _, id := range pathArr {
			if c, ok := data[strutil.Int64(id)]; ok {
				names = append(names, c.Name)
			}
		}
		if c, ok := data[category.Id]; ok {
			names = append(names, c.Name)
		}
		return strings.Join(names, "-")
	} else {
		return category.Name
	}
}
