package router

import (
	"github.com/kuma0328/web3hackathon/infrastructure/database"
	"github.com/kuma0328/web3hackathon/infrastructure/persistance"
	"github.com/kuma0328/web3hackathon/presentation/handler"
	"github.com/kuma0328/web3hackathon/usecase"
)

func (r Router) InitUserRouter(conn *database.Conn) {
	repo := persistance.NewUserRepository(conn)
	uc := usecase.NewUserUsecase(repo)
	h := handler.NewUserHandler(uc)

	loginCheckGroup := r.Engine.Group("/user", checkLogin())
	{
		loginCheckGroup.GET("/logout", h.Logout)
	}

	logoutCheckGroup := r.Engine.Group("/user", checkLogout())
	{
		logoutCheckGroup.POST("/signup", h.Sigunup)
		logoutCheckGroup.POST("/login", h.Login)
	}

}
