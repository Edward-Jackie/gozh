package dao

import "github.com/Edward-Jackie/gotool/pkg/web"

type SystemRoleDao struct {
	context *web.Context
}

func NewSystemRoleDao(context *web.Context) *SystemRoleDao {
	return &SystemRoleDao{context: context}
}
