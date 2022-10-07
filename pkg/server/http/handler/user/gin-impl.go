package user

import (
	"log"
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

func (u *UserHdlImpl) GetUserByEmailHdl(ctx *gin.Context) {

}

// Insert New User
// @Summary this api will insert user with unique email
// @Schemes
// @Description insert new user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Router /v1/user [post]
func (u *UserHdlImpl) InsertUserHdl(ctx *gin.Context) {
	// JSON: struktur data yang bisa dibaca secara manusiawi
	// dan digunakan secara masive untuk mengirimkan payload
	// dari client -> server atau sebaliknya
	// {"first_name":"Tara", "last_name":"Calman", "email":"calman@email.com"}
	// first_name, last_name, email -> json property
	// Tara, Calman, calman@email.com -> json property value
	// yang ingin dipecahkan oleh json, standardize payload around world
	// selain json: protobuf, form, csv

	log.Printf("%T - InsertUserHdl is invoked]\n", u)
	defer log.Printf("%T - InsertUserHdl executed\n", u)

	// binding / mendapatkan body payload dari request
	log.Println("binding body payload from request")
	var user user.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Status:  "fail",
			Message: "failed to bind payload",
		})
		return
	}
	// check apakah email kosong atau tidak: kalau kosong lempar BAD_REQUEST
	log.Println("check email from request")
	if user.Email == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
			Status:  "fail",
			Message: "email should not be empty",
		})
		return
	}
	// call service/usecase untuk menginsert data
	log.Println("calling insert service usecase")
	result, err := u.userUsecase.InsertUserSvc(ctx, user)
	if err != nil {
		switch err.Error() {
		case "BAD_REQUEST":
			ctx.AbortWithStatusJSON(http.StatusBadRequest, message.Response{
				Status:  "fail",
				Message: "invalid processing payload",
			})
			return
		case "INTERNAL_SERVER_ERROR":
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, message.Response{
				Status:  "fail",
				Message: "something went wrong",
			})
			return
		}
	}
	// response result for the user if success
	ctx.JSONP(http.StatusOK, message.Response{
		Status:  "success",
		Message: "success insert user",
		Data:    result,
	})
}
