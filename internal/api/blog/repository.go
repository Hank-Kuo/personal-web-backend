package blog

import (
	"context"

	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type Repository interface {
	Fetch(ctx context.Context, p *utils.PaginationQuery) (*[]model.Blog, *utils.Pagination, error)
	GetByID(ctx context.Context, id int) (*model.Blog, error)
	Create(ctx context.Context, blog *model.Blog) (int, error)
	Update(ctx context.Context, id int, blog *model.Blog) error
}
