package repository

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/jinzhu/copier"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/infrastructure/model"
)

type SystemRoleRepository struct {
	context *web.Context
}

func NewSystemRoleRepository(context *web.Context) *SystemRoleRepository {
	return &SystemRoleRepository{context: context}
}

// Count 统计数量
func (repository *SystemRoleRepository) Count(conditions *orm.Conditions) (int64, error) {
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemRole{}, conditions)
	err := query.Count(&total).Error
	return total, err
}

// Save 保存
func (repository *SystemRoleRepository) Save(systemRole *entity.SystemRole) (*entity.SystemRole, error) {
	systemRoleModel := &model.SystemRole{}
	err := copier.Copy(systemRoleModel, systemRole)
	if err != nil {
		return systemRole, err
	}
	if systemRoleModel.Id <= 0 {
		systemRoleModel.Id, err = tools.NewSnowflakeId()
		if err != nil {
			return systemRole, err
		}
		systemRole.Id = systemRoleModel.Id
		systemRoleModel.CreatedBy = repository.context.Auth().UserId
		err = repository.context.Transaction().Context.Create(systemRoleModel).Error
	} else {
		//systemRoleModel.UpdatedBy = repository.context.Auth().UserId
		err = repository.context.Transaction().Context.Model(systemRoleModel).Updates(systemRoleModel).Error
	}
	return systemRole, err
}

// Find 获取列表
func (repository *SystemRoleRepository) Find(conditions *orm.Conditions) (int64, []*entity.SystemRole, error) {
	list := make([]*model.SystemRole, 0)
	result := make([]*entity.SystemRole, 0)
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemRole{}, conditions)
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
func (repository *SystemRoleRepository) FindOne(conditions *orm.Conditions) (*entity.SystemRole, error) {
	systemRole := &entity.SystemRole{}
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemRole{}, conditions)
	err := query.First(systemRole).Error
	return systemRole, err
}

// FindById 根据ID查询
func (repository *SystemRoleRepository) FindById(id int64) (*entity.SystemRole, error) {
	systemRoleModel := &model.SystemRole{}
	systemRole := &entity.SystemRole{}
	err := orm.AdvanceSearch(repository.context.Transaction(), &model.SystemRole{}, &orm.Conditions{
		Equal: map[string]interface{}{"id": id},
	}).First(systemRoleModel).Error
	if err != nil {
		return systemRole, err
	}
	err = copier.Copy(systemRole, systemRoleModel)
	return systemRole, err
}

// Delete 删除
func (repository *SystemRoleRepository) Delete(systemRole *entity.SystemRole) (*entity.SystemRole, error) {
	systemRoleModel := &model.SystemRole{}
	_ = copier.Copy(systemRoleModel, systemRole)
	err := repository.context.Transaction().Context.Delete(systemRoleModel).Error
	return systemRole, err
}
