package types

import "time"

type CategoryData struct {
	Id          int64     `json:"id"`
	ParentId    int64     `json:"parent_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
}

type PostCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryRes struct {
	Id          int64  `json:"id"`
	ParentId    int64  `json:"parent_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
