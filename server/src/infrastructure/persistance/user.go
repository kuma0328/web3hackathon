package persistance

import (
	"context"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
	"github.com/kuma0328/web3hackathon/infrastructure/database"
)

var _ repository.IUserRepository = &UserRepository{}

type UserRepository struct {
	conn *database.Conn
}

func NewUserRepository(conn *database.Conn) repository.IUserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {

	return nil, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {

	return nil, nil
}

func (ur *UserRepository) GetUser(ctx context.Context, id string) (*entity.User, error) {

	return nil, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id string) error {

	return nil
}

type userDto struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Mail     string `db:"mail"`
	Password string `db:"password"`
}

func userDtoToEntity(dto *userDto) *entity.User {
	return &entity.User{
		Id:       dto.Id,
		Name:     dto.Name,
		Mail:     dto.Mail,
		Password: dto.Password,
	}
}

func userEntityToDto(u *entity.User) userDto {
	return userDto{
		Id:       u.Id,
		Name:     u.Name,
		Mail:     u.Mail,
		Password: u.Password,
	}
}
