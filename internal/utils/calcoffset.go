package utils

func CalcOffset(page int64, pageSize int64) (int64, int64) {
	// 处理分页参数，确保它们是有效的
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || 20 < pageSize {
		pageSize = 10 // 默认每页大小
	}
	offset := (page - 1) * pageSize
	return offset, pageSize
}
