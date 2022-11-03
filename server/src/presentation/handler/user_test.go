package handler

import (
	"context"

	"github.com/kuma0328/web3hackathon/domain/entity"
)

var (
	testCreateUser *entity.User = &entity.User{
		Id:       1,
		Name:     "hoge",
		Mail:     "hoge@hoge.com",
		Password: "hogehoge",
	}
	testUpdateUser *entity.User = &entity.User{
		Id:       1,
		Name:     "hoge",
		Mail:     "hoge2@hoge.com",
		Password: "hogehoge2",
	}
)

type UserUsecaseMock struct{}

func (h *UserUsecaseMock) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return testCreateUser, nil
}

func (ur UserUsecaseMock) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return testUpdateUser, nil
}

func (ur UserUsecaseMock) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	return testCreateUser, nil
}

func (ur UserUsecaseMock) DeleteUser(ctx context.Context, id int) error {
	return nil
}

func (ur UserUsecaseMock) GetUserByMail(ctx context.Context, mail string) (*entity.User, error) {
	return testCreateUser, nil
}

func (ur UserUsecaseMock) LoginUser(ctx context.Context, user *entity.User) error {
	return nil
}
