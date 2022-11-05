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
	GetCommunityById(id int) (*entity.Community, error)
	UpdateCommunityOfId(id int, community *entity.Community)(*entity.Community,error)
	DeleteCommunityOfId(id int)error
	CreateNewCommunity(community *entity.Community)(*entity.Community,error)
	GetCommunityAll()(entity.Communities, error)
}

func NewCommunityUsecase(repo repository.ICommunityRepository) ICommunityUsecase {
	return &CommunityUsecase{
		repo: repo,
	}
}

func (uc *CommunityUsecase) GetCommunityAll() (entity.Communities, error) {
	return uc.repo.GetCommunityAll()
}

func (uc *CommunityUsecase) GetCommunityById(id int) (*entity.Community, error) {
	if id == 0 {
		return nil, fmt.Errorf("Usecase.GetCommunityById Error: id invalied")
	}
	return uc.repo.GetCommunityById(id)
}

func (uc *CommunityUsecase) UpdateCommunityOfId(id int, community *entity.Community)(*entity.Community,error){
	if id == 0 {
		return nil,fmt.Errorf("Usecase.UpdateCommunityOfId Error: id invalied")
	}
	community.Id = id
	return uc.repo.UpdateCommunityOfId(community)
}

func (uc *CommunityUsecase) DeleteCommunityOfId(id int)error{
	if id == 0 {
		return fmt.Errorf("Usecase.DeleteCommunityOfId Error: id invalied")
	}
	return uc.repo.DeleteCommunityOfId(id)
}

func (uc *CommunityUsecase) CreateNewCommunity(community *entity.Community)(*entity.Community,error){
	if community.Name == "" {
		return nil,fmt.Errorf("Usecase.CreateNewCommunity Error: name is empty")
	}
	return uc.repo.CreateNewCommunity(community)
}

