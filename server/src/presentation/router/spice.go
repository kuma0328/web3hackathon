package router

import (
	"github.com/kuma0328/web3hackathon/infrastructure/database"
	"github.com/kuma0328/web3hackathon/infrastructure/persistance"
	"github.com/kuma0328/web3hackathon/presentation/handler"
	"github.com/kuma0328/web3hackathon/usecase"
)

// TODO r *Routerのほうが良いのでは?
func (r Router) InitSpiceRouter(conn *database.Conn){
	repoSpice  := persistance.NewSpiceRepository(conn)
	repoRecipe := persistance.NewRecipeRepository(conn)
	ucSpice    := usecase.NewSpiceUsecase(repoSpice,repoRecipe)
	hSpice     := handler.NewSpiceHandler(ucSpice)

	gSpice := r.Engine.Group("/community/:community_id/recipe")
	gSpice.POST("/spice",hSpice.PostNewSpice)
}