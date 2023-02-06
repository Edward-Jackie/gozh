package repository

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/jinzhu/copier"
	"gozh/internal/video/domain/entity"
	"gozh/internal/video/infrastructure/model"
)

type VideoCategoryRepository struct {
	context *web.Context
}

func NewVideoCategoryRepository(context *web.Context) *VideoCategoryRepository {
	return &VideoCategoryRepository{context: context}
}

// Count 统计数量
func (repository *VideoCategoryRepository) Count(conditions *orm.Conditions) (int64, error) {
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.VideoCategory{}, conditions)
	err := query.Count(&total).Error
	return total, err
}

// Save 保存
func (repository *VideoCategoryRepository) Save(videoCategory *entity.VideoCategory) (*entity.VideoCategory, error) {
	videoCategoryModel := &model.VideoCategory{}
	err := copier.Copy(videoCategoryModel, videoCategory)
	if err != nil {
		return videoCategory, err
	}
	if videoCategoryModel.Id <= 0 {
		videoCategoryModel.Id, err = tools.NewSnowflakeId()
		if err != nil {
			return videoCategory, err
		}
		videoCategory.Id = videoCategoryModel.Id
		err = repository.context.Transaction().Context.Create(videoCategoryModel).Error
	} else {
		err = repository.context.Transaction().Context.Model(videoCategoryModel).Updates(videoCategoryModel).Error
	}
	return videoCategory, err
}

// Find 获取列表
func (repository *VideoCategoryRepository) Find(conditions *orm.Conditions) (int64, []*entity.VideoCategory, error) {
	list := make([]*entity.VideoCategory, 0)
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.VideoCategory{}, conditions)
	err := query.Count(&total).Error
	if err != nil {
		return total, list, err
	}
	err = orm.OffsetLimit(query, conditions).Find(&list).Error
	return total, list, err
}

// FindOne 获取单条记录
func (repository *VideoCategoryRepository) FindOne(conditions *orm.Conditions) (*entity.VideoCategory, error) {
	videoCategory := &entity.VideoCategory{}
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.VideoCategory{}, conditions)
	err := query.First(videoCategory).Error
	return videoCategory, err
}

// FindById 根据ID查询
func (repository *VideoCategoryRepository) FindById(id int64) (*entity.VideoCategory, error) {
	videoCategoryModel := &model.VideoCategory{}
	videoCategory := &entity.VideoCategory{}
	err := orm.AdvanceSearch(repository.context.Transaction(), &model.VideoCategory{}, &orm.Conditions{
		Equal: map[string]interface{}{"id": id},
	}).First(videoCategoryModel).Error
	if err != nil {
		return videoCategory, err
	}
	err = copier.Copy(videoCategory, videoCategoryModel)
	return videoCategory, err
}

// Delete 删除
func (repository *VideoCategoryRepository) Delete(videoCategory *entity.VideoCategory) (*entity.VideoCategory, error) {
	videoCategoryModel := &model.VideoCategory{}
	_ = copier.Copy(videoCategoryModel, videoCategory)
	err := repository.context.Transaction().Context.Delete(videoCategoryModel).Error
	return videoCategory, err
}
