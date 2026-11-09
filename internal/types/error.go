package types

import "fmt"

type BusinessError struct {
	HttpCode int
	ErrCode  int64  `json:"errCode"`
	Msg      string `json:"msg"`
	Detail   string `json:"detail"`
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("错误代码 %d : %s", e.ErrCode, e.Msg)
}

func NewBusinessError(httpCode int, code int64, msg, detail string) *BusinessError {
	return &BusinessError{
		HttpCode: httpCode,
		ErrCode:  code,
		Msg:      msg,
		Detail:   detail,
	}
}

var (
	// 基本错误
	ErrBindingDataError = NewBusinessError(400, 100001, "数据绑定错误", "请求参数有误,请检查输入的数据格式")
	// 文章相关
	ErrPostNotExist = NewBusinessError(400, 200101, "找不到文章数据", "找不到对应的文章数据,请检查查询参数")
	// 作者相关
	ErrAuthorNotExist = NewBusinessError(400, 300101, "找不到作者", "找不到对应的作者数据,请检查查询参数")
)
