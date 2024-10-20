package repository

import (
	"context"
	"inspobox/webook/internal/domain"
	"inspobox/webook/internal/repository/dao"
)

var ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(d *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: d,
	}
}

func (ur *UserRepository) Create(ctx context.Context, u domain.User) error {
	err := ur.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
	return err
}
