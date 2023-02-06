package cache

import (
	"errors"
	"github.com/Edward-Jackie/gotool/pkg/cache"
	"github.com/Edward-Jackie/gotool/pkg/database/orm"
	"github.com/Edward-Jackie/gotool/pkg/web"
	"gozh/internal/system/domain/entity"
	"gozh/internal/system/infrastructure/repository"
)

var cacheRoleKey = "roles"

var cacheRoleTime = int64(3600 * 24 * 30)

func GetRoles(context *web.Context) (map[int64]*entity.SystemRole, error) {
	result := make(map[int64]*entity.SystemRole)
	_ = cache.Default().Get(cacheRoleKey).ToStruct(&result)
	if len(result) > 0 {
		return result, nil
	}
	_, data, err := repository.NewSystemRoleRepository(context).Find(&orm.Conditions{})
	if err != nil {
		return result, err
	}
	for _, item := range data {
		result[item.Id] = item
	}
	err = cache.Default().Set(cacheRoleKey, result, cacheRoleTime)
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetRoleById(context *web.Context, id int64) (*entity.SystemRole, error) {
	data, err := GetRoles(context)
	if err != nil {
		return nil, err
	}
	if _, ok := data[id]; ok {
		return data[id], nil
	} else {
		return nil, errors.New("角色不存在")
	}
}

func RemoveRoles() {
	_ = cache.Default().Remove(cacheRoleKey)
}
