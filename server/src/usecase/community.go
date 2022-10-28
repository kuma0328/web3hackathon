package usecase

import (
	"fmt"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
)

var _ ICommunityUsecase = &CommunityUsecase{}

type CommunityUsecase struct {
	repo repository.ICommunityRepository
}

type ICommunityUsecase interface {
	GetCommunityById(id string) (*entity.Community, error)
}

func NewCommunityUsecase(repo repository.ICommunityRepository) ICommunityUsecase {
	return &CommunityUsecase{
		repo: repo,
	}
}

func (uc *CommunityUsecase) GetCommunityById(id string) (*entity.Community, error) {
	if id == "" {
		return nil, fmt.Errorf("Usecase.GetCommunityById Error: id is empty")
	}
	return uc.repo.GetCommunityById(id)
}
