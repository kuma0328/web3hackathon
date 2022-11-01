package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kuma0328/web3hackathon/domain/entity"
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
	newName := ctx.Param("name")
	newMail := ctx.Param("mail")
	newMail := ctx.Param("password")

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
