package vcommon

type CommonPageReq struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type CommonPageRes struct {
	Page int `json:"page" d:"1" dc:"当前页数"`
	//TotalPage int `json:"total_page,omitempty" dc:"总页数"`
	Size  int `json:"size,omitempty" dc:"页面大小"`
	Total int `json:"total" dc:"总条数"`
}
