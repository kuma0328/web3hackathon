package persistance

import (
	"context"
	"fmt"

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
	query := `
	INSERT INTO users (name, mail, password)
	VALUES (:name,:img_url,:content)
	`
	dto := userEntityToDto(user)
	res, err := ur.conn.DB.NamedExecContext(ctx, query, &dto)

	id, _ := res.LastInsertId()
	dto.Id = (int)(id)

	if err != nil {
		return nil, fmt.Errorf("UserRepository.CreateNewUser NamedExec Error : %w", err)
	}
	return userDtoToEntity(&dto), nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `
	UPDATE communities
	SET name     = :name,
		mail 	 = :mail,
		passqord = :password
	WHERE id = :id
	`
	dto := userEntityToDto(user)
	_, err := ur.conn.DB.NamedExecContext(ctx, query, &dto)

	if err != nil {
		return nil, fmt.Errorf("UserRepository.CreateNewUser NamedExec Error : %w", err)
	}
	return userDtoToEntity(&dto), nil
}

func (ur *UserRepository) GetUser(ctx context.Context, id string) (*entity.User, error) {
	var dto userDto

	query := `
	SELECT * 
	FROM users
	WHERE id = ?
	`

	err := ur.conn.DB.GetContext(ctx, &dto, query, id)
	if err != nil {
		return nil, fmt.Errorf("CommunityRepository.GetCommunityById Get Error : %w", err)
	}
	return userDtoToEntity(&dto), nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id string) error {
	query := `
	DELETE FROM communities
	WHERE id = :id
	`

	_, err := ur.conn.DB.NamedExec(query, map[string]interface{}{"id": id})
	if err != nil {
		return fmt.Errorf("CommunityRepository.DeleteCommunityOfId NamedExec Error : %w", err)
	}

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
