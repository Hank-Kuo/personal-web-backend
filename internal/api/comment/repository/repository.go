package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/Hank-Kuo/personal-web-backend/internal/api/comment"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type commentRepo struct {
	db *sqlx.DB
}

func NewReposity(db *sqlx.DB) comment.Repository {
	return &commentRepo{db: db}
}

func (r *commentRepo) GetByID(ctx context.Context, id int) (*model.Comment, error) {
	var comment model.Comment
	if err := r.db.GetContext(ctx, &comment, "SELECT * FROM comment WHERE id = ?", id); err != nil {
		return nil, errors.Wrap(err, "commentRepo.GetByBlogID")
	}
	return &comment, nil
}

func (r *commentRepo) GetByBlogID(ctx context.Context, blogID int, p *utils.PaginationQuery) (*[]model.Comment, *utils.Pagination, error) {
	pagination, err := utils.GetPagination(p, "comment", r.db)
	if err != nil {
		return nil, nil, errors.Wrap(err, "commentRepo.GetByBlogID")
	}

	comment := []model.Comment{}
	if err = r.db.SelectContext(ctx, &comment, "SELECT * FROM comment WHERE blog_id = ? LIMIT ?, ?", blogID, p.GetOffset(), p.GetLimit()); err != nil {
		return nil, nil, errors.Wrap(err, "commentRepo.GetByBlogID")
	}

	return &comment, pagination, nil
}

func (r *commentRepo) Create(ctx context.Context, comment *model.Comment) (int, error) {
	sqlQuery := `INSERT INTO comment(blog_id, name, character, comment) VALUES
				(?, ?, ?, ?)`

	res, err := r.db.ExecContext(ctx, sqlQuery, comment.BlogID, comment.Name, comment.Character, comment.Comment)
	if err != nil {
		return -1, errors.Wrap(err, "commentRepo.Create")
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return -1, errors.Wrap(err, "commentRepo.Create")
	}

	if affect != 1 {
		return -1, errors.Wrap(err, "commentRepo.Create")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return -1, errors.Wrap(err, "commentRepo.Create")
	}
	return int(id), nil
}
