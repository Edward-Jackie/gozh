package service

import (
	"fmt"
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/structs"
	"github.com/jinzhu/copier"
	"gozh/internal/system/application/system/command"
	"gozh/internal/system/application/system/dto"
	"gozh/internal/system/application/system/query"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/facade"
	"gozh/internal/system/infrastructure/cache"
	"strconv"
)

type SystemRoleService struct {
	context *web.Context
}

func NewSystemRoleService(context *web.Context) *SystemRoleService {
	return &SystemRoleService{context: context}
}

// Create 新增
func (service *SystemRoleService) Create(createSystemRole *command.CreateSystemRole) (interface{}, error) {
	if err := service.context.Validate(createSystemRole); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	systemRole := &entity.SystemRole{}
	_ = copier.Copy(systemRole, createSystemRole)
	systemRole, err := facade.SystemRoleRepository(service.context).Save(systemRole)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	//清除缓存
	cache.RemoveRoles()
	_ = NewSystemLoggerService(service.context).Save(&command.CreateSystemLogger{
		UserId:     service.context.Auth().UserId,
		Comment:    "新增角色:" + systemRole.Name,
		LoggerType: entity.LoggerTypeLogin,
	})
	return systemRole, nil
}

// Update 修改
func (service *SystemRoleService) Update(updateSystemRole *command.UpdateSystemRole) (interface{}, error) {
	if err := service.context.Validate(updateSystemRole); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	systemRoleRepository := facade.SystemRoleRepository(service.context)
	systemRole, err := systemRoleRepository.FindById(updateSystemRole.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	systemRole.Update(structs.ToMap(updateSystemRole))
	systemRole, err = facade.SystemRoleRepository(service.context).Save(systemRole)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	//清除缓存
	cache.RemoveRoles()
	_ = NewSystemLoggerService(service.context).Save(&command.CreateSystemLogger{
		UserId:     service.context.Auth().UserId,
		Comment:    "编辑角色:" + systemRole.Name,
		LoggerType: entity.LoggerTypeLogin,
	})
	return systemRole, nil
}

// List 列表
func (service *SystemRoleService) List(listSystemRole *query.ListSystemRole) (interface{}, error) {
	total, list, err := facade.SystemRoleRepository(service.context).Find(&orm.Conditions{
		Pagination: &orm.Pagination{
			Page:     listSystemRole.Page,
			PageSize: listSystemRole.PageSize,
		},
		Like: map[string]interface{}{
			"name": listSystemRole.Name,
		},
		OrderBy:     listSystemRole.OrderBy,
		OrderColumn: listSystemRole.OrderColumn,
	})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	result := make([]*dto.SystemRole, 0)
	for _, item := range list {
		role := &dto.SystemRole{
			Id:          item.Id,
			Code:        item.Code,
			Name:        item.Name,
			Permissions: make([]string, 0),
			Comment:     item.Comment,
		}
		if len(item.Permissions) > 0 {
			role.Permissions = arrutil.MustToStrings(item.Permissions)
		}
		result = append(result, role)
	}
	return web.NewPagination(total, result), nil
}

// Get 获取信息
func (service *SystemRoleService) Get(getSystemRole *query.GetSystemRole) (interface{}, error) {
	if err := service.context.Validate(getSystemRole); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	systemRole, err := facade.SystemRoleRepository(service.context).FindOne(&orm.Conditions{
		Equal: map[string]interface{}{"id": getSystemRole.Id},
	})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return systemRole, nil
}

// Delete 删除
func (service *SystemRoleService) Delete(deleteSystemRole *command.DeleteSystemRole) (interface{}, error) {
	if err := service.context.Validate(deleteSystemRole); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	systemRole, err := facade.SystemRoleRepository(service.context).FindById(deleteSystemRole.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	systemRole, err = facade.SystemRoleRepository(service.context).Delete(systemRole)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	//清除缓存
	cache.RemoveRoles()
	_ = NewSystemLoggerService(service.context).Save(&command.CreateSystemLogger{
		UserId:     service.context.Auth().UserId,
		Comment:    "删除角色:" + systemRole.Name,
		LoggerType: entity.LoggerTypeLogin,
	})
	return systemRole, nil
}

// GetPermissions 获取菜单权限
func (service *SystemRoleService) GetPermissions() (interface{}, error) {
	_, menus, err := facade.SettingMenuRepository(service.context).Find(map[string]interface{}{})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return dto.TreeMenu(menus, 0), nil
}

// SavePermissions 保存菜单权限
func (service *SystemRoleService) SavePermissions(permissions *command.SavePermissions) (interface{}, error) {
	if err := service.context.Validate(permissions); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	role, err := facade.SystemRoleRepository(service.context).FindById(permissions.RoleId)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	permissionIds := make([]int64, 0)
	for _, item := range permissions.PermissionIds {
		pid, err := strconv.Atoi(item)
		if err != nil {
			return nil, web.ThrowError(web.ArgError, "权限值不正确")
		}
		permissionIds = append(permissionIds, int64(pid))
	}
	role.Update(map[string]interface{}{"Permissions": permissionIds})
	role, err = facade.SystemRoleRepository(service.context).Save(role)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, "权限保存失败")
	}
	//清除缓存
	cache.RemoveRoles()
	_ = NewSystemLoggerService(service.context).Save(&command.CreateSystemLogger{
		UserId:     service.context.Auth().UserId,
		Comment:    "角色[" + role.Name + "]设置权限",
		LoggerType: entity.LoggerTypeLogin,
	})
	return role, nil
}

// All 获取所有角色
func (service *SystemRoleService) All() (interface{}, error) {
	_, list, err := facade.SystemRoleRepository(service.context).Find(&orm.Conditions{})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	result := make([]entity.CommonLabelValue, 0)
	for _, item := range list {
		result = append(result, entity.CommonLabelValue{
			Label: item.Name,
			Value: fmt.Sprintf("%v", item.Id),
		})
	}
	return result, nil
}
