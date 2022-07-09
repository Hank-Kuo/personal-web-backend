package blog

import (
	"context"

	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"

	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type Usecase interface {
	Fetch(ctx context.Context, p *utils.PaginationQuery) (*dto.FetchBlogDto, error)
	GetByID(ctx context.Context, id int) (*dto.GetBlogByIDDto, error)
	Create(ctx context.Context, blog *dto.CreateBlogDto) error
	Update(ctx context.Context, id int, blog *dto.UpdateBlogDto) error
}
