package user

import (
	"context"
	"log"

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
	log.Printf("%T - GetUserByEmail is invoked]\n", u)
	defer log.Printf("%T - GetUserByEmail executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()
	// insert new user
	db.Model(&user.User{}).
		Where("email = ?", email).
		Find(&result)
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when getting user with email %v\n",
			email)
	}
	return result, err
}

func (u *UserRepoImpl) InsertUser(ctx context.Context, insertedUser *user.User) (err error) {
	log.Printf("%T - InsertUser is invoked]\n", u)
	defer log.Printf("%T - InsertUser executed\n", u)
	// get gorm client first
	db := u.pgCln.GetClient()
	// insert new user
	db.Model(&user.User{}).
		Create(&insertedUser)
	//check error
	if err = db.Error; err != nil {
		log.Printf("error when inserting user with email %v\n",
			insertedUser.Email)
	}
	return err
}

func (u *UserRepoImpl) GetUsers(ctx context.Context) ([]user.User, error) {
	var result []user.User

	if err := u.pgCln.GetClient().Model(&user.User{}).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
