package database

import (
	"context"
	"database/sql"
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
		log.Printf("在获取文章列表时出现错误:%s", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var postId int64
		err := rows.Scan(&postId)
		if err != nil {
			log.Printf("在获取文章列表时出现错误:%s", err)
			return nil, err
		}
		id_list = append(id_list, postId)
	}
	if err = rows.Err(); err != nil {
		log.Printf("在获取文章列表时出现错误:%s", err)
		return nil, err
	}
	return id_list, nil
}

func (q *PostQuery) GetPostTagsByPostId(ctx context.Context, postId int64) ([]types.TagData, error) {
	query := "SELECT `tag_id` FROM `post_tags` WHERE `post_id` = ?;"
	rows, err := q.db.QueryContext(ctx, query, postId)
	if err != nil {
		log.Printf("在获取文章标签时出现错误:%s", err)
		return nil, err
	}
	defer rows.Close()
	var tagIds []any
	for rows.Next() {
		var t int64
		err := rows.Scan(&t)
		if err != nil {
			log.Printf("在获取文章标签ID时出现错误:%s", err)
			return nil, err
		}
		tagIds = append(tagIds, t)
	}
	if err = rows.Err(); err != nil {
		log.Printf("在获取文章标签ID时出现错误:%s", err)
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
		log.Printf("在获取文章标签数据时出现错误:%s", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t types.TagData
		err := rows.Scan(&t.Id, &t.Name, &t.Usage)
		if err != nil {
			log.Printf("在获取文章标签数据时出现错误:%s", err)
			return nil, err
		}
		tagDatas = append(tagDatas, t)
	}
	if err = rows.Err(); err != nil {
		log.Printf("在获取文章标签数据时出现错误:%s", err)
		return nil, err
	}
	return tagDatas, nil
}
