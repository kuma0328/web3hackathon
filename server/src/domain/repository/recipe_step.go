package repository

import (
	"context"

	"github.com/kuma0328/web3hackathon/domain/entity"
)

type IRecipeStepRepository interface {
	GetRecipeStepsByRecipeId(ctx context.Context, recipeId int) (entity.RecipeSteps, error)
}
