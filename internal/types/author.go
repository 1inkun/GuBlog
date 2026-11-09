package types

import "time"

type AuthorData struct {
	Id         int64     `json:"id"`
	Nickname   string    `json:"nickname"`
	AvatarFile string    `json:"avatar_file"`
	Email      string    `json:"email"`
	GitHub     string    `json:"github"`
	Twitter    string    `json:"twitter"`
	Bilibili   string    `json:"bilibili"`
	Active     string    `json:"active"`
	CreateAt   time.Time `json:"create_at"`
}

type PostAuthor struct {
	Nickname   string    `json:"nickname"`
	AvatarFile string    `json:"avatar_file"`
	CreateAt   time.Time `json:"create_at"`
}

type AuthorRes struct {
	Nickname   string    `json:"nickname"`
	AvatarFile string    `json:"avatar_file"`
	Email      string    `json:"email"`
	GitHub     string    `json:"github"`
	Twitter    string    `json:"twitter"`
	Bilibili   string    `json:"bilibili"`
	CreateAt   time.Time `json:"create_at"`
}
