package persistance

import (
	"context"
	"log"

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
	statement := "INSERT INTO users (name,mail) VALUES($1,$2)"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Name, user.Mail)

	if err != nil {
		log.Println(err)
		return nil, db_error.QueryError
	}

	resuser := &entity.User{}
	resuser.Id = user.Id
	resuser.Name = user.Name
	resuser.Mail = user.Mail

	return resuser, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	statement := "UPDATE users SET name = $2, mail = $3 WHERE id = $1"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return nil, db_error.StatementError
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Id, user.Name, user.Mail)

	if err != nil {
		log.Println(err)
		return nil, db_error.ExecError
	}

	resuser := &entity.User{}
	resuser.Id = user.Id
	resuser.Name = user.Name
	resuser.Mail = user.Mail

	return resuser, nil
}

func (ur *UserRepository) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	statement := "SELECT * FROM users WHERE id = $1"

	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return nil, db_error.StatementError
	}
	defer stmt.Close()
	resuser := &entity.User{}

	err = stmt.QueryRowContext(ctx, id).Scan(&resuser.Id, &resuser.Name, &resuser.Mail)

	if err != nil {
		log.Println(err)
		return nil, db_error.QueryError
	}

	return resuser, nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	statement := "DELETE FROM users WHERE id = $1"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return db_error.StatementError
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)

	if err != nil {
		log.Println(err)
		return db_error.ExecError
	}

	return nil
}
