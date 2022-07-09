package utils

import (
	"math"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const defaultLimit = 10

type PaginationQuery struct {
	Limit   int    `form:"limit"`
	Page    int    `form:"page"`
	OrderBy string `form:"order_by"`
}

type Pagination struct {
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	OrderBy   string `json:"order_by"`
	TotalPage int    `json:"total_page"`
}

func GetPagination(p *PaginationQuery, table string, db *sqlx.DB) (*Pagination, error) {
	var totalCount int

	db.QueryRow("SELECT count(*) FROM " + table).Scan(&totalCount)
	totalPages := int(math.Ceil(float64(totalCount) / float64(p.GetLimit())))

	return &Pagination{
		Limit:     p.GetLimit(),
		Page:      p.GetPage(),
		OrderBy:   p.GetOrderBy(),
		TotalPage: totalPages,
	}, nil
}

func (p *PaginationQuery) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *PaginationQuery) GetOrderBy() string {
	if p.OrderBy == "" {
		p.OrderBy = "id"
	}
	return p.OrderBy
}

func (p *PaginationQuery) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = defaultLimit
	}
	return p.Limit
}

func (p *PaginationQuery) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func GetPaginationFromGin(c *gin.Context) (*PaginationQuery, error) {
	var p PaginationQuery
	if err := c.BindQuery(&p); err != nil {
		return nil, err
	}

	return &p, nil
}
