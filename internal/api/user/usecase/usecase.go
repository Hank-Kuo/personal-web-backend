package usecase

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"regexp"

	"github.com/Hank-Kuo/personal-web-backend/config"
	"github.com/Hank-Kuo/personal-web-backend/internal/api/user"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type userUC struct {
	cfg      *config.Config
	userRepo user.Repository
	logger   logger.Logger
}

func NewUsecase(cfg *config.Config, userRepo user.Repository, log logger.Logger) user.Usecase {
	return &userUC{cfg: cfg, userRepo: userRepo, logger: log}
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func (u *userUC) GetUser(ctx context.Context, account string) (*model.User, error) {
	if isEmailValid(account) {
		user, err := u.userRepo.GetByEmail(ctx, account)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, utils.HttpError{Status: http.StatusNotFound, Message: "not found user", Detail: err}
			} else {
				return nil, utils.HttpError{Message: "can't get user from user repo", Detail: err}
			}
		}
		return user, nil
	} else {
		user, err := u.userRepo.GetByAccount(ctx, account)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, utils.HttpError{Status: http.StatusNotFound, Message: "not found user", Detail: err}
			} else {
				return nil, utils.HttpError{Message: "can't get user from user repo", Detail: err}
			}
		}
		return user, nil
	}
}

func (u *userUC) Login(c context.Context, body *dto.LoginReqDto) (*dto.LoginResDto, error) {
	ctx, cancel := context.WithTimeout(c, u.cfg.Server.ContextTimeout)
	defer cancel()

	user, err := u.GetUser(ctx, body.Account)
	if err != nil {
		return nil, err
	}

	if err := utils.CheckPasswordHash(body.Password, user.Password); err != nil {
		return nil, utils.HttpError{Message: "incorrect password", Detail: err}
	}

	if err := u.userRepo.UpdateLoginTime(ctx, user.ID, utils.GetCurrentTime(u.cfg)); err != nil {
		return nil, utils.HttpError{Message: "can't update login time", Detail: err}
	}

	accessJWT, err := utils.GetJwt(u.cfg, user, "access")
	if err != nil {
		return nil, utils.HttpError{Message: "generate access token error", Detail: err}
	}

	refreshJWT, err := utils.GetJwt(u.cfg, user, "refresh")
	if err != nil {
		return nil, utils.HttpError{Message: "generate refresh token error", Detail: err}
	}

	return &dto.LoginResDto{
		User:         *user,
		AccessToken:  accessJWT,
		Refreshtoken: refreshJWT,
		ExpiredIn:    u.cfg.Server.JwtExpireTime,
		TokenType:    "bear",
	}, nil
}
