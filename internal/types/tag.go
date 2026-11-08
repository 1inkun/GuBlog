package types

type TagData struct {
	Id    int    `json:"id"`
	Name  string `json:"tag_name"`
	Usage int    `json:"usage_count"`
}
