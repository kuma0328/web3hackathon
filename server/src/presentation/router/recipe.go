package router

import (
	"github.com/kuma0328/web3hackathon/infrastructure/database"
	"github.com/kuma0328/web3hackathon/infrastructure/persistance"
	"github.com/kuma0328/web3hackathon/presentation/handler"
	"github.com/kuma0328/web3hackathon/usecase"
)

func (r Router) InitRecipeRouter(conn *database.Conn) {
	repoRecipe := persistance.NewRecipeRepository(conn)
	repoRecipeStep := persistance.NewRecipeStepRepository(conn)
	repoSpice := persistance.NewSpiceRepository(conn)
	ucRecipe := usecase.NewRecipeUsecase(repoRecipe,repoRecipeStep,repoSpice)
	hRecipe := handler.NewRecipeHandler(ucRecipe)

	gRecipe := r.Engine.Group("/recipe")
	gRecipe.GET("/:community_id",hRecipe.GetRecipeByCommunityId)

}
