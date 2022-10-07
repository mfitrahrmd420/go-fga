package user

import (
	"fmt"
	"net/http"

	"github.com/Calmantara/go-fga/pkg/domain/message"
	"github.com/Calmantara/go-fga/pkg/domain/user"
	"github.com/gin-gonic/gin"
)

type UserHdlImpl struct {
	userUsecase user.UserUsecase
}

func NewUserHandler(userUsecase user.UserUsecase) user.UserHandler {
	return &UserHdlImpl{userUsecase: userUsecase}
}

// GetUserByEmailHdl godoc
// @Summary get user by email
// @Description this api will get a user with specific email
// @Tags users
// @Accept json
// @Produce json
// @Param email path string true "user email"
// @Success 200 {object} user.User
// @Router /v1/users [get]
func (u *UserHdlImpl) GetUserByEmailHdl(ctx *gin.Context) {
	email := ctx.Param("email")

	usr, err := u.userUsecase.GetUserByEmailSvc(ctx, email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, message.Response{
			Status:  "fail",
			Message: "something went wrong",
		})

		return
	}

	if usr == (user.User{}) {
		ctx.JSON(http.StatusNotFound, message.Response{
			Status:  "fail",
			Message: fmt.Sprintf("user with email %s was not found", email),
		})

		return
	}

	ctx.JSON(http.StatusOK, message.Response{
		Status: "success",
		Data:   usr,
	})
}

// InsertUserHdl godoc
// @Summary insert new user
// @Description this api will insert user with unique email
// @Tags users
// @Accept json
// @Produce json
// @Param user body user.User true "user object body"
// @Success 201 {object} user.User
// @Router /v1/users [post]
func (u *UserHdlImpl) InsertUserHdl(ctx *gin.Context) {
	var newUser user.User

	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Status:  "fail",
			Message: "failed to bind payload",
		})

		return
	}

	if newUser.Email == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Status:  "fail",
			Message: "email should not be empty",
		})

		return
	}

	result, err := u.userUsecase.InsertUserSvc(ctx, newUser)
	if err != nil {
		switch err.Error() {
		case message.EMAIL_USED:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Status:  "fail",
				Message: fmt.Sprintf("email '%s' already used", newUser.Email),
			})

			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Status:  "fail",
				Message: "something went wrong",
			})

			return
		}
	}

	ctx.JSON(http.StatusCreated, message.Response{
		Status: "success",
		Data:   result,
	})
}

func (u *UserHdlImpl) GetUsersHdl(ctx *gin.Context) {
	users, err := u.userUsecase.GetUsersSvc(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
			Status:  "fail",
			Message: "something went wrong",
		})

		return
	}

	ctx.JSON(http.StatusOK, message.Response{
		Status: "success",
		Data:   users,
	})
}
