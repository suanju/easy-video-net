package common

//PageInfo 分页
type PageInfo struct {
	Page    int    `json:"page" form:"page"`                 // 页码
	Size    int    `json:"size" form:"size"`                 // 每页大小
	Keyword string `json:"keyword,omitempty" form:"keyword"` //关键字
}

func (p *PageInfo) Init() {
	if p.Size == 0 {
		p.Size = 20
	}
	if p.Page == 0 {
		p.Page = 1
	}
}
