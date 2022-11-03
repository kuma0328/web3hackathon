package persistance

import (
	"context"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
	"github.com/kuma0328/web3hackathon/infrastructure/database"
)

var _ repository.IPostRepository = &PostRepository{}

type PostRepository struct {
	conn *database.Conn
}

func NewPostRepository(conn *database.Conn) repository.IPostRepository {
	return &PostRepository{
		conn: conn,
	}
}

func (repo *PostRepository) GetPostById(ctx context.Context, id int) (*entity.Post, error) {
	return nil, nil
}

func (repo *PostRepository) UpdatePostOfId(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	return nil, nil
}
func (repo *PostRepository) DeletePostOfId(ctx context.Context, id int) error {
	return nil
}
func (repo *PostRepository) CreateNewPost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	return nil, nil
}
