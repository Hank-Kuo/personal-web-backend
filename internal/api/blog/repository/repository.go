package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/Hank-Kuo/personal-web-backend/internal/api/blog"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type blogRepo struct {
	db *sqlx.DB
}

func NewReposity(db *sqlx.DB) blog.Repository {
	return &blogRepo{db: db}
}

func (r *blogRepo) Create(ctx context.Context, blog *model.Blog) (int, error) {
	sqlQuery := "INSERT INTO blog(title, link, img_link) VALUES (?, ?, ?)"

	res, err := r.db.ExecContext(ctx, sqlQuery, blog.Title, blog.Link, blog.ImgLink)

	if err != nil {
		return -1, errors.Wrap(err, "blogRepo.Create")
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return -1, errors.Wrap(err, "blogRepo.Create")
	}

	if affect != 1 {
		return -1, errors.Wrap(err, "blogRepo.Create")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, errors.Wrap(err, "commentRepo.Create")
	}
	return int(id), nil

}

func (r *blogRepo) Fetch(ctx context.Context, p *utils.PaginationQuery) (*[]model.Blog, *utils.Pagination, error) {
	pagination, err := utils.GetPagination(p, "blog", r.db)
	if err != nil {
		return nil, nil, errors.Wrap(err, "blogRepo.Pagination")
	}

	blogs := []model.Blog{}
	if err = r.db.SelectContext(ctx, &blogs, "SELECT * FROM blog ORDER BY "+p.GetOrderBy()+" DESC LIMIT ?, ?", p.GetOffset(), p.GetLimit()); err != nil {
		return nil, nil, errors.Wrap(err, "blogRepo.Fetch")
	}
	return &blogs, pagination, nil
}

func (r *blogRepo) GetByID(ctx context.Context, id int) (*model.Blog, error) {
	var blog model.Blog

	err := r.db.GetContext(ctx, &blog, "SELECT * FROM blog WHERE id = ?", id)
	if err != nil {
		return nil, errors.Wrap(err, "blogRepo.GetByID")
	}

	return &blog, nil
}

func (r *blogRepo) Pagination(ctx context.Context, p *utils.PaginationQuery) (*utils.Pagination, error) {
	pagination, err := utils.GetPagination(p, "blog", r.db)
	if err != nil {
		return nil, errors.Wrap(err, "blogRepo.Pagination")
	}
	return pagination, nil
}

func (r *blogRepo) Update(ctx context.Context, blogID int, blog *model.Blog) error {
	var valueArgs []interface{}
	sqlArsgs := []string{}

	if blog.Title != "" {
		sqlArsgs = append(sqlArsgs, "title=?")
		valueArgs = append(valueArgs, blog.Title)
	}
	if blog.Link != "" {
		sqlArsgs = append(sqlArsgs, "link=?")
		valueArgs = append(valueArgs, blog.Link)
	}
	if blog.ImgLink != "" {
		sqlArsgs = append(sqlArsgs, "img_link=?")
		valueArgs = append(valueArgs, blog.ImgLink)
	}
	if blog.Visitor != 0 {
		sqlArsgs = append(sqlArsgs, "visitor=?")
		valueArgs = append(valueArgs, blog.Visitor)
	}
	sqlArsgs = append(sqlArsgs, "updated_at=?")
	valueArgs = append(valueArgs, blog.UpdatedAt)
	valueArgs = append(valueArgs, blogID)

	sqlQuery := fmt.Sprintf("UPDATE blog set %s WHERE id =?", strings.Join(sqlArsgs, ","))

	res, err := r.db.ExecContext(ctx, sqlQuery, valueArgs...)

	if err != nil {
		return errors.Wrap(err, "blogRepo.Update")
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "blogRepo.Update")
	}

	if affect != 1 {
		return errors.Wrap(err, "blogRepo.Update")
	}

	return nil
}
