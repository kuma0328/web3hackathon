package repository

import (
	"context"

	"github.com/kuma0328/web3hackathon/domain/entity"
)

type UserRepository interface {
	CreateUser(context.Context, *entity.User) (*entity.User, error)
	UpdateUser(context.Context, *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, id string) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
}
