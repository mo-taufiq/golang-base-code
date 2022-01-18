package pagination

import (
	"math"
)

type Pagination struct {
	PageNumber        int64 `json:"page_number"`
	TotalItemsPerPage int64 `json:"total_items_per_page"`
	TotalAllRecords   int64 `json:"total_all_records"`
	TotalPages        int64 `json:"total_pages"`
	Limit             int64 `json:"limit"`
	Offset            int64 `json:"offset"`
	PreviousPage      bool  `json:"previous_page"`
	NextPage          bool  `json:"next_page"`
}

func NewPagination(pageNumber, totalItemsPerPage, totalAllRecords int64) *Pagination {
	p := Pagination{
		PageNumber:        pageNumber,
		TotalItemsPerPage: totalItemsPerPage,
		TotalAllRecords:   totalAllRecords,
		Limit:             totalItemsPerPage,
	}
	p.countTotalPages()
	p.setOffset()
	p.setPreviousNextPage()

	return &p
}

func (p *Pagination) countTotalPages() {
	if p.TotalItemsPerPage != 0 {
		p.TotalPages = int64(math.Ceil(float64(p.TotalAllRecords) / float64(p.TotalItemsPerPage)))
	}
}

func (p *Pagination) setOffset() {
	p.Offset = (p.PageNumber - 1) * p.TotalItemsPerPage

	if p.Offset < 0 {
		p.TotalItemsPerPage = p.TotalAllRecords
		p.Limit = p.TotalAllRecords
		p.TotalPages = 1
		p.Offset = 0
	}
}

func (p *Pagination) setPreviousNextPage() {
	if p.PageNumber != 0 {
		if p.PageNumber == 1 && p.PageNumber != p.TotalPages {
			p.PreviousPage = false
			p.NextPage = true
		} else if p.PageNumber == p.TotalPages && p.PageNumber != 1 {
			p.PreviousPage = true
			p.NextPage = false
		}
	}
}
