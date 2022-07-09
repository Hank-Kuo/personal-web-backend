package tag

import (
	"context"

	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"

	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type Usecase interface {
	Fetch(ctx context.Context, p *utils.PaginationQuery) (*dto.FetchTagDto, error)
	GetByID(ctx context.Context, id int) (*dto.GetTagByIDDto, error)
	Update(ctx context.Context, tagID int, tag *model.Tag) (*model.Tag, error)
}
