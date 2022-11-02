package persistance

import (
	"context"
	"database/sql"
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
	INSERT INTO users (name, mail,password)
	VALUES (:name,:mail,:password)
	`
	hashpass, err := PasswordEncrypt(user.Password)
	user.Password = hashpass
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
	UPDATE users
	SET name     = :name,
		mail 	 = :mail,
		password = :password
	WHERE id = :id
	`

	hashpass, err := PasswordEncrypt(user.Password)
	user.Password = hashpass
	dto := userEntityToDto(user)
	_, err = ur.conn.DB.NamedExecContext(ctx, query, &dto)

	if err != nil {
		return nil, fmt.Errorf("UserRepository.CreateNewUser NamedExec Error : %w", err)
	}
	return userDtoToEntity(&dto), nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	var dto userDto

	query := `
	SELECT id, name, mail  
	FROM users
	WHERE id = ?
	`

	err := ur.conn.DB.GetContext(ctx, &dto, query, id)
	if err != nil {
		return nil, fmt.Errorf("CommunityRepository.GetCommunityById Get Error : %w", err)
	}
	return userDtoToEntity(&dto), nil
}

func (ur *UserRepository) GetUserByMail(ctx context.Context, mail string) (*entity.User, error) {
	var dto userDto

	query := `
	SELECT * 
	FROM users
	WHERE mail = ?
	LIMIT 1
	`

	err := ur.conn.DB.Get(&dto, query, mail)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return userDtoToEntity(&dto), nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id int) error {
	query := `
	DELETE FROM users
	WHERE id = :id
	`

	_, err := ur.conn.DB.NamedExecContext(ctx, query, map[string]interface{}{"id": id})
	if err != nil {
		return fmt.Errorf("CommunityRepository.DeleteCommunityOfId NamedExec Error : %w", err)
	}

	return nil
}

func (ur *UserRepository) LoginUser(ctx context.Context, user *entity.User) error {
	dbuser, err := ur.GetUserByMail(ctx, user.Mail)
	if err != nil {
		return fmt.Errorf("GetUserByID Error : %w", err)
	}
	if dbuser.Name == "" {
		return fmt.Errorf("mail not found")
	}

	err = CompareHashAndPassword(dbuser.Password, user.Password)

	if err != nil {
		return fmt.Errorf("PassWord not ok Error : %w", err)
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
