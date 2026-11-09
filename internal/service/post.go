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
	if page == 0 {
		res.Page = 1
	}
	res.PageSize = pageSize
	postIds, err := s.postQuery.GetPostsId(ctx, pageSize, offset)
	if err != nil {
		return res, err
	}
	if len(postIds) < 1 {
		return res, nil
	}
	var data = make([]types.PostListRes, len(postIds))
	eg, ctx := errgroup.WithContext(ctx)
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

func (s *PostService) GetPostListWithTags(ctx context.Context, tagIds []int64, page int64, pageSize int64) (utils.Pagination, error) {
	var res utils.Pagination
	offset, pageSize := utils.CalcOffset(page, pageSize)
	res.Page = page
	if page == 0 {
		res.Page = 1
	}
	res.PageSize = pageSize
	// 先通过tag查出满足要求的文章id
	postIds, err := s.postQuery.GetPostIdByTagId(ctx, tagIds, pageSize, offset)
	if err != nil {
		return res, err
	}
	if len(postIds) == 0 {
		// 没找到满足条件数据的情况
		return res, nil
	}
	// 再通过postId去查
	var data = make([]types.PostListRes, len(postIds))
	eg, gctx := errgroup.WithContext(ctx)
	for k, v := range postIds {
		eg.Go(func() error {
			tagData, err := s.postQuery.GetPostTagsByPostId(gctx, v)
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

func (s *PostService) GetPostDataById(ctx context.Context, postId int64) (types.PostDataRes, error) {
	var res types.PostDataRes
	// 获取基本信息
	rawPostData, err := s.postQuery.GetPostDataById(ctx, postId)
	if err != nil {
		return res, err
	}
	postId = rawPostData.Id
	authorId := rawPostData.AuthorId
	var postAuthorData types.PostAuthor
	var tagDatas []types.TagData

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// 获取作者信息
		data, err := s.postQuery.GetPostAuthorDataById(ctx, authorId)
		if err != nil {
			return err
		}
		postAuthorData = data
		return nil
	})
	// 获取tag信息
	eg.Go(func() error {
		datas, err := s.postQuery.GetPostTagsByPostId(ctx, postId)
		if err != nil {
			return err
		}
		tagDatas = datas
		return nil
	})
	// 统一处理错误
	if err := eg.Wait(); err != nil {
		return res, err
	}
	// 组装数据
	res.Tags = tagDatas
	res.Author = postAuthorData
	res.Title = rawPostData.Title
	res.Body = rawPostData.Body
	res.PublishAt = rawPostData.PublishAt
	res.ViewCount = rawPostData.ViewCount
	return res, nil
}
