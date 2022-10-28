package repository

import "github.com/kuma0328/web3hackathon/domain/entity"

type ICommunityRepository interface {
	GetCommunityById(id int) (*entity.Community, error)
	UpdateCommunityOfId(community *entity.Community)(*entity.Community,error)
	DeleteCommunityOfId(id int)error
	CreateNewCommunity(community *entity.Community)(*entity.Community,error)
}
