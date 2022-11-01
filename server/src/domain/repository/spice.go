package repository

import (
	"context"

	"github.com/kuma0328/web3hackathon/domain/entity"
)

type ISpiceRepository interface {
	GetSpicesByRecipeId(ctx context.Context, recipeId int) (entity.Spices,error)
	PostNewSpice(ctx context.Context,spice *entity.Spice,recipeId int) (*entity.Spice,error)
}
