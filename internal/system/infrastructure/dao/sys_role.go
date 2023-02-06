package dao

import "github.com/Edward-Jackie/gotool/pkg/web"

type SysRoleDao struct {
	context *web.Context
}

func NewSysRoleDao(context *web.Context) *SysRoleDao {
	return &SysRoleDao{context: context}
}
