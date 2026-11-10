package database

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/1inkun/GuBlog/internal/types"
	"github.com/redis/go-redis/v9"
)

type AuthorQuery struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewAuthorQuery(db *sql.DB, rdb *redis.Client) *AuthorQuery {
	return &AuthorQuery{
		db:  db,
		rdb: rdb,
	}
}

func (q *AuthorQuery) GetAuthorDataById(ctx context.Context, authorId int64) (types.AuthorRes, error) {
	var res types.AuthorRes
	var avatar, github, twitter, bilibili sql.NullString
	query := "SELECT `nickname`, `avatar_file`, `email`, `github`, `twitter`, `bilibili`, `create_at` FROM `authors` WHERE `id` = ?;"
	row := q.db.QueryRowContext(ctx, query, authorId)
	err := row.Scan(&res.Nickname, &avatar, &res.Email, &github, &twitter, &bilibili, &res.CreateAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, types.ErrAuthorNotExist
		}
		log.Printf("通过Id查找作者信息时出现错误:%v", err)
		return res, err
	}
	res.AvatarFile = nullToStr(avatar)
	res.GitHub = nullToStr(github)
	res.Twitter = nullToStr(twitter)
	res.Bilibili = nullToStr(bilibili)
	return res, nil
}

func nullToStr(v sql.NullString) string {
	if v.Valid {
		return v.String
	}
	return ""
}
