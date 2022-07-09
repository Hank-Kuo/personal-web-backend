package usecase

import (
	"context"
	"net/http"

	"github.com/Hank-Kuo/personal-web-backend/config"
	"github.com/Hank-Kuo/personal-web-backend/internal/api/blog"
	"github.com/Hank-Kuo/personal-web-backend/internal/api/comment"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type commentUC struct {
	cfg         *config.Config
	commentRepo comment.Repository
	blogRepo    blog.Repository
	logger      logger.Logger
}

func NewUsecase(cfg *config.Config, commentRepo comment.Repository, blogRepo blog.Repository, log logger.Logger) comment.Usecase {
	return &commentUC{cfg: cfg, commentRepo: commentRepo, blogRepo: blogRepo, logger: log}
}

func (u *commentUC) GetByBlogID(c context.Context, blogID int, p *utils.PaginationQuery) (*dto.GetCommentByBlogIDDto, error) {
	ctx, cancel := context.WithTimeout(c, u.cfg.Server.ContextTimeout)
	defer cancel()

	comment, pagination, err := u.commentRepo.GetByBlogID(ctx, blogID, p)
	if err != nil {
		return nil, utils.HttpError{Message: "can't fetch data from comment table", Detail: err}
	}

	return &dto.GetCommentByBlogIDDto{
		Comment: *comment,
		Meta:    *pagination,
	}, nil
}

func (u *commentUC) Create(c context.Context, comment *dto.CraeteCommentDto) (*model.Comment, error) {
	ctx, cancel := context.WithTimeout(c, u.cfg.Server.ContextTimeout)
	defer cancel()

	if _, err := u.blogRepo.GetByID(ctx, comment.BlogID); err != nil {
		return nil, utils.HttpError{Status: http.StatusNotFound, Message: "not found blog, can't create", Detail: err}
	}

	id, err := u.commentRepo.Create(ctx, &model.Comment{
		BlogID:    comment.BlogID,
		Name:      comment.Name,
		Comment:   comment.Comment,
		Character: utils.GetCharacter(),
	})
	if err != nil {
		return nil, utils.HttpError{Message: "can't create comment", Detail: err}
	}

	newComment, err := u.commentRepo.GetByID(ctx, id)

	if err != nil {
		return nil, utils.HttpError{Message: "can't get comment", Detail: err}
	}

	return newComment, nil
}
