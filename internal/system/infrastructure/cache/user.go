package cache

import (
	"errors"
	"github.com/Edward-Jackie/gotool/pkg/cache"
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/infrastructure/repository"
)

var cacheUserKey = "users"

var cacheUserTime = int64(3600 * 24 * 30)

func GetUsers(context *web.Context) (map[int64]*entity.SystemUser, error) {
	result := make(map[int64]*entity.SystemUser)
	_ = cache.Default().Get(cacheUserKey).ToStruct(&result)
	if len(result) > 0 {
		return result, nil
	}
	_, data, err := repository.NewSystemUserRepository(context).Find(&orm.Conditions{})
	if err != nil {
		return result, err
	}
	for _, item := range data {
		result[item.Id] = item
	}
	err = cache.Default().Set(cacheUserKey, result, cacheUserTime)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetUserById(context *web.Context, id int64) (*entity.SystemUser, error) {
	data, err := GetUsers(context)
	if err != nil {
		return nil, err
	}
	if _, ok := data[id]; ok {
		return data[id], nil
	} else {
		return nil, errors.New("用户不存在")
	}
}

func RemoveUsers() {
	_ = cache.Default().Remove(cacheUserKey)
}
