package repository

import (
	"context"

	"github.com/kuma0328/web3hackathon/domain/entity"
)

type IRecipeRepository interface {
	GetRecipeByCommunityId(ctx context.Context, communityId int) (*entity.Recipe,error)
	CreateNewRecipeOfCommunity(ctx context.Context, recipe *entity.Recipe, communityId int) (*entity.Recipe,error)
}

