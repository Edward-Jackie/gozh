package service

import (
	"fmt"
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/tools"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/Edward-Jackie/gotool/pkg/web/auth"
	"github.com/gookit/goutil/arrutil"
	"gozh/internal/system/application/setting/dto"
	"gozh/internal/system/application/system/command"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/facade"
	"gozh/internal/system/infrastructure/cache"
	"strconv"
)

type AuthService struct {
	context *web.Context
}

func NewAuthService(context *web.Context) *AuthService {
	return &AuthService{context: context}
}

// Login 登录
func (service *AuthService) Login(login *command.Login) (interface{}, error) {
	if err := service.context.Validate(login); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	user, err := facade.SystemUserRepository(service.context).FindOne(&orm.Conditions{
		Equal: map[string]interface{}{"userName": login.UserName},
	})
	if err != nil {
		return nil, web.ThrowError(web.ArgError, "该用户不存在")
	}
	password, err := tools.DecryptPassword(user.Password)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, "您输入的密码不正确")
	}
	if password != login.Password {
		return nil, web.ThrowError(web.InternalServerError, "您输入的密码不正确")
	}
	user.LastLoginIp = service.context.ClientIP()
	user.LastLoginTime = tools.Now()
	_, err = facade.SystemUserRepository(service.context).Save(user)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, "更新用户信息失败")
	}
	token, err := auth.CreateTokenString(&auth.User{UserId: user.Id, UserName: user.UserName, NickName: user.NickName})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, "生成token失败")
	}
	_ = NewSystemLoggerService(service.context).Save(&command.CreateSystemLogger{
		UserId:     user.Id,
		Comment:    "用户登录",
		LoggerType: entity.LoggerTypeLogin,
	})
	return map[string]interface{}{
		"token":    token,
		"userInfo": user,
	}, nil
}

// User 用户登录信息
func (service *AuthService) User() (interface{}, error) {
	authUser := service.context.Auth()
	//用户信息
	user, err := facade.SystemUserRepository(service.context).FindById(authUser.UserId)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, "用户不存在")
	}
	return user, nil
}

// Menu 获取登录用户菜单
func (service *AuthService) Menu() (interface{}, error) {
	list, err := cache.GetMenu(service.context)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, "获取菜单失败")
	}
	//获取用户角色
	userInfo, err := cache.GetUserById(service.context, service.context.Auth().UserId)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, "获取用户失败")
	}
	checkMenus := make([]*entity.SettingMenu, 0)
	//获取对应的角色
	if userInfo.UserName == "admin" {
		checkMenus = list
	} else {
		checkIds := make([]int64, 0)
		for _, roleId := range userInfo.RoleIds {
			r, _ := strconv.Atoi(roleId)
			if r > 0 {
				role, err := cache.GetRoleById(service.context, int64(r))
				if err == nil && role.Id > 0 {
					for _, p := range role.Permissions {
						checkIds = append(checkIds, p)
						parentIds, err := cache.GetMenuParentIds(service.context, p)
						if err == nil && len(parentIds) > 0 {
							checkIds = append(checkIds, parentIds...)
						}
					}
				}
			}
		}
		for _, item := range list {
			if arrutil.InStrings(fmt.Sprintf("%v", item.Id), arrutil.MustToStrings(checkIds)) {
				checkMenus = append(checkMenus, item)
			}
		}
	}
	result := dto.TreeMenu(checkMenus, 0)
	return map[string]interface{}{
		"menu": result,
	}, nil
}

// Dict 获取数据字典
func (service *AuthService) Dict() (interface{}, error) {
	_, list, err := facade.SettingDictRepository(service.context).Find(map[string]interface{}{})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, "获取数据字典失败")
	}
	_, values, err := facade.SettingDictValuesRepository(service.context).Find(&orm.Conditions{})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, "获取字典值失败")
	}
	result := make(map[string]interface{})
	for _, item := range list {
		data := make([]*entity.DictOption, 0)
		for _, val := range values {
			if val.DictId == item.Id {
				data = append(data, &entity.DictOption{
					Label: val.Value,
					Value: val.Key,
				})
			}
		}
		result[item.Code] = data
	}
	return result, nil
}
