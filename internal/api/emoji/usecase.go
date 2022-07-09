package emoji

import (
	"context"

	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
)

type Usecase interface {
	Update(ctx context.Context, blogID int, emoji *dto.EmojiDto) (*model.Emoji, error)
}
