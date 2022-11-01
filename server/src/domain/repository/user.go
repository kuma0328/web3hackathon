package repository

import (
	"context"

	"github.com/kuma0328/web3hackathon/domain/entity"
)

type IUserRepository interface {
	CreateUser(context.Context, *entity.User) (*entity.User, error)
	UpdateUser(context.Context, *entity.User) (*entity.User, error)
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	GetUserByMail(ctx context.Context, mail string) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}
