package comment

import (
	"context"

	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"

	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type Usecase interface {
	GetByBlogID(ctx context.Context, blogID int, p *utils.PaginationQuery) (*dto.GetCommentByBlogIDDto, error)
	Create(ctx context.Context, comment *dto.CraeteCommentDto) (*model.Comment, error)
}
