package dao

import "github.com/Edward-Jackie/gotool/pkg/web"

type VideoCategoryDao struct {
	context *web.Context
}

func NewVideoCategoryDao(context *web.Context) *VideoCategoryDao {
	return &VideoCategoryDao{context: context}
}
