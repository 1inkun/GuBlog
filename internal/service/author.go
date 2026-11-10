package service

import (
	"context"

	"github.com/1inkun/GuBlog/internal/database"
	"github.com/1inkun/GuBlog/internal/types"
)

type AuthorService struct {
	authorQuery *database.AuthorQuery
}

func NewAuthorService(authorQuery *database.AuthorQuery) *AuthorService {
	return &AuthorService{
		authorQuery: authorQuery,
	}
}

func (s *AuthorService) GetAuthorDataById(ctx context.Context, authorId int64) (types.AuthorRes, error) {
	res, err := s.authorQuery.GetAuthorDataById(ctx, authorId)
	if err != nil {
		return types.AuthorRes{}, err
	}
	return res, err
}
