package usecase

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/Hank-Kuo/personal-web-backend/config"
	"github.com/Hank-Kuo/personal-web-backend/internal/api/emoji"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type emojiUC struct {
	cfg       *config.Config
	emojiRepo emoji.Repository
	logger    logger.Logger
}

func NewUsecase(cfg *config.Config, emojiRepo emoji.Repository, log logger.Logger) emoji.Usecase {
	return &emojiUC{cfg: cfg, emojiRepo: emojiRepo, logger: log}
}

func (u *emojiUC) Update(c context.Context, blogID int, emoji *dto.EmojiDto) (*model.Emoji, error) {
	ctx, cancel := context.WithTimeout(c, u.cfg.Server.ContextTimeout)
	defer cancel()

	if err := u.emojiRepo.Update(ctx, blogID, &model.Emoji{
		Funny:    *emoji.Funny,
		Sad:      *emoji.Sad,
		Wow:      *emoji.Wow,
		Clap:     *emoji.Clap,
		Perfect:  *emoji.Perfect,
		Love:     *emoji.Love,
		Hard:     *emoji.Hard,
		Good:     *emoji.Good,
		Mad:      *emoji.Mad,
		UpdateAt: utils.GetCurrentTime(u.cfg),
	}); err != nil {
		return nil, utils.HttpError{Message: "can't update data from emoji table", Detail: err}
	}

	e, err := u.emojiRepo.GetByBlogID(ctx, blogID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.HttpError{Status: http.StatusNotFound, Message: "emoji not found", Detail: err}
		} else {
			return nil, utils.HttpError{Message: "can't fetch data from emoji table", Detail: err}
		}
	}

	return e, nil
}
