package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/Hank-Kuo/personal-web-backend/internal/api/user"
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
)

type userRepo struct {
	db *sqlx.DB
}

func NewReposity(db *sqlx.DB) user.Repository {
	return &userRepo{db: db}
}

func (r *userRepo) UpdateLoginTime(ctx context.Context, id int, updateTime time.Time) error {
	sqlQuery := `UPDATE user set login_time=? WHERE id = ?`

	res, err := r.db.ExecContext(ctx, sqlQuery, updateTime, id)

	if err != nil {
		return errors.Wrap(err, "userRepo.UpdateLoginTime")
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "userRepo.UpdateLoginTime")
	}

	if affect != 1 {
		return errors.Wrap(err, "userRepo.UpdateLoginTime")
	}

	return nil
}
func (r *userRepo) GetByAccount(ctx context.Context, account string) (*model.User, error) {
	var user model.User
	if err := r.db.GetContext(ctx, &user, "SELECT * FROM user WHERE account = ? ", account); err != nil {
		return nil, errors.Wrap(err, "userRepo.GetByAccount")
	}
	return &user, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.db.GetContext(ctx, &user, "SELECT * FROM user WHERE email = ? ", email); err != nil {
		return nil, errors.Wrap(err, "userRepo.GetByEmail")
	}

	return &user, nil
}
