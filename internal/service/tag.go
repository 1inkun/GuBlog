package service

import (
	"context"

	"github.com/1inkun/GuBlog/internal/database"
	"github.com/1inkun/GuBlog/internal/utils"
)

type TagService struct {
	tagQuery *database.TagQuery
}

func NewTagService(tagQuery *database.TagQuery) *TagService {
	return &TagService{
		tagQuery: tagQuery,
	}
}

func (s *TagService) GetTagList(ctx context.Context, page, pageSize int64) (utils.Pagination, error) {
	var res utils.Pagination
	offset, pageSize := utils.CalcOffset(page, pageSize)
	res.Page = page
	if page == 0 {
		res.Page = 1
	}
	res.PageSize = pageSize
	// 获取Tag数据列表
	data, err := s.tagQuery.GetTagList(ctx, pageSize, offset)
	if err != nil {
		return res, err
	}
	// 组装响应
	res.Data = data
	return res, nil
}
