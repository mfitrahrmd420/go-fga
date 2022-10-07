package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/Calmantara/go-fga/pkg/domain/message"
	"log"

	"github.com/Calmantara/go-fga/pkg/domain/user"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func NewUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}

func (u *UserUsecaseImpl) GetUserByEmailSvc(ctx context.Context, email string) (result user.User, err error) {
	result, err = u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		log.Printf("error when getting user data with given email : %v\n", err.Error())

		return
	}

	return
}

func (u *UserUsecaseImpl) InsertUserSvc(ctx context.Context, input user.User) (result user.User, err error) {
	usr, err := u.GetUserByEmailSvc(ctx, input.Email)
	if err != nil {
		fmt.Printf("error when checking email : %v\n", err.Error())

		return result, err
	}

	if usr != (user.User{}) {
		return result, errors.New(message.EMAIL_USED)
	}

	if err = u.userRepo.InsertUser(ctx, &input); err != nil {
		log.Printf("error when inserting user : %v\n", err.Error())

		return result, err
	}

	return input, err
}

func (u *UserUsecaseImpl) GetUsersSvc(ctx context.Context) ([]user.User, error) {
	users, err := u.userRepo.GetUsers(ctx)
	if err != nil {
		log.Printf("error when getting users data : %v\n", err.Error())

		return nil, err
	}

	return users, nil
}
