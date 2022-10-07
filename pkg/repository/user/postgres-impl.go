package user

import (
	"context"
	"github.com/Calmantara/go-fga/config/postgres"
	"github.com/Calmantara/go-fga/pkg/domain/user"
)

type UserRepoImpl struct {
	pgCln postgres.PostgresClient
}

func NewUserRepo(pgCln postgres.PostgresClient) user.UserRepo {
	return &UserRepoImpl{pgCln: pgCln}
}

func (u *UserRepoImpl) GetUserByEmail(ctx context.Context, email string) (result user.User, err error) {
	err = u.pgCln.GetClient().
		Model(&user.User{}).
		Where("email = ?", email).
		Find(&result).Error

	return
}

func (u *UserRepoImpl) InsertUser(ctx context.Context, insertedUser *user.User) (err error) {
	err = u.pgCln.GetClient().
		Model(&user.User{}).
		Create(&insertedUser).Error

	return
}

func (u *UserRepoImpl) GetUsers(ctx context.Context) (result []user.User, err error) {
	err = u.pgCln.GetClient().
		Model(&user.User{}).
		Find(&result).Error

	return
}
