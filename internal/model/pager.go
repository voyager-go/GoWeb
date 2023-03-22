package model

type Pagination struct {
	PageNum    int         `json:"pageNum"`    // 当前页码
	PageSize   int         `json:"pageSize"`   // 每页数量
	Total      int64       `json:"total"`      // 总记录数
	TotalPages int         `json:"totalPages"` // 总页数
	List       interface{} `json:"list"`       // 分页数据列表
}
