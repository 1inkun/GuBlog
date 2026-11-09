package types

import "time"

type PostData struct {
	Id           int64     `json:"id"`
	AuthorId     int64     `json:"author_id"`
	CategoriesId int64     `json:"categories_id"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	PublishAt    time.Time `json:"publish_at"`
	Status       int64     `json:"status"`
	ViewCount    int64     `json:"view_count"`
	UpdateAt     time.Time `json:"update_at"`
	CreateAt     time.Time `json:"create_at"`
}

type PostDataRes struct {
	Title     string     `json:"title"`
	Author    PostAuthor `json:"author"`
	Tags      []TagData  `json:"tags"`
	Body      string     `json:"body"`
	ViewCount int64      `json:"view_count"`
	PublishAt time.Time  `json:"publish_at"`
}

type PostListRes struct {
	PostId int64     `json:"post_id"`
	Tags   []TagData `json:"tags"`
}
