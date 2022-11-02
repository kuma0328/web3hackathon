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
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	GetUserByMail(ctx context.Context, mail string) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	LoginUser(ctx context.Context, user *entity.User) error
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		repo: ur,
	}
}

func (uu *UserUsecase) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if user.Name == "" {
		return nil, usecase_error.NameEmptyError
	}
	if user.Mail == "" {
		return nil, usecase_error.MailEmptyError
	}
	if user.Password == "" {
		return nil, usecase_error.PassWordEmptyError
	}

	user, err := uu.repo.CreateUser(ctx, user)
	return user, err
}

func (uu *UserUsecase) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	if user.Name == "" {
		return nil, usecase_error.NameEmptyError
	}
	if user.Mail == "" {
		return nil, usecase_error.MailEmptyError
	}
	if user.Password == "" {
		return nil, usecase_error.PassWordEmptyError
	}

	user, err := uu.repo.UpdateUser(ctx, user)
	return user, err
}

func (uu *UserUsecase) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	if id == 0 {
		return nil, usecase_error.IdEmptyError
	}

	user, err := uu.repo.GetUserByID(ctx, id)
	return user, err
}

func (uu *UserUsecase) GetUserByMail(ctx context.Context, mail string) (*entity.User, error) {
	if mail == "" {
		return nil, usecase_error.MailEmptyError
	}

	user, err := uu.repo.GetUserByMail(ctx, mail)
	return user, err
}

func (uu *UserUsecase) DeleteUser(ctx context.Context, id int) error {
	if id == 0 {
		return usecase_error.IdEmptyError
	}

	err := uu.repo.DeleteUser(ctx, id)
	return err
}

func (uu *UserUsecase) LoginUser(ctx context.Context, user *entity.User) error {
	if user.Mail == "" {
		return usecase_error.MailEmptyError
	}
	if user.Password == "" {
		return usecase_error.PassWordEmptyError
	}

	err := uu.repo.LoginUser(ctx, user)
	return err
}
