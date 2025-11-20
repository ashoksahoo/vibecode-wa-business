package utils

import "gorm.io/gorm"

// Pagination represents pagination parameters
type Pagination struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
	HasMore bool `json:"has_more"`
}

// PaginationResponse represents pagination in API responses
type PaginationResponse struct {
	Limit   int   `json:"limit"`
	Offset  int   `json:"offset"`
	Total   int64 `json:"total"`
	HasMore bool  `json:"has_more"`
}

// NewPagination creates a new Pagination instance with defaults
func NewPagination(limit, offset int) *Pagination {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}
	return &Pagination{
		Limit:  limit,
		Offset: offset,
	}
}

// ApplyToQuery applies pagination to a GORM query
func (p *Pagination) ApplyToQuery(db *gorm.DB) *gorm.DB {
	return db.Limit(p.Limit).Offset(p.Offset)
}

// SetTotal sets the total count and calculates HasMore
func (p *Pagination) SetTotal(total int64) {
	p.Total = total
	p.HasMore = int64(p.Offset+p.Limit) < total
}

// ToResponse converts Pagination to PaginationResponse
func (p *Pagination) ToResponse() PaginationResponse {
	return PaginationResponse{
		Limit:   p.Limit,
		Offset:  p.Offset,
		Total:   p.Total,
		HasMore: p.HasMore,
	}
}

// GetPage calculates the current page number (1-indexed)
func (p *Pagination) GetPage() int {
	if p.Limit == 0 {
		return 1
	}
	return (p.Offset / p.Limit) + 1
}

// GetTotalPages calculates the total number of pages
func (p *Pagination) GetTotalPages() int {
	if p.Limit == 0 || p.Total == 0 {
		return 0
	}
	pages := int(p.Total) / p.Limit
	if int(p.Total)%p.Limit > 0 {
		pages++
	}
	return pages
}
