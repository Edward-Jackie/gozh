package service

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/jinzhu/copier"
	"gozh/internal/video/application/video/command"
	"gozh/internal/video/application/video/dto"
	"gozh/internal/video/application/video/query"
	"gozh/internal/video/domain/entity"
	"gozh/internal/video/facade"
)

type VideoCategoryService struct {
	context *web.Context
}

func NewVideoCategoryService(context *web.Context) *VideoCategoryService {
	return &VideoCategoryService{context: context}
}

// Create 新增
func (service *VideoCategoryService) Create(createVideoCategory *command.CreateVideoCategory) (interface{}, error) {
	if err := service.context.Validate(createVideoCategory); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	//验证名称是否存在
	count, err := facade.VideoCategoryRepository(service.context).Count(&orm.Conditions{
		Equal: map[string]interface{}{
			"name": createVideoCategory.Name,
		},
	})
	if err != nil || count > 0 {
		return nil, web.ThrowError(web.ArgError, "类目名称已存在")
	}
	videoCategory := &entity.VideoCategory{}
	_ = copier.Copy(videoCategory, createVideoCategory)
	if createVideoCategory.ParentId > 0 {
		parent, err := facade.VideoCategoryRepository(service.context).FindById(createVideoCategory.ParentId)
		if err != nil {
			return nil, web.ThrowError(web.ArgError, "父级类目不存在")
		}
		videoCategory.Level = parent.Level + 1
		videoCategory.Path = videoCategory.CreatePath(parent.Path, parent.Id)
	} else {
		videoCategory.Level = 1
	}
	videoCategory, err = facade.VideoCategoryRepository(service.context).Save(videoCategory)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return videoCategory, nil
}

// Update 修改
func (service *VideoCategoryService) Update(updateVideoCategory *command.UpdateVideoCategory) (interface{}, error) {
	if err := service.context.Validate(updateVideoCategory); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	//验证名称是否存在
	count, err := facade.VideoCategoryRepository(service.context).Count(&orm.Conditions{
		Equal: map[string]interface{}{
			"name": updateVideoCategory.Name,
		},
		NotEqual: map[string]interface{}{
			"id": updateVideoCategory.Id,
		},
	})
	if err != nil || count > 0 {
		return nil, web.ThrowError(web.ArgError, "类目名称已存在")
	}
	videoCategoryRepository := facade.VideoCategoryRepository(service.context)
	videoCategory, err := videoCategoryRepository.FindById(updateVideoCategory.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	videoCategory.Name = updateVideoCategory.Name
	videoCategory.ParentId = updateVideoCategory.ParentId
	if updateVideoCategory.ParentId > 0 {
		parent, err := videoCategoryRepository.FindById(updateVideoCategory.ParentId)
		if err != nil || parent.Id < 0 {
			return nil, web.ThrowError(web.ArgError, "父级类目不存在")
		}
		videoCategory.Level = parent.Level + 1
		videoCategory.Path = videoCategory.CreatePath(parent.Path, parent.Id)
	} else {
		videoCategory.Level = 1
		videoCategory.Path = ""
	}
	videoCategory, err = facade.VideoCategoryRepository(service.context).Save(videoCategory)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return videoCategory, nil
}

// List 列表
func (service *VideoCategoryService) List(listVideoCategory *query.ListVideoCategory) (interface{}, error) {
	total, list, err := facade.VideoCategoryRepository(service.context).Find(&orm.Conditions{})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return web.NewPagination(total, dto.TreeCategory(list, 0)), nil
}

// Get 获取信息
func (service *VideoCategoryService) Get(getVideoCategory *query.GetVideoCategory) (interface{}, error) {
	if err := service.context.Validate(getVideoCategory); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	videoCategory, err := facade.VideoCategoryRepository(service.context).FindOne(&orm.Conditions{
		Equal: map[string]interface{}{"id": getVideoCategory.Id},
	})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return videoCategory, nil
}

// Delete 删除
func (service *VideoCategoryService) Delete(deleteVideoCategory *command.DeleteVideoCategory) (interface{}, error) {
	if err := service.context.Validate(deleteVideoCategory); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	//是否存在子集
	count, err := facade.VideoCategoryRepository(service.context).Count(&orm.Conditions{
		Equal: map[string]interface{}{
			"parentId": deleteVideoCategory.Id,
		},
	})
	if err != nil || count > 0 {
		return nil, web.ThrowError(web.ArgError, "请先删除子集数据")
	}
	videoCategory, err := facade.VideoCategoryRepository(service.context).FindById(deleteVideoCategory.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	videoCategory, err = facade.VideoCategoryRepository(service.context).Delete(videoCategory)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return videoCategory, nil
}
