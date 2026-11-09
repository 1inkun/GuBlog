package types

type Pagination struct {
	Page     int64 `json:"page" query:"page"`
	PageSize int64 `json:"page_size" query:"page_size"`
}

type GetPostDataReq struct {
	Tags       []int64 `query:"tag"`
	Pagination Pagination
}
