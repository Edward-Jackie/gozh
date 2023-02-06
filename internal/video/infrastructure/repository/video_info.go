package repository

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/jinzhu/copier"
	"gozh/internal/video/domain/entity"
	"gozh/internal/video/infrastructure/model"
	"time"
)

type VideoInfoRepository struct {
	context *web.Context
}

func NewVideoInfoRepository(context *web.Context) *VideoInfoRepository {
	return &VideoInfoRepository{context: context}
}

// Count 统计数量
func (repository *VideoInfoRepository) Count(conditions *orm.Conditions) (int64, error) {
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.VideoInfo{}, conditions)
	err := query.Count(&total).Error
	return total, err
}

// Save 保存
func (repository *VideoInfoRepository) Save(videoInfo *entity.VideoInfo) (*entity.VideoInfo, error) {
	videoInfoModel := &model.VideoInfo{}
	videoInfo.CreatedBy = repository.context.Auth().UserId
	mTime := time.Now()
	err := copier.Copy(videoInfoModel, videoInfo)
	if err != nil {
		return videoInfo, err
	}
	if videoInfoModel.Id <= 0 {
		videoInfoModel.Id, err = tools.NewSnowflakeId()
		if err != nil {
			return videoInfo, err
		}
		videoInfo.CreatedAt = mTime
		videoInfo.Id = videoInfoModel.Id
		videoInfoModel.UpdatedAt = &mTime
		err = repository.context.Transaction().Context.Create(videoInfoModel).Error
	} else {
		videoInfoModel.UpdatedAt = &mTime
		err = repository.context.Transaction().Context.Model(videoInfoModel).Updates(videoInfoModel).Error
	}
	return videoInfo, err
}

// Find 获取列表
func (repository *VideoInfoRepository) Find(conditions *orm.Conditions) (int64, []*entity.VideoInfo, error) {
	list := make([]*entity.VideoInfo, 0)
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.VideoInfo{}, conditions)
	err := query.Count(&total).Error
	if err != nil {
		return total, list, err
	}
	err = orm.OffsetLimit(query, conditions).Find(&list).Error
	return total, list, err
}

// FindOne 获取单条记录
func (repository *VideoInfoRepository) FindOne(conditions *orm.Conditions) (*entity.VideoInfo, error) {
	videoInfo := &entity.VideoInfo{}
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.VideoInfo{}, conditions)
	err := query.First(videoInfo).Error
	return videoInfo, err
}

// FindById 根据ID查询
func (repository *VideoInfoRepository) FindById(id int64) (*entity.VideoInfo, error) {
	videoInfoModel := &model.VideoInfo{}
	videoInfo := &entity.VideoInfo{}
	err := orm.AdvanceSearch(repository.context.Transaction(), &model.VideoInfo{}, &orm.Conditions{
		Equal: map[string]interface{}{"id": id},
	}).First(videoInfoModel).Error
	if err != nil {
		return videoInfo, err
	}
	err = copier.Copy(videoInfo, videoInfoModel)
	return videoInfo, err
}

// Delete 删除
func (repository *VideoInfoRepository) Delete(videoInfo *entity.VideoInfo) (*entity.VideoInfo, error) {
	videoInfoModel := &model.VideoInfo{}
	_ = copier.Copy(videoInfoModel, videoInfo)
	err := repository.context.Transaction().Context.Delete(videoInfoModel).Error
	return videoInfo, err
}
