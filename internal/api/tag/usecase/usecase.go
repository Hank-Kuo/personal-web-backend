package usecase

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/Hank-Kuo/personal-web-backend/config"
	"github.com/Hank-Kuo/personal-web-backend/internal/api/tag"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type tagUC struct {
	cfg     *config.Config
	tagRepo tag.Repository
	logger  logger.Logger
}

func NewUsecase(cfg *config.Config, tagRepo tag.Repository, log logger.Logger) tag.Usecase {
	return &tagUC{cfg: cfg, tagRepo: tagRepo, logger: log}
}

func (u *tagUC) Fetch(c context.Context, p *utils.PaginationQuery) (*dto.FetchTagDto, error) {
	ctx, cancel := context.WithTimeout(c, u.cfg.Server.ContextTimeout)
	defer cancel()

	tags, pagination, err := u.tagRepo.Fetch(ctx, p)
	if err != nil {
		return nil, utils.HttpError{Message: "can't fetch data from tag table", Detail: err}

	}

	return &dto.FetchTagDto{
		Tags: *tags,
		Meta: *pagination,
	}, nil
}

func (u *tagUC) GetByID(c context.Context, id int) (*dto.GetTagByIDDto, error) {
	ctx, cancel := context.WithTimeout(c, u.cfg.Server.ContextTimeout)
	defer cancel()

	tag, err := u.tagRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.HttpError{Status: http.StatusNotFound, Message: "not found tag id", Detail: err}
		} else {
			return nil, utils.HttpError{Message: "can't get data from tag table", Detail: err}
		}

	}

	return &dto.GetTagByIDDto{Tag: *tag}, nil
}

func (u *tagUC) Update(c context.Context, tagID int, tag *model.Tag) (*model.Tag, error) {
	ctx, cancel := context.WithTimeout(c, u.cfg.Server.ContextTimeout)
	defer cancel()

	if err := u.tagRepo.Update(ctx, tagID, tag); err != nil {
		return nil, utils.HttpError{Message: "can't update data from tag table", Detail: err}
	}

	tag, err := u.tagRepo.GetByID(ctx, tagID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.HttpError{Status: http.StatusNotFound, Message: "not found tag id", Detail: err}
		} else {
			return nil, utils.HttpError{Message: "can't get data from tag table", Detail: err}
		}
	}

	return tag, nil
}
