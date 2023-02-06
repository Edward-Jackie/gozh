package repository

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/jinzhu/copier"
	"gozh/internal/video/domain/entity"
	"gozh/internal/video/infrastructure/model"
)

type VideoInfoItemRepository struct {
	context *web.Context
}

func NewVideoInfoItemRepository(context *web.Context) *VideoInfoItemRepository {
	return &VideoInfoItemRepository{context: context}
}

// Count 统计数量
func (repository *VideoInfoItemRepository) Count(conditions *orm.Conditions) (int64, error) {
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.VideoInfoItem{}, conditions)
	err := query.Count(&total).Error
	return total, err
}

// Save 保存
func (repository *VideoInfoItemRepository) Save(videoInfoItem *entity.VideoInfoItem) (*entity.VideoInfoItem, error) {
	videoInfoItemModel := &model.VideoInfoItem{}
	err := copier.Copy(videoInfoItemModel, videoInfoItem)
	if err != nil {
		return videoInfoItem, err
	}
	if videoInfoItemModel.Id <= 0 {
		videoInfoItemModel.Id, err = tools.NewSnowflakeId()
		if err != nil {
			return videoInfoItem, err
		}
		videoInfoItem.Id = videoInfoItemModel.Id
		err = repository.context.Transaction().Context.Create(videoInfoItemModel).Error
	} else {
		err = repository.context.Transaction().Context.Model(videoInfoItemModel).Updates(videoInfoItemModel).Error
	}
	return videoInfoItem, err
}

// Find 获取列表
func (repository *VideoInfoItemRepository) Find(conditions *orm.Conditions) (int64, []*entity.VideoInfoItem, error) {
	list := make([]*entity.VideoInfoItem, 0)
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.VideoInfoItem{}, conditions)
	err := query.Count(&total).Error
	if err != nil {
		return total, list, err
	}
	err = orm.OffsetLimit(query, conditions).Find(&list).Error
	return total, list, err
}

// FindOne 获取单条记录
func (repository *VideoInfoItemRepository) FindOne(conditions *orm.Conditions) (*entity.VideoInfoItem, error) {
	videoInfoItem := &entity.VideoInfoItem{}
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.VideoInfoItem{}, conditions)
	err := query.First(videoInfoItem).Error
	return videoInfoItem, err
}

// FindById 根据ID查询
func (repository *VideoInfoItemRepository) FindById(id int64) (*entity.VideoInfoItem, error) {
	videoInfoItemModel := &model.VideoInfoItem{}
	videoInfoItem := &entity.VideoInfoItem{}
	err := orm.AdvanceSearch(repository.context.Transaction(), &model.VideoInfoItem{}, &orm.Conditions{
		Equal: map[string]interface{}{"id": id},
	}).First(videoInfoItemModel).Error
	if err != nil {
		return videoInfoItem, err
	}
	err = copier.Copy(videoInfoItem, videoInfoItemModel)
	return videoInfoItem, err
}

// Delete 删除
func (repository *VideoInfoItemRepository) Delete(videoInfoItem *entity.VideoInfoItem) (*entity.VideoInfoItem, error) {
	videoInfoItemModel := &model.VideoInfoItem{}
	_ = copier.Copy(videoInfoItemModel, videoInfoItem)
	err := repository.context.Transaction().Context.Delete(videoInfoItemModel).Error
	return videoInfoItem, err
}
