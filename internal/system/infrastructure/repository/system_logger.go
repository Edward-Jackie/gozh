package repository

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/jinzhu/copier"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/infrastructure/model"
)

type SystemLoggerRepository struct {
	context *web.Context
}

func NewSystemLoggerRepository(context *web.Context) *SystemLoggerRepository {
	return &SystemLoggerRepository{context: context}
}

// Count 统计数量
func (repository *SystemLoggerRepository) Count(conditions *orm.Conditions) (int64, error) {
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemLogger{}, conditions)
	err := query.Count(&total).Error
	return total, err
}

// Save 保存
func (repository *SystemLoggerRepository) Save(systemLogger *entity.SystemLogger) (*entity.SystemLogger, error) {
	systemLoggerModel := &model.SystemLogger{}
	err := copier.Copy(systemLoggerModel, systemLogger)
	if err != nil {
		return systemLogger, err
	}
	if systemLoggerModel.Id <= 0 {
		systemLoggerModel.Id, err = tools.NewSnowflakeId()
		if err != nil {
			return systemLogger, err
		}
		systemLogger.Id = systemLoggerModel.Id
		err = repository.context.Transaction().Context.Create(systemLoggerModel).Error
	} else {
		err = repository.context.Transaction().Context.Model(systemLoggerModel).Updates(systemLoggerModel).Error
	}
	return systemLogger, err
}

// Find 获取列表
func (repository *SystemLoggerRepository) Find(conditions *orm.Conditions) (int64, []*entity.SystemLogger, error) {
	list := make([]*entity.SystemLogger, 0)
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemLogger{}, conditions)
	err := query.Count(&total).Error
	if err != nil {
		return total, list, err
	}
	err = orm.OffsetLimit(query, conditions).Find(&list).Error
	return total, list, err
}

// FindOne 获取单条记录
func (repository *SystemLoggerRepository) FindOne(conditions *orm.Conditions) (*entity.SystemLogger, error) {
	systemLogger := &entity.SystemLogger{}
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemLogger{}, conditions)
	err := query.First(systemLogger).Error
	return systemLogger, err
}

// FindById 根据ID查询
func (repository *SystemLoggerRepository) FindById(id int64) (*entity.SystemLogger, error) {
	systemLoggerModel := &model.SystemLogger{}
	systemLogger := &entity.SystemLogger{}
	err := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemLogger{}, &orm.Conditions{
		Equal: map[string]interface{}{"id": id},
	}).First(systemLoggerModel).Error
	if err != nil {
		return systemLogger, err
	}
	err = copier.Copy(systemLogger, systemLoggerModel)
	return systemLogger, err
}

// Delete 删除
func (repository *SystemLoggerRepository) Delete(systemLogger *entity.SystemLogger) (*entity.SystemLogger, error) {
	systemLoggerModel := &model.SystemLogger{}
	_ = copier.Copy(systemLoggerModel, systemLogger)
	err := repository.context.Transaction().Context.Delete(systemLoggerModel).Error
	return systemLogger, err
}
