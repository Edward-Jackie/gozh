package repository

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/jinzhu/copier"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/infrastructure/model"
)

type SystemEventRepository struct {
	context *web.Context
}

func NewSystemEventRepository(context *web.Context) *SystemEventRepository {
	return &SystemEventRepository{context: context}
}

// Count 统计数量
func (repository *SystemEventRepository) Count(conditions *orm.Conditions) (int64, error) {
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemEvent{}, conditions)
	err := query.Count(&total).Error
	return total, err
}

// Save 保存
func (repository *SystemEventRepository) Save(systemEvent *entity.SystemEvent) (*entity.SystemEvent, error) {
	systemEventModel := &model.SystemEvent{}
	err := copier.Copy(systemEventModel, systemEvent)
	if err != nil {
		return systemEvent, err
	}
	if systemEventModel.Id <= 0 {
		systemEventModel.Id, err = tools.NewSnowflakeId()
		if err != nil {
			return systemEvent, err
		}
		systemEvent.Id = systemEventModel.Id
		err = repository.context.Transaction().Context.Create(systemEventModel).Error
	} else {
		err = repository.context.Transaction().Context.Model(systemEventModel).Updates(systemEventModel).Error
	}
	return systemEvent, err
}

// Find 获取列表
func (repository *SystemEventRepository) Find(conditions *orm.Conditions) (int64, []*entity.SystemEvent, error) {
	list := make([]*entity.SystemEvent, 0)
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemEvent{}, conditions)
	err := query.Count(&total).Error
	if err != nil {
		return total, list, err
	}
	err = orm.OffsetLimit(query, conditions).Find(&list).Error
	return total, list, err
}

// FindOne 获取单条记录
func (repository *SystemEventRepository) FindOne(conditions *orm.Conditions) (*entity.SystemEvent, error) {
	systemEvent := &entity.SystemEvent{}
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemEvent{}, conditions)
	err := query.First(systemEvent).Error
	return systemEvent, err
}

// FindById 根据ID查询
func (repository *SystemEventRepository) FindById(id int64) (*entity.SystemEvent, error) {
	systemEventModel := &model.SystemEvent{}
	systemEvent := &entity.SystemEvent{}
	err := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemEvent{}, &orm.Conditions{
		Equal: map[string]interface{}{"id": id},
	}).First(systemEventModel).Error
	if err != nil {
		return systemEvent, err
	}
	err = copier.Copy(systemEvent, systemEventModel)
	return systemEvent, err
}

// Delete 删除
func (repository *SystemEventRepository) Delete(systemEvent *entity.SystemEvent) (*entity.SystemEvent, error) {
	systemEventModel := &model.SystemEvent{}
	_ = copier.Copy(systemEventModel, systemEvent)
	err := repository.context.Transaction().Context.Delete(systemEventModel).Error
	return systemEvent, err
}
