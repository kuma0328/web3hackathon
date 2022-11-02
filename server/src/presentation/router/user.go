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

	g := r.Engine.Group("/user")
	g.GET("/:id", h.GetUserByID)

	loginCheckGroup := r.Engine.Group("/user", checkLogin())
	loginCheckGroup.GET("/logout", h.Logout)
	loginCheckGroup.PUT("/update", h.UpdateUser)
	loginCheckGroup.DELETE("/:id", h.DeleteUser)

	logoutCheckGroup := r.Engine.Group("/user", checkLogout())
	logoutCheckGroup.POST("/signup", h.Sigunup)
	logoutCheckGroup.POST("/login", h.Login)

}
