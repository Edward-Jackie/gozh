package dao

import "github.com/Edward-Jackie/gotool/pkg/web"

type SysMenuDao struct {
	context *web.Context
}

func NewSysMenuDao(context *web.Context) *SysMenuDao {
	return &SysMenuDao{context: context}
}
