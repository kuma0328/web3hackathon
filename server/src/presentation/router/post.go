package router

import (
	"github.com/kuma0328/web3hackathon/infrastructure/database"
	"github.com/kuma0328/web3hackathon/infrastructure/persistance"
	"github.com/kuma0328/web3hackathon/presentation/handler"
	"github.com/kuma0328/web3hackathon/usecase"
)

func (r Router) InitPostRouter(conn *database.Conn) {
	repo := persistance.NewPostRepository(conn)
	uc := usecase.NewPostUsecase(repo)
	h := handler.NewPostHandler(uc)

	g := r.Engine.Group("/post")
	g.GET("/data/:id", h.GetDataByPostById)
	g.GET("/img/:id", h.GetImgByPostById)
	g.PUT("/update", h.UpdatePostOfId)
	g.DELETE("/:id", h.DeletePostOfId)
	g.POST("/create", h.CreateNewPost)
}
