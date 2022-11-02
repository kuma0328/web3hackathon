package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/infrastructure/persistance"
	"github.com/kuma0328/web3hackathon/usecase"
)

type UserHandler struct {
	uc usecase.IUserUsecase
}

func NewUserHandler(uc usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		uc: uc,
	}
}

func (u *UserHandler) Sigunup(ctx *gin.Context) {
	var json userJson
	if err := ctx.BindJSON(&json); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	user, err := u.uc.GetUserByMail(ctx, userJsonToEntity(&json).Mail)

	if user.Name != "" {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "mail already used"},
		)
		return
	}

	user, err = u.uc.CreateUser(ctx, userJsonToEntity(&json))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	cookieKey := "_cookie"
	persistance.NewSession(ctx, cookieKey, user.Id)

	userJson := userEntityToJson(user)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": userJson},
	)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var json userJson
	if err := ctx.BindJSON(&json); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	err := u.uc.LoginUser(ctx, userJsonToEntity(&json))
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	cookieKey := "_cookie"
	persistance.NewSession(ctx, cookieKey, userJsonToEntity(&json).Id)

	ctx.JSON(
		http.StatusOK,
		gin.H{"login": "ok"},
	)
}

func (u *UserHandler) Logout(ctx *gin.Context) {
	cookieKey := "_cookie"
	persistance.DeleteSession(ctx, cookieKey)

	ctx.JSON(
		http.StatusOK,
		gin.H{"delete": "ok"},
	)
}

func (u *UserHandler) GetUserByID(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	user, err := u.uc.GetUserByID(ctx, id)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	userJson := userEntityToJson(user)
	ctx.JSON(
		http.StatusOK,
		gin.H{"data": userJson},
	)

}

type userJson struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type usersJson []userJson

func userEntityToJson(c *entity.User) userJson {
	return userJson{
		Id:       c.Id,
		Name:     c.Name,
		Mail:     c.Mail,
		Password: c.Password,
	}
}

func userJsonToEntity(j *userJson) *entity.User {
	return &entity.User{
		Id:       j.Id,
		Name:     j.Name,
		Mail:     j.Mail,
		Password: j.Password,
	}
}
