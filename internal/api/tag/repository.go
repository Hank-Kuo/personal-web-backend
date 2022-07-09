package tag

import (
	"context"

	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type Repository interface {
	Fetch(ctx context.Context, p *utils.PaginationQuery) (*[]model.Tag, *utils.Pagination, error)
	GetByID(ctx context.Context, tagID int) (*model.Tag, error)
	GetByBlogID(ctx context.Context, blogID int) (*model.BlogToTag, error)
	Update(ctx context.Context, tagID int, tag *model.Tag) error
	Create(ctx context.Context, tagName string) error
	QueryByName(ctx context.Context, tagName string) (*model.Tag, error)
	CreateBlogTags(ctx context.Context, blogID int, tags *[]model.Tag) error
	DeleteBlogTags(ctx context.Context, blogID int, tags *[]model.Tag) error
}
