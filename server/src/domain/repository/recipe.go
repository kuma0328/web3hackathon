package repository

import (
	"context"

	"github.com/kuma0328/web3hackathon/domain/entity"
)

type IRecipeRepository interface {
	GetRecipeByCommunityId(ctx context.Context, communityId int) (*entity.Recipe,error)
}

