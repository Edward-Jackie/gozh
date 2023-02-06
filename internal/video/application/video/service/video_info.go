package service

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/gookit/goutil/structs"
	"github.com/jinzhu/copier"
	"gozh/internal/system"
	"gozh/internal/system/infrastructure/cache"
	"gozh/internal/video/application/video/command"
	"gozh/internal/video/application/video/dto"
	"gozh/internal/video/application/video/query"
	"gozh/internal/video/domain/entity"
	"gozh/internal/video/facade"
)

type VideoInfoService struct {
	context *web.Context
}

func NewVideoInfoService(context *web.Context) *VideoInfoService {
	return &VideoInfoService{context: context}
}

// Create 新增
func (service *VideoInfoService) Create(createVideoInfo *command.CreateVideoInfo) (interface{}, error) {
	if err := service.context.Validate(createVideoInfo); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	count, err := facade.VideoInfoRepository(service.context).Count(&orm.Conditions{Equal: map[string]interface{}{"name": createVideoInfo.Name}})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	if count > 0 {
		return nil, web.ThrowError(web.ArgError, createVideoInfo.Name+" 名称已存在")
	}
	videoInfo := &entity.VideoInfo{}
	_ = copier.Copy(videoInfo, createVideoInfo)
	videoInfo, err = facade.VideoInfoRepository(service.context).Save(videoInfo)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	//保存视频详情
	for _, item := range createVideoInfo.Video {
		_, err = facade.VideoInfoItemRepository(service.context).Save(&entity.VideoInfoItem{
			VideoId: videoInfo.Id,
			Etag:    item.Etag,
			Url:     item.Url,
			Poster:  item.Poster,
			Title:   item.Title,
		})
		if err != nil {
			return nil, web.ThrowError(web.InternalServerError, err.Error())
		}
	}
	_ = system.AddLogger(service.context, &system.Logger{
		UserId:     service.context.Auth().UserId,
		Comment:    "新增视频:" + videoInfo.Name,
		LoggerType: system.LoggerTypeVideo,
	})
	return videoInfo, nil
}

// Update 修改
func (service *VideoInfoService) Update(updateVideoInfo *command.UpdateVideoInfo) (interface{}, error) {
	if err := service.context.Validate(updateVideoInfo); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	videoInfoRepository := facade.VideoInfoRepository(service.context)
	videoInfo, err := videoInfoRepository.FindById(updateVideoInfo.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	videoInfo.Update(structs.ToMap(updateVideoInfo))
	videoInfo, err = facade.VideoInfoRepository(service.context).Save(videoInfo)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	//保存视频详情
	updateItemIds := make([]int64, 0)
	for _, item := range updateVideoInfo.Video {
		videoInfoItem, err := facade.VideoInfoItemRepository(service.context).Save(&entity.VideoInfoItem{
			Id:      item.Id,
			VideoId: videoInfo.Id,
			Etag:    item.Etag,
			Url:     item.Url,
			Poster:  item.Poster,
			Title:   item.Title,
		})
		if err != nil {
			return nil, web.ThrowError(web.InternalServerError, err.Error())
		}
		updateItemIds = append(updateItemIds, videoInfoItem.Id)
	}
	err = facade.VideoInfoItemDao(service.context).BatchDelete(updateItemIds, videoInfo.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	_ = system.AddLogger(service.context, &system.Logger{
		UserId:     service.context.Auth().UserId,
		Comment:    "编辑视频:" + videoInfo.Name,
		LoggerType: system.LoggerTypeVideo,
	})
	return videoInfo, nil
}

// List 列表
func (service *VideoInfoService) List(listVideoInfo *query.ListVideoInfo) (interface{}, error) {
	condition := &orm.Conditions{
		Like: map[string]interface{}{
			"name": listVideoInfo.Name,
		},
		Pagination: &orm.Pagination{
			Page:     listVideoInfo.Page,
			PageSize: listVideoInfo.PageSize,
		},
		OrderColumn: "created_at",
		OrderBy:     "desc",
	}
	if listVideoInfo.CategoryId > 0 {
		condition.Equal = map[string]interface{}{
			"categoryId": listVideoInfo.CategoryId,
		}
	}
	total, list, err := facade.VideoInfoRepository(service.context).Find(condition)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	//获取所有类目
	_, categories, _ := facade.VideoCategoryRepository(service.context).Find(&orm.Conditions{})
	categoryData := make(map[int64]*entity.VideoCategory)
	for _, item := range categories {
		categoryData[item.Id] = item
	}
	result := make([]*dto.VideoInfo, 0)
	videoIds := make([]int64, 0)
	for _, item := range list {
		videoInfo := dto.TransferVideoInfo(item)
		user, err := cache.GetUserById(service.context, item.CreatedBy)
		if err != nil {
			return nil, web.ThrowError(web.InternalServerError, err.Error())
		}
		videoInfo.CreatedName = user.NickName
		if cate, ok := categoryData[videoInfo.CategoryId]; ok {
			videoInfo.Category = cate
			videoInfo.CategoryPath = dto.TransferCategoryPath(categoryData, cate)
		}
		//视频详情
		_, videoItems, err := facade.VideoInfoItemRepository(service.context).Find(&orm.Conditions{
			Equal: map[string]interface{}{
				"videoId": videoInfo.Id,
			},
			OrderBy:     "desc",
			OrderColumn: "created_at",
		})
		if err != nil {
			return nil, web.ThrowError(web.InternalServerError, err.Error())
		}
		videoInfo.Video = videoItems
		result = append(result, videoInfo)
		videoIds = append(videoIds, videoInfo.Id)
	}
	if len(videoIds) > 0 {
		playCounts, _ := facade.VideoInfoDao(service.context).GetPlayCount(videoIds)
		if len(playCounts) > 0 {
			for _, item := range result {
				for _, play := range playCounts {
					if play.Id == item.Id {
						item.PlayCount = play.Count
					}
				}
			}
		}
	}
	return web.NewPagination(total, result), nil
}

// Get 获取信息
func (service *VideoInfoService) Get(getVideoInfo *query.GetVideoInfo) (interface{}, error) {
	if err := service.context.Validate(getVideoInfo); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	videoInfo, err := facade.VideoInfoRepository(service.context).FindOne(&orm.Conditions{
		Equal: map[string]interface{}{"id": getVideoInfo.Id},
	})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	//获取视频详情
	_, videoItems, err := facade.VideoInfoItemRepository(service.context).Find(&orm.Conditions{
		Equal: map[string]interface{}{
			"videoId": videoInfo.Id,
		},
		OrderBy:     "desc",
		OrderColumn: "created_at",
	})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	result := dto.TransferVideoInfo(videoInfo)
	result.Video = videoItems
	return result, nil
}

// Delete 删除
func (service *VideoInfoService) Delete(deleteVideoInfo *command.DeleteVideoInfo) (interface{}, error) {
	if err := service.context.Validate(deleteVideoInfo); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	videoInfo, err := facade.VideoInfoRepository(service.context).FindById(deleteVideoInfo.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	videoInfo, err = facade.VideoInfoRepository(service.context).Delete(videoInfo)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	_ = system.AddLogger(service.context, &system.Logger{
		UserId:     service.context.Auth().UserId,
		Comment:    "删除视频:" + videoInfo.Name,
		LoggerType: system.LoggerTypeVideo,
	})
	return videoInfo, nil
}

// GetItemPlayCount 获取视频明细播放统计
func (service *VideoInfoService) GetItemPlayCount(getItemPlayCount *query.GetItemPlayCount) (interface{}, error) {
	if err := service.context.Validate(getItemPlayCount); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	return facade.VideoInfoDao(service.context).GetItemPlayCount(getItemPlayCount.VideoId)
}
