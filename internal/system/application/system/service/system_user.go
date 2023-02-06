package service

import (
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/strutil"
	"github.com/jinzhu/copier"
	"gozh/internal/system/application/system/command"
	"gozh/internal/system/application/system/dto"
	"gozh/internal/system/application/system/query"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/facade"
	"gozh/internal/system/infrastructure/cache"
	"strconv"
	"strings"
	"time"
)

type SystemUserService struct {
	context *web.Context
}

func NewSystemUserService(context *web.Context) *SystemUserService {
	return &SystemUserService{context: context}
}

// Create 新增
func (service *SystemUserService) Create(createSystemUser *command.CreateSystemUser) (interface{}, error) {
	err := service.context.Validate(createSystemUser)
	if err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	systemUser := &entity.SystemUser{}
	_ = copier.Copy(systemUser, createSystemUser)
	systemUser.Password = strutil.Md5(entity.DefaultPassword)
	systemUser.Password, err = tools.EncryptPassword(systemUser.Password)
	if err != nil {
		return nil, web.ThrowError(web.ArgError, "生成加密密码失败")
	}
	mTime := time.Now()
	systemUser.RegisterTime = &mTime
	systemUser, err = facade.SystemUserRepository(service.context).Save(systemUser)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	//清除缓存
	cache.RemoveUsers()
	_ = NewSystemLoggerService(service.context).Save(&command.CreateSystemLogger{
		UserId:     service.context.Auth().UserId,
		Comment:    "新增用户:" + systemUser.UserName + "[" + systemUser.NickName + "]",
		LoggerType: entity.LoggerTypeLogin,
	})
	return systemUser, nil
}

// Update 修改
func (service *SystemUserService) Update(updateSystemUser *command.UpdateSystemUser) (interface{}, error) {
	if err := service.context.Validate(updateSystemUser); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	systemUserRepository := facade.SystemUserRepository(service.context)
	systemUser, err := systemUserRepository.FindById(updateSystemUser.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	if systemUser.UserName == "admin" && systemUser.UserName != updateSystemUser.UserName {
		return nil, web.ThrowError(web.ArgError, "管理员账户用户名不可修改")
	}
	systemUser.Update(structs.ToMap(updateSystemUser))
	systemUser, err = facade.SystemUserRepository(service.context).Save(systemUser)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	//清除缓存
	cache.RemoveUsers()
	_ = NewSystemLoggerService(service.context).Save(&command.CreateSystemLogger{
		UserId:     service.context.Auth().UserId,
		Comment:    "编辑用户:" + systemUser.UserName + "[" + systemUser.NickName + "]",
		LoggerType: entity.LoggerTypeLogin,
	})
	return systemUser, nil
}

// List 列表
func (service *SystemUserService) List(listSystemUser *query.ListSystemUser) (interface{}, error) {
	total, list, err := facade.SystemUserRepository(service.context).Find(&orm.Conditions{
		Pagination: &orm.Pagination{
			Page:     listSystemUser.Page,
			PageSize: listSystemUser.PageSize,
		},
		Like: map[string]interface{}{
			"nickName":      listSystemUser.NickName,
			"userName":      listSystemUser.UserName,
			"lastLoginTime": listSystemUser.LastLoginDate,
		},
		OrderColumn: listSystemUser.OrderColumn,
		OrderBy:     listSystemUser.OrderBy,
	})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	result := make([]*dto.SystemUser, 0)
	for _, item := range list {
		user := &dto.SystemUser{
			Id:            item.Id,
			UserName:      item.UserName,
			NickName:      item.NickName,
			Phone:         item.Phone,
			RoleIds:       item.RoleIds,
			Roles:         "",
			RegisterTime:  item.RegisterTime,
			LastLoginTime: item.LastLoginTime,
			LastLoginIp:   item.LastLoginIp,
		}
		roles := make([]string, 0)
		for _, r := range item.RoleIds {
			roleId, _ := strconv.Atoi(r)
			if roleId > 0 {
				role, err := cache.GetRoleById(service.context, int64(roleId))
				if err == nil && role.Id > 0 {
					roles = append(roles, role.Name)
				}
			}
		}
		user.Roles = strings.Join(roles, ",")
		result = append(result, user)
	}
	return web.NewPagination(total, result), nil
}

// Get 获取信息
func (service *SystemUserService) Get(getSystemUser *query.GetSystemUser) (interface{}, error) {
	if err := service.context.Validate(getSystemUser); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	systemUser, err := facade.SystemUserRepository(service.context).FindOne(&orm.Conditions{
		Equal: map[string]interface{}{"id": getSystemUser.Id},
	})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return systemUser, nil
}

// Delete 删除
func (service *SystemUserService) Delete(deleteSystemUser *command.DeleteSystemUser) (interface{}, error) {
	if err := service.context.Validate(deleteSystemUser); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	systemUser, err := facade.SystemUserRepository(service.context).FindById(deleteSystemUser.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	if systemUser.UserName == "admin" {
		return nil, web.ThrowError(web.ArgError, "管理员账号不可删除")
	}
	systemUser, err = facade.SystemUserRepository(service.context).Delete(systemUser)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	//清除缓存
	cache.RemoveUsers()
	_ = NewSystemLoggerService(service.context).Save(&command.CreateSystemLogger{
		UserId:     service.context.Auth().UserId,
		Comment:    "删除用户:" + systemUser.UserName + "[" + systemUser.NickName + "]",
		LoggerType: entity.LoggerTypeLogin,
	})
	return systemUser, nil
}

// ModifyPassword 修改密码
func (service *SystemRoleService) ModifyPassword(modifyPassword *command.ModifyPassword) (interface{}, error) {
	if err := service.context.Validate(modifyPassword); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	if modifyPassword.Password != modifyPassword.ConfirmPassword {
		return nil, web.ThrowError(web.ArgError, "两次输入密码不一致")
	}
	userInfo, err := facade.SystemUserRepository(service.context).FindById(service.context.Auth().UserId)
	if err != nil {
		return nil, web.ThrowError(web.ArgError, "未查询到用户")
	}
	password, err := tools.DecryptPassword(userInfo.Password)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	if modifyPassword.OldPassword != password {
		return nil, web.ThrowError(web.ArgError, "旧密码不正确")
	}
	userInfo.Password, _ = tools.EncryptPassword(modifyPassword.Password)
	userInfo, err = facade.SystemUserRepository(service.context).Save(userInfo)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	_ = NewSystemLoggerService(service.context).Save(&command.CreateSystemLogger{
		UserId:     service.context.Auth().UserId,
		Comment:    "修改密码:" + userInfo.UserName + "[" + userInfo.NickName + "]",
		LoggerType: entity.LoggerTypeLogin,
	})
	return userInfo, nil
}

// ResetPassword 重置密码
func (service *SystemUserService) ResetPassword(resetPassword *command.ResetPassword) (interface{}, error) {
	if err := service.context.Validate(resetPassword); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	if resetPassword.Password != resetPassword.ConfirmPassword {
		return nil, web.ThrowError(web.ArgError, "两次输入密码不一致")
	}
	userInfo, err := facade.SystemUserRepository(service.context).FindById(resetPassword.UserId)
	if err != nil {
		return nil, web.ThrowError(web.ArgError, "未查询到用户")
	}
	userInfo.Password, _ = tools.EncryptPassword(resetPassword.Password)
	userInfo, err = facade.SystemUserRepository(service.context).Save(userInfo)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	_ = NewSystemLoggerService(service.context).Save(&command.CreateSystemLogger{
		UserId:     service.context.Auth().UserId,
		Comment:    "重置密码:" + userInfo.UserName + "[" + userInfo.NickName + "]",
		LoggerType: entity.LoggerTypeLogin,
	})
	return userInfo, nil
}
