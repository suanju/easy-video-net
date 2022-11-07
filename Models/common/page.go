package common

//PageInfo 分页
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

func (p *PageInfo) Init() {
	if p.PageSize == 0 {
		p.PageSize = 20
	}
	if p.Page == 0 {
		p.Page = 1
	}
}
