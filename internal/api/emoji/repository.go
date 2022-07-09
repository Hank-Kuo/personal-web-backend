package emoji

import (
	"context"

	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
)

type Repository interface {
	GetByBlogID(ctx context.Context, blogID int) (*model.Emoji, error)
	Update(ctx context.Context, blogID int, emoji *model.Emoji) error
	Create(ctx context.Context, blogID int) error
}
