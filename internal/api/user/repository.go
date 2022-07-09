package user

import (
	"context"
	"time"

	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
)

type Repository interface {
	GetByAccount(ctx context.Context, account string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateLoginTime(ctx context.Context, id int, updateTIme time.Time) error
}
