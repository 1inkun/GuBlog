package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/1inkun/GuBlog/internal/types"
	"github.com/redis/go-redis/v9"
)

type PostQuery struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewPostQuery(db *sql.DB, rdb *redis.Client) *PostQuery {
	return &PostQuery{
		db:  db,
		rdb: rdb,
	}
}

func (q *PostQuery) GetPostsId(ctx context.Context, pageSize int64, offset int64) ([]int64, error) {
	query := "SELECT `id` FROM `posts` ORDER BY `create_at` DESC LIMIT ? OFFSET ?;"
	var id_list = make([]int64, 0, pageSize)
	rows, err := q.db.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		log.Printf("在获取文章列表时出现错误:%v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var postId int64
		err := rows.Scan(&postId)
		if err != nil {
			log.Printf("在获取文章列表时出现错误:%v", err)
			return nil, err
		}
		id_list = append(id_list, postId)
	}
	if err = rows.Err(); err != nil {
		log.Printf("在获取文章列表时出现错误:%v", err)
		return nil, err
	}
	return id_list, nil
}

func (q *PostQuery) GetPostTagsByPostId(ctx context.Context, postId int64) ([]types.TagData, error) {
	query := "SELECT `tag_id` FROM `post_tags` WHERE `post_id` = ?;"
	rows, err := q.db.QueryContext(ctx, query, postId)
	if err != nil {
		log.Printf("在获取文章标签时出现错误:%v", err)
		return nil, err
	}
	defer rows.Close()
	var tagIds []any
	for rows.Next() {
		var t int64
		err := rows.Scan(&t)
		if err != nil {
			log.Printf("在获取文章标签ID时出现错误:%v", err)
			return nil, err
		}
		tagIds = append(tagIds, t)
	}
	if err = rows.Err(); err != nil {
		log.Printf("在获取文章标签ID时出现错误:%v", err)
		return nil, err
	}
	if len(tagIds) < 1 {
		log.Println("没有设置标签")
		return nil, nil
	}
	// 按照tagId列表查找对应的tag
	// 处理占位符个数
	placeholders := strings.Repeat("?,", len(tagIds))
	placeholders = placeholders[:len(placeholders)-1]
	query = fmt.Sprintf("SELECT id, name, usage_count FROM tags WHERE id IN (%s) ORDER BY usage_count DESC;", placeholders)
	var tagDatas = make([]types.TagData, 0, len(tagIds))
	rows, err = q.db.QueryContext(ctx, query, tagIds...)
	if err != nil {
		log.Printf("在获取文章标签数据时出现错误:%v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t types.TagData
		err := rows.Scan(&t.Id, &t.Name, &t.Usage)
		if err != nil {
			log.Printf("在获取文章标签数据时出现错误:%v", err)
			return nil, err
		}
		tagDatas = append(tagDatas, t)
	}
	if err = rows.Err(); err != nil {
		log.Printf("在获取文章标签数据时出现错误:%v", err)
		return nil, err
	}
	return tagDatas, nil
}

func (q *PostQuery) GetPostIdByTagId(ctx context.Context, tagIds []int64, pageSize int64, offset int64) ([]int64, error) {
	var postIds = make([]int64, 0, pageSize)
	var args = make([]any, 0, len(tagIds)+2)
	// 处理tagIds
	for _, v := range tagIds {
		args = append(args, v)
	}
	args = append(args, pageSize, offset)
	// 处理占位符个数
	placeholders := strings.Repeat("?,", len(tagIds))
	placeholders = placeholders[:len(placeholders)-1]
	query := fmt.Sprintf("SELECT `post_id` FROM `post_tags` WHERE `tag_id` IN (%s) GROUP BY `post_id` ORDER BY MAX(`create_at`) DESC, `post_id` DESC LIMIT ? OFFSET ?;", placeholders)
	rows, err := q.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("在通过tag获取文章id时发生错误:%v", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t int64
		if err := rows.Scan(&t); err != nil {
			log.Printf("在通过tag获取文章id时发生错误:%v", err)
			return nil, err
		}
		postIds = append(postIds, t)
	}
	if err := rows.Err(); err != nil {
		log.Printf("在通过tag获取文章id时发生错误:%v", err)
		return nil, err
	}
	return postIds, nil
}

func (q *PostQuery) GetPostDataById(ctx context.Context, postId int64) (types.PostData, error) {
	var data types.PostData
	query := "SELECT `id`, `title`, `author_id`, `body`, `view_count`, `publish_at` FROM `posts` WHERE `id` = ? AND `status` = 1;"
	row := q.db.QueryRowContext(ctx, query, postId)
	var publishAt sql.NullTime
	err := row.Scan(&data.Id, &data.Title, &data.AuthorId, &data.Body, &data.ViewCount, &publishAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("在查询文章数据时找不到数据:%v", types.ErrPostNotExist)
			return data, types.ErrPostNotExist
		}
		log.Printf("在查询文章数据时出现错误:%v", err)
		return data, err
	}
	if publishAt.Valid {
		data.PublishAt = publishAt.Time
	}
	return data, nil
}

func (q *PostQuery) GetPostAuthorDataById(ctx context.Context, authorId int64) (types.PostAuthor, error) {
	query := "SELECT `nickname`, `avatar_file`, `create_at` FROM `authors` WHERE `id` = ?;"
	var data types.PostAuthor
	row := q.db.QueryRowContext(ctx, query, authorId)
	err := row.Scan(&data.Nickname, &data.AvatarFile, &data.CreateAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("在查询文章作者数据时找不到数据:%v", types.ErrAuthorNotExist)
			return data, types.ErrAuthorNotExist
		}
		log.Printf("在查询文章作者数据时出现错误:%v", err)
		return data, err
	}
	return data, nil
}
