package system

import (
	"github.com/Edward-Jackie/gotool/pkg/web"
	"github.com/jinzhu/copier"
	"gozh/internal/system/api/route"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/infrastructure/cache"
)

func Initialization() {
	web.Default().InitRoute(
		route.Auth,
		route.SettingMenu,
		route.SettingDict,
		route.SettingDictValues,
		route.SystemUser,
		route.SystemRole,
		route.SystemEvent,
	)
}

type User struct {
	entity.SystemUser
}

// GetUserInfo 获取用户信息
func GetUserInfo(context *web.Context, userId int64) (*User, error) {
	systemUser, err := cache.GetUserById(context, userId)
	if err != nil {
		return nil, err
	}
	user := &User{}
	err = copier.Copy(user, systemUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUsers 获取所有用户
func GetUsers(context *web.Context) ([]*User, error) {
	users := make([]*User, 0)
	systemUsers, err := cache.GetUsers(context)
	if err != nil {
		return users, err
	}
	for _, item := range systemUsers {
		user := &User{}
		err = copier.Copy(user, item)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
