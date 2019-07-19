package dlldb

type IPaging interface {
}
type Paging struct {
	Offset     int `json:"offset"`     // 页码
	Limit      int `json:"limit"`      // 每页条数
	TotalPage  int `json:"totalPage"`  // 总数据条数
	TotalCount int `json:"totalCount"` // 总数据条数
}

func NewPaging() *Paging {
	return &Paging{}
}
func (p *Paging) TotalPages() int {
	if p.TotalCount == 0 || p.Limit == 0 {
		return 0
	}
	totalPage := p.TotalCount / p.Limit
	if p.TotalCount%p.Limit > 0 {
		totalPage = totalPage + 1
	}
	return totalPage
}
