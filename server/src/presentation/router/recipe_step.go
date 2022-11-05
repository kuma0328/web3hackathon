package router

import (
	"github.com/kuma0328/web3hackathon/infrastructure/database"
	"github.com/kuma0328/web3hackathon/infrastructure/persistance"
	"github.com/kuma0328/web3hackathon/presentation/handler"
	"github.com/kuma0328/web3hackathon/usecase"
)

func (r Router) InitRecipeStepRouter(conn *database.Conn) {
	repoStep   := persistance.NewRecipeStepRepository(conn)
	repoRecipe := persistance.NewRecipeRepository(conn)
	ucStep     := usecase.NewRecipeStepUsecase(repoStep,repoRecipe)
	hStep      := handler.NewRecipeStepHandler(ucStep)

	gStep := r.Engine.Group("/community/:community_id/recipe")
	gStep.POST("/step",hStep.PostNewRecipeStep)

}