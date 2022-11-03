package usecase

import (
	"context"
	"fmt"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
)

var _ ISpiceUsecase = &SpiceUsecase{}

type SpiceUsecase struct {
	repoSpice repository.ISpiceRepository
	repoRecipe repository.IRecipeRepository
}

type ISpiceUsecase interface {
	PostNewSpice(ctx context.Context, spice *entity.Spice, communityId int) (*entity.Spice,error)
}

func NewSpiceUsecase(repoSpice repository.ISpiceRepository,repoRecipe repository.IRecipeRepository) ISpiceUsecase {
	return &SpiceUsecase{
		repoSpice: repoSpice,
		repoRecipe: repoRecipe,
	}
}

func (uc *SpiceUsecase) PostNewSpice(ctx context.Context, spice *entity.Spice, communityId int) (*entity.Spice,error) {
	if spice.Name == ""{
		return nil,fmt.Errorf("SpiceUsecase.PostNewSpice Error: Name is empty")
	}
	recipe,err := uc.repoRecipe.GetRecipeByCommunityId(ctx,communityId)
	if err != nil {
		return nil,fmt.Errorf("SpiceUsecase.PostNewSpice Error : %w",err)
	}
	return uc.repoSpice.PostNewSpice(ctx,spice,recipe.Id)
}

