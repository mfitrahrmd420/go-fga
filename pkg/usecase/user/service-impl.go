package user

import (
	"context"
	"errors"
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
	log.Printf("%T - GetUserByEmail is invoked]\n", u)
	defer log.Printf("%T - GetUserByEmail executed\n", u)
	// get user from repository (database)
	log.Println("getting user from user repository")
	result, err = u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		// ini berarti ada yang salah dengan connection di database
		log.Fatal("error when fetching data from database: " + err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")
		return result, err
	}
	// check user id > 0 ?
	log.Println("checking user id")
	if result.ID <= 0 {
		// kalau tidak berarti user not found
		log.Fatal("user is not found: " + email)
		err = errors.New("NOT_FOUND")
		return result, err
	}
	return result, err
}

func (u *UserUsecaseImpl) InsertUserSvc(ctx context.Context, input user.User) (result user.User, err error) {
	log.Printf("%T - InsertUserSvc is invoked]\n", u)
	defer log.Printf("%T - InsertUserSvc executed\n", u)
	// get user for input email first
	usrCheck, err := u.GetUserByEmailSvc(ctx, input.Email)

	// check user is exist or not
	if err == nil {
		// user found
		log.Fatalf("user has been registered with id: %v", usrCheck.ID)
		err = errors.New("BAD_REQUEST")
		return result, err
	}
	// internal server error condition
	if err.Error() != "NOT_FOUND" {
		// internal server error
		log.Fatal("got error when checking user from database")
		return result, err
	}
	// valid condition: NOT_FOUND
	log.Println("insert user to database process")
	if err = u.userRepo.InsertUser(ctx, input); err != nil {
		log.Fatalf("error when inserting user:%v", err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")
	}
	return result, err
}
