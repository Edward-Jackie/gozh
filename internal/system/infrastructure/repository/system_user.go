package repository

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/jinzhu/copier"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/infrastructure/model"
)

type SystemUserRepository struct {
	context *web.Context
}

func NewSystemUserRepository(context *web.Context) *SystemUserRepository {
	return &SystemUserRepository{context: context}
}

// Count 统计数量
func (repository *SystemUserRepository) Count(conditions *orm.Conditions) (int64, error) {
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemUser{}, conditions)
	err := query.Count(&total).Error
	return total, err
}

// Save 保存
func (repository *SystemUserRepository) Save(systemUser *entity.SystemUser) (*entity.SystemUser, error) {
	systemUserModel := &model.SystemUser{}
	err := copier.Copy(systemUserModel, systemUser)
	if err != nil {
		return systemUser, err
	}
	if systemUserModel.Id <= 0 {
		systemUserModel.Id, err = tools.NewSnowflakeId()
		if err != nil {
			return systemUser, err
		}
		systemUser.Id = systemUserModel.Id
		auth := repository.context.Auth()
		if auth != nil {
			systemUserModel.CreatedBy = auth.UserId
		}
		err = repository.context.Transaction().Context.Create(systemUserModel).Error
	} else {
		auth := repository.context.Auth()
		if auth != nil {
			systemUserModel.UpdatedBy = auth.UserId
		}
		err = repository.context.Transaction().Context.Model(systemUserModel).Updates(systemUserModel).Error
	}
	return systemUser, err
}

// Find 获取列表
func (repository *SystemUserRepository) Find(conditions *orm.Conditions) (int64, []*entity.SystemUser, error) {
	list := make([]*model.SystemUser, 0)
	result := make([]*entity.SystemUser, 0)
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemUser{}, conditions)
	err := query.Count(&total).Error
	if err != nil {
		return total, result, err
	}
	err = orm.OffsetLimit(query, conditions).Find(&list).Error
	if err != nil {
		return total, result, err
	}
	_ = copier.Copy(&result, &list)
	return total, result, err
}

// FindOne 获取单条记录
func (repository *SystemUserRepository) FindOne(conditions *orm.Conditions) (*entity.SystemUser, error) {
	systemUser := &entity.SystemUser{}
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemUser{}, conditions)
	err := query.First(systemUser).Error
	return systemUser, err
}

// FindById 根据ID查询
func (repository *SystemUserRepository) FindById(id int64) (*entity.SystemUser, error) {
	systemUserModel := &model.SystemUser{}
	systemUser := &entity.SystemUser{}
	err := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemUser{}, &orm.Conditions{
		Equal: map[string]interface{}{"id": id},
	}).First(systemUserModel).Error
	if err != nil {
		return systemUser, err
	}
	err = copier.Copy(systemUser, systemUserModel)
	return systemUser, err
}

// Delete 删除
func (repository *SystemUserRepository) Delete(systemUser *entity.SystemUser) (*entity.SystemUser, error) {
	systemUserModel := &model.SystemUser{}
	_ = copier.Copy(systemUserModel, systemUser)
	err := repository.context.Transaction().Context.Delete(systemUserModel).Error
	return systemUser, err
}
