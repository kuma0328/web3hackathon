package router

import (
	"github.com/kuma0328/web3hackathon/infrastructure/database"
	"github.com/kuma0328/web3hackathon/infrastructure/persistance"
	"github.com/kuma0328/web3hackathon/presentation/handler"
	"github.com/kuma0328/web3hackathon/usecase"
)

func (r Router) InitCommunityRouter(conn *database.Conn) {
	repo := persistance.NewCommunityRepository(conn)
	uc := usecase.NewCommunityUsecase(repo)
	h := handler.NewCommunityHandler(uc)

	g := r.Engine.Group("/community")
	g.GET("/:community_id", h.GetCommunityById)
	g.PUT("/:community_id",h.UpdateCommunityOfId)
	g.DELETE("/:community_id",h.DeleteCommunityOfId)
	g.POST("/",h.CreateNewCommunity)
	g.GET("/all",h.GetCommunityAll)
}
