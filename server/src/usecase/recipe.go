package usecase

import (
	"context"
	"fmt"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
)

type RecipeUsecase struct {
	repoRecipe repository.IRecipeRepository
	repoRecipeStep repository.IRecipeStepRepository
	repoSpice  repository.ISpiceRepository
}

type IRecipeUsecase interface {
	GetRecipeByCommunityId(ctx context.Context, communityId int) (*entity.Recipe,error)
	CreateNewRecipeOfCommunity(ctx context.Context, recipe *entity.Recipe, communityId int) (*entity.Recipe,error)
}

func NewRecipeUsecase(
	repoRecipe repository.IRecipeRepository,
	repoRecipeStep repository.IRecipeStepRepository, 
	repoSpice repository.ISpiceRepository) *RecipeUsecase {
	return &RecipeUsecase{
		repoRecipe: repoRecipe,
		repoRecipeStep: repoRecipeStep,
		repoSpice: repoSpice,
	}
}

func (uc *RecipeUsecase) GetRecipeByCommunityId(ctx context.Context, communityId int) (*entity.Recipe,error) {
	if communityId == 0{
		return nil,fmt.Errorf("Usecase.GetRecipeByCommunityId Error: community Id invalied")
	}
	recipe,err := uc.repoRecipe.GetRecipeByCommunityId(ctx,communityId)
	if err != nil {
		return nil,fmt.Errorf("Usecase.GetRecipeByCommunityId Error: %w",err)
	}
	if recipe.Id == 0 {
		return nil,fmt.Errorf("Usecase.GetRecipeByCommunityId Error: recipe not found")
	}

	recipeSteps,err := uc.repoRecipeStep.GetRecipeStepsByRecipeId(ctx,recipe.Id)
	if err != nil {
		return nil,fmt.Errorf("Usecase.GetRecipeByCommunityId GetRecipeStepsByRecipeId Error: %w",err)
	}
	recipe.RecipeSteps = recipeSteps

	spices,err := uc.repoSpice.GetSpicesByRecipeId(ctx,recipe.Id)
	if err != nil {
		return nil,fmt.Errorf("Usecase.GetRecipeByCommunityId GetSpicesByRecipeId Error: %w",err)
	}
	recipe.Spices = spices

	return recipe,nil
}

func (uc *RecipeUsecase) CreateNewRecipeOfCommunity(ctx context.Context, recipe *entity.Recipe, communityId int) (*entity.Recipe,error){
	if recipe.Name == "" {
		return nil,fmt.Errorf("Usecase.CreateNewRecipeOfCommunity Error : recipe name is empty")
	}
	if communityId == 0 {
		return nil,fmt.Errorf("Usecase.CreateNewRecipeOfCommunity Error : community Id invalied")
	}
	return uc.repoRecipe.CreateNewRecipeOfCommunity(ctx,recipe,communityId)
}
