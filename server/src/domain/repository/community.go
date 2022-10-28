package repository

import "github.com/kuma0328/web3hackathon/domain/entity"

type ICommunityRepository interface {
	GetCommunityById(id string) (*entity.Community, error)
}
