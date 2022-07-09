package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/Hank-Kuo/personal-web-backend/internal/api/tag"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type tagRepo struct {
	db *sqlx.DB
}

func NewReposity(db *sqlx.DB) tag.Repository {
	return &tagRepo{db: db}
}

func (r *tagRepo) Fetch(ctx context.Context, p *utils.PaginationQuery) (*[]model.Tag, *utils.Pagination, error) {
	pagination, err := utils.GetPagination(p, "tag", r.db)
	if err != nil {
		return nil, nil, errors.Wrap(err, "tagRepo.Pagination")
	}

	tags := []model.Tag{}
	if err = r.db.SelectContext(ctx, &tags, "SELECT * FROM tag LIMIT ?, ?", p.GetOffset(), p.GetLimit()); err != nil {
		return nil, nil, errors.Wrap(err, "tagRepo.Fetch")
	}

	return &tags, pagination, nil
}

func (r *tagRepo) GetByID(ctx context.Context, tagID int) (*model.Tag, error) {
	var tag model.Tag

	err := r.db.GetContext(ctx, &tag, "SELECT * FROM tag WHERE id = ?", tagID)
	if err != nil {
		return nil, errors.Wrap(err, "tagRepo.GetByID")

	}

	return &tag, nil
}

func (r *tagRepo) GetByBlogID(ctx context.Context, blogID int) (*model.BlogToTag, error) {
	var tags []model.Tag

	err := r.db.SelectContext(ctx, &tags, `SELECT t.id, t.name, t.created_at  FROM 
											(SELECT * FROM blog_to_tag  WHERE blog_id = ?) as b_t 
											JOIN tag as t ON t.id = b_t.tag_id`, blogID)
	if err != nil {
		return nil, errors.Wrap(err, "tagRepo.GetByBlogID")
	}

	return &model.BlogToTag{ID: blogID, Tags: tags}, nil
}

func (r *tagRepo) Update(ctx context.Context, tagID int, tag *model.Tag) error {
	sqlQuery := `UPDATE tag set name=? WHERE id = ?`

	res, err := r.db.ExecContext(ctx, sqlQuery, tag.Name, tagID)
	if err != nil {
		return errors.Wrap(err, "tagRepo.Update")
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "tagRepo.Update")
	}

	if affect != 1 {
		return errors.Wrap(err, "tagRepo.Update")
	}

	return nil
}

func (r *tagRepo) QueryByName(ctx context.Context, tagName string) (*model.Tag, error) {
	var tag model.Tag
	if err := r.db.GetContext(ctx, &tag, "SELECT * FROM tag WHERE name = ?", tagName); err != nil {
		return nil, errors.Wrap(err, "tagRepo.QueryByName")
	}
	return &tag, nil
}

func (r *tagRepo) Create(ctx context.Context, tagName string) error {
	sqlQuery := `INSERT INTO tag (name) VALUES (?)`

	res, err := r.db.ExecContext(ctx, sqlQuery, tagName)
	if err != nil {
		return errors.Wrap(err, "tagRepo.Create")
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "tagRepo.Create")
	}

	if affect != 1 {
		return errors.Wrap(err, "tagRepo.Create")
	}
	return nil
}

func (r *tagRepo) CreateBlogTags(ctx context.Context, blogID int, tags *[]model.Tag) error {
	valueStrings := make([]string, 0, len(*tags))
	valueArgs := make([]interface{}, 0, len(*tags)*2)

	for _, tag := range *tags {
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, blogID)
		valueArgs = append(valueArgs, tag.ID)
	}
	stmt := fmt.Sprintf("INSERT INTO blog_to_tag (blog_id, tag_id) VALUES %s", strings.Join(valueStrings, ","))

	_, err := r.db.ExecContext(ctx, stmt, valueArgs...)
	if err != nil {
		return errors.Wrap(err, "tagRepo.CreateBlogTags")
	}

	return nil
}

func (r *tagRepo) DeleteBlogTags(ctx context.Context, blogID int, tags *[]model.Tag) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM blog_to_tag WHERE blog_id = ? ", blogID)

	if err != nil {
		return errors.Wrap(err, "tagRepo.DeleteBlogTags")
	}

	return nil
}
