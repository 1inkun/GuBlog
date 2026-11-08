package service

import (
	"context"

	"github.com/1inkun/GuBlog/internal/database"
	"github.com/1inkun/GuBlog/internal/types"
	"github.com/1inkun/GuBlog/internal/utils"
	"golang.org/x/sync/errgroup"
)

type PostService struct {
	postQuery *database.PostQuery
}

func NewPostService(postQuery *database.PostQuery) *PostService {
	return &PostService{
		postQuery: postQuery,
	}
}

func (s *PostService) GetPostList(ctx context.Context, page int64, pageSize int64) (utils.Pagination, error) {
	var res utils.Pagination
	offset, pageSize := utils.CalcOffset(page, pageSize)
	res.Page = page
	res.PageSize = pageSize
	postIds, err := s.postQuery.GetPostsId(ctx, pageSize, offset)
	if err != nil {
		return res, err
	}
	if len(postIds) < 1 {
		return res, nil
	}
	var data = make([]types.PostListRes, len(postIds))
	var eg errgroup.Group
	for k, v := range postIds {
		eg.Go(func() error {
			tagData, err := s.postQuery.GetPostTagsByPostId(ctx, v)
			if err != nil {
				return err
			}
			var t = types.PostListRes{
				PostId: v,
				Tags:   tagData,
			}
			data[k] = t
			return nil
		})
	}
	if err = eg.Wait(); err != nil {
		return res, err
	}
	// 组装返回数据
	res.Data = data
	return res, nil
}
