package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/Hank-Kuo/personal-web-backend/internal/api/emoji"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
)

type emojiRepo struct {
	db *sqlx.DB
}

func NewReposity(db *sqlx.DB) emoji.Repository {
	return &emojiRepo{db: db}
}
func (r *emojiRepo) GetByBlogID(ctx context.Context, blogID int) (*model.Emoji, error) {
	var emoji model.Emoji
	if err := r.db.GetContext(ctx, &emoji, "SELECT * FROM emoji WHERE blog_id = ? ", blogID); err != nil {
		return nil, errors.Wrap(err, "emojiRepo.GetByBlogID")
	}

	return &emoji, nil
}

func (r *emojiRepo) Update(ctx context.Context, blogID int, emoji *model.Emoji) error {
	sqlQuery := `UPDATE emoji set funny=?, sad=?, wow=?, 
					clap=?, perfect=?, good=?, love=?, hard=?, 
					mad=?, updated_at=? WHERE blog_id = ?`

	res, err := r.db.ExecContext(ctx, sqlQuery, emoji.Funny, emoji.Sad, emoji.Wow,
		emoji.Clap, emoji.Perfect, emoji.Good, emoji.Love,
		emoji.Hard, emoji.Mad, emoji.UpdateAt, blogID)

	if err != nil {
		return errors.Wrap(err, "emojiRepo.Update")
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "emojiRepo.Update")
	}

	if affect != 1 {
		return errors.Wrap(err, "emojiRepo.Update")
	}

	return nil
}

func (r *emojiRepo) Create(ctx context.Context, blogID int) error {
	sqlQuery := `INSERT INTO emoji(blog_id) VALUES (?)`

	res, err := r.db.ExecContext(ctx, sqlQuery, blogID)

	if err != nil {
		return errors.Wrap(err, "emojiRepo.Create")
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "emojiRepo.Create")
	}

	if affect != 1 {
		return errors.Wrap(err, "emojiRepo.Create")
	}

	return nil
}
