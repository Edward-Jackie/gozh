package dao

import "github.com/Edward-Jackie/gotool/pkg/web"

type SystemUserDao struct {
	context *web.Context
}

func NewSystemUserDao(context *web.Context) *SystemUserDao {
	return &SystemUserDao{context: context}
}
