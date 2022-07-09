package user

import (
	"context"

	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
)

type Usecase interface {
	Login(ctx context.Context, user *dto.LoginReqDto) (*dto.LoginResDto, error)
	GetUser(ctx context.Context, account string) (*model.User, error)
}
