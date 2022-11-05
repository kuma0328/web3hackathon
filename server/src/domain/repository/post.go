package repository

import (
	"context"

	"github.com/kuma0328/web3hackathon/domain/entity"
)

type IPostRepository interface {
	GetPostById(ctx context.Context, id int) (*entity.Post, error)
	UpdatePostOfId(ctx context.Context, post *entity.Post) (*entity.Post, error)
	DeletePostOfId(ctx context.Context, id int) error
	CreateNewPost(ctx context.Context, post *entity.Post) (*entity.Post, error)
}
