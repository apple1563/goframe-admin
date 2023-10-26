package vcommon

type CommonPageReq struct {
	PageIndex int `json:"page"`
	PageSize  int `json:"size"`
}

type CommonPageRes struct {
	PageIndex int `json:"page" d:"1" dc:"当前页数"`
	//TotalPage int `json:"total_page,omitempty" dc:"总页数"`
	PageSize int `json:"size,omitempty" dc:"页面大小"`
	Total    int `json:"total" dc:"总条数"`
}
