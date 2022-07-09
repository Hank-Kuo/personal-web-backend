package usecase

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"golang.org/x/sync/errgroup"

	"github.com/Hank-Kuo/personal-web-backend/config"
	"github.com/Hank-Kuo/personal-web-backend/internal/api/blog"
	"github.com/Hank-Kuo/personal-web-backend/internal/api/emoji"
	"github.com/Hank-Kuo/personal-web-backend/internal/api/tag"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/dto"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/logger"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type blogUC struct {
	cfg       *config.Config
	blogRepo  blog.Repository
	tagRepo   tag.Repository
	emojiRepo emoji.Repository
	logger    logger.Logger
}

func NewUsecase(cfg *config.Config, blogRepo blog.Repository, tagRepo tag.Repository, emojiRepo emoji.Repository, log logger.Logger) blog.Usecase {
	return &blogUC{cfg: cfg, blogRepo: blogRepo, tagRepo: tagRepo, emojiRepo: emojiRepo, logger: log}
}

func (u *blogUC) fillTagsToBlogs(c context.Context, blogs *[]model.Blog) (*[]dto.BlogDto, error) {
	g, ctx := errgroup.WithContext(c)
	mapTags := map[int][]model.Tag{} //  [blog_id]-> []tag

	// Using goroutine to fetch the tag's detail
	chanTag := make(chan model.BlogToTag)
	for _, blog := range *blogs {
		id := blog.ID
		g.Go(func() error {
			tags, err := u.tagRepo.GetByBlogID(ctx, id)
			if err != nil {
				return err
			}
			chanTag <- *tags
			return nil
		})
	}

	go func() {
		if err := g.Wait(); err != nil {
			return
		}
		close(chanTag)
	}()

	for tags := range chanTag {
		mapTags[tags.ID] = tags.Tags
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	blogWithTag := []dto.BlogDto{}

	for _, blog := range *blogs {
		if tags, ok := mapTags[blog.ID]; ok {
			blogWithTag = append(blogWithTag, dto.BlogDto{Blog: blog, Tags: tags})
		}
	}

	return &blogWithTag, nil
}

func (u *blogUC) createAndGetTag(c context.Context, tags []string) (*[]model.Tag, error) {
	g, ctx := errgroup.WithContext(c)

	var detailTags []model.Tag

	// Using goroutine to fetch the tag's detail
	chanTag := make(chan model.Tag)
	for _, tag := range tags {
		t := tag
		g.Go(func() error {
			tags, err := u.tagRepo.QueryByName(ctx, t)
			if err != nil {
				err := u.tagRepo.Create(ctx, t)
				if err != nil {
					return err
				}
				tags, err := u.tagRepo.QueryByName(ctx, t)
				if err != nil {
					return err
				}
				chanTag <- *tags
				return nil
			}
			chanTag <- *tags
			return nil
		})
	}

	go func() {
		if err := g.Wait(); err != nil {
			return
		}
		close(chanTag)
	}()

	for tags := range chanTag {
		detailTags = append(detailTags, tags)
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &detailTags, nil
}

func (u *blogUC) getDiffTags(c context.Context, oldTags []model.Tag, newTags []model.Tag) (*[]model.Tag, error) {
	g, _ := errgroup.WithContext(c)
	var oldTagMap = make(map[int]bool)

	for _, tag := range oldTags {
		oldTagMap[tag.ID] = true
	}

	chanTag := make(chan model.Tag)
	for _, tag := range newTags {
		t := tag
		g.Go(func() error {
			if _, ok := oldTagMap[t.ID]; !ok {
				chanTag <- t
			}
			return nil
		})
	}

	go func() {
		if err := g.Wait(); err != nil {
			return
		}
		close(chanTag)
	}()

	var diffTags []model.Tag
	for tag := range chanTag {
		diffTags = append(diffTags, tag)
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &diffTags, nil

}

func (u *blogUC) Fetch(c context.Context, p *utils.PaginationQuery) (*dto.FetchBlogDto, error) {
	ctx, cancel := context.WithTimeout(c, u.cfg.Server.ContextTimeout)
	defer cancel()

	blogs, pagination, err := u.blogRepo.Fetch(ctx, p)
	if err != nil {
		return nil, utils.HttpError{Message: "can't fetch data from blog table", Detail: err}
	}

	blogWithTag, err := u.fillTagsToBlogs(ctx, blogs)
	if err != nil {
		return nil, utils.HttpError{Message: "can't merge tag into blogs", Detail: err}
	}

	return &dto.FetchBlogDto{
		Blogs: *blogWithTag,
		Meta:  *pagination,
	}, nil
}

func (u *blogUC) GetByID(c context.Context, id int) (*dto.GetBlogByIDDto, error) {
	ctx, cancel := context.WithTimeout(c, u.cfg.Server.ContextTimeout)
	defer cancel()

	var blogDetail dto.GetBlogByIDDto

	blog, err := u.blogRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.HttpError{Status: http.StatusNotFound, Message: "not found blog", Detail: err}
		} else {
			return nil, utils.HttpError{Message: "can't get data from blog table", Detail: err}
		}
	}

	tags, err := u.tagRepo.GetByBlogID(ctx, blog.ID)
	if err != nil {
		return nil, utils.HttpError{Message: "can't get tags from tag repo", Detail: err}
	}

	emoji, err := u.emojiRepo.GetByBlogID(ctx, blog.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.HttpError{Status: http.StatusNotFound, Message: "not found emoji", Detail: err}
		} else {
			return nil, utils.HttpError{Message: "can't get emoji from emoji repo", Detail: err}
		}
	}

	blogDetail.Blog = *blog
	blogDetail.Tags = tags.Tags
	blogDetail.Emoji = *emoji

	return &blogDetail, nil
}

func (u *blogUC) Create(c context.Context, blog *dto.CreateBlogDto) error {
	ctx, cancel := context.WithTimeout(c, u.cfg.Server.ContextTimeout)
	defer cancel()

	id, err := u.blogRepo.Create(ctx, &model.Blog{Title: blog.Title, Link: blog.Link, ImgLink: blog.ImgLink})
	if err != nil {
		return utils.HttpError{Message: "can't create blog from blog repo", Detail: err}
	}

	if err := u.emojiRepo.Create(ctx, id); err != nil {
		return utils.HttpError{Message: "can't create emoji from emoji repo", Detail: err}
	}

	tags, err := u.createAndGetTag(ctx, blog.Tags)
	if err != nil {
		return utils.HttpError{Message: "can't get & create tag id", Detail: err}
	}

	if err := u.tagRepo.CreateBlogTags(ctx, id, tags); err != nil {
		return utils.HttpError{Message: "can't create blog to tag", Detail: err}
	}

	return nil
}

func (u *blogUC) Update(c context.Context, blogID int, blog *dto.UpdateBlogDto) error {
	ctx, cancel := context.WithTimeout(c, u.cfg.Server.ContextTimeout)
	defer cancel()

	if err := u.blogRepo.Update(ctx, blogID, &model.Blog{Title: blog.Title, Link: blog.Link,
		ImgLink: blog.ImgLink, Visitor: blog.Visitor, UpdatedAt: utils.GetCurrentTime(u.cfg)}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HttpError{Status: http.StatusNotFound, Message: "not found blog", Detail: err}
		} else {
			return utils.HttpError{Message: "can't update blog from blog repo", Detail: err}
		}
	}

	if len(blog.Tags) != 0 {

		tags, err := u.createAndGetTag(ctx, blog.Tags)
		if err != nil {
			return utils.HttpError{Message: "can't get & create tag id", Detail: err}
		}
		if err := u.tagRepo.DeleteBlogTags(ctx, blogID, tags); err != nil {
			return utils.HttpError{Message: "can't delete tags", Detail: err}
		}

		if err := u.tagRepo.CreateBlogTags(ctx, blogID, tags); err != nil {
			return utils.HttpError{Message: "can't create tags", Detail: err}
		}

	}

	return nil
}
