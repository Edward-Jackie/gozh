package dao

import "github.com/Edward-Jackie/gotool/pkg/web"

type SysUserDao struct {
	context *web.Context
}

func NewSysUserDao(context *web.Context) *SysUserDao {
	return &SysUserDao{context: context}
}
