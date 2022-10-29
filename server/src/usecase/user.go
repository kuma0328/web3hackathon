package usecase

import (
	"context"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
	usecase_error "github.com/kuma0328/web3hackathon/error/usecase"
)

var _ IUserUsecase = &UserUsecase{}

type UserUsecase struct {
	repo repository.IUserRepository
}

type IUserUsecase interface {
	CreateUser(context.Context, *entity.User) (*entity.User, error)
	UpdateUser(context.Context, *entity.User) (*entity.User, error)
	GetUser(ctx context.Context, id string) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		repo: ur,
	}
}

func (ur *UserUsecase) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if user.Name == "" {
		return nil, usecase_error.IdEmptyError
	}
	if user.Mail == "" {
		return nil, usecase_error.NameEmptyError
	}
	if user.Password == "" {
		return nil, usecase_error.PassWordEmptyError
	}

	user, err := ur.repo.CreateUser(ctx, user)
	return user, err
}

func (ur *UserUsecase) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if user.Name == "" {
		return nil, usecase_error.IdEmptyError
	}
	if user.Mail == "" {
		return nil, usecase_error.NameEmptyError
	}
	if user.Password == "" {
		return nil, usecase_error.PassWordEmptyError
	}

	user, err := ur.repo.UpdateUser(ctx, user)
	return user, err
}

func (ur *UserUsecase) GetUser(ctx context.Context, id string) (*entity.User, error) {
	user, err := ur.repo.GetUser(ctx, id)
	return user, err
}

func (ur *UserUsecase) DeleteUser(ctx context.Context, id string) error {
	err := ur.repo.DeleteUser(ctx, id)
	return err
}
