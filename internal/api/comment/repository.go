package comment

import (
	"context"

	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type Repository interface {
	GetByID(ctx context.Context, id int) (*model.Comment, error)
	GetByBlogID(ctx context.Context, blogID int, p *utils.PaginationQuery) (*[]model.Comment, *utils.Pagination, error)
	Create(ctx context.Context, comment *model.Comment) (int, error)
}
