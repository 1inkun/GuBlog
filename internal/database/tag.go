package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/1inkun/GuBlog/internal/types"
	"github.com/redis/go-redis/v9"
)

type TagQuery struct {
	db  *sql.DB
	rdb *redis.Client
}

func NewTagQuery(db *sql.DB, rdb *redis.Client) *TagQuery {
	return &TagQuery{
		db:  db,
		rdb: rdb,
	}
}

func (q *TagQuery) GetTagList(ctx context.Context, pageSize, offset int64) ([]types.TagData, error) {
	var data = make([]types.TagData, 0, pageSize)
	query := "SELECT `id`, `name`, `usage_count` FROM `tags` ORDER BY `usage_count` DESC, `id` ASC LIMIT ? OFFSET ?;"
	rows, err := q.db.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		log.Printf("在获取tag列表时出错:%v", err)
		return data, err
	}
	defer rows.Close()
	for rows.Next() {
		var t types.TagData
		err := rows.Scan(&t.Id, &t.Name, &t.Usage)
		if err != nil {
			log.Printf("在获取tag列表时出错:%v", err)
			return data, err
		}
		data = append(data, t)
	}
	if err := rows.Err(); err != nil {
		log.Printf("在获取tag列表时出错:%v", err)
		return data, err
	}
	return data, nil
}
