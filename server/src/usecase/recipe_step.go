package usecase

import (
	"context"
	"fmt"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
)

type RecipeStepUsecase struct {
	repoStep repository.IRecipeStepRepository
	repoRecipe repository.IRecipeRepository
}

type IRecipeStepUsecase interface {
	PostNewRecipeStep(ctx context.Context,step *entity.RecipeStep,communityId int) (*entity.RecipeStep,error)
}

func NewRecipeStepUsecase(repoStep repository.IRecipeStepRepository,repoRecipe repository.IRecipeRepository) IRecipeStepUsecase {
	return &RecipeStepUsecase{
		repoStep:repoStep,
		repoRecipe: repoRecipe,
	}
}

func (uc *RecipeStepUsecase) PostNewRecipeStep(ctx context.Context,step *entity.RecipeStep,communityId int) (*entity.RecipeStep,error) {
	if step.Number == 0 {
		return nil,fmt.Errorf("Usecase.GetRecipeByCommunityId Error: Number is empty or 0")
	}
	if step.Content == "" {
		return nil,fmt.Errorf("Usecase.GetRecipeByCommunityId Error: Content is empty")
	}

	recipe,err := uc.repoRecipe.GetRecipeByCommunityId(ctx,communityId)
	if err != nil {
		return nil,fmt.Errorf("RecipeStepUsecase.PostNewRecipeStep Error : %w",err)
	}
	return uc.repoStep.PostNewRecipeStep(ctx,step,recipe.Id)
}

