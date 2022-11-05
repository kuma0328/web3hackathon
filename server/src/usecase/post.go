package usecase

import (
	"context"

	"github.com/kuma0328/web3hackathon/domain/entity"
	"github.com/kuma0328/web3hackathon/domain/repository"
	usecase_error "github.com/kuma0328/web3hackathon/error/usecase"
)

var _ IPostUsecase = &PostUsecase{}

type PostUsecase struct {
	repo repository.IPostRepository
}

type IPostUsecase interface {
	GetPostById(ctx context.Context, id int) (*entity.Post, error)
	UpdatePostOfId(ctx context.Context, post *entity.Post) (*entity.Post, error)
	DeletePostOfId(ctx context.Context, id int) error
	CreateNewPost(ctx context.Context, post *entity.Post) (*entity.Post, error)
}

func NewPostUsecase(repo repository.IPostRepository) IPostUsecase {
	return &PostUsecase{
		repo: repo,
	}
}

func (pu *PostUsecase) GetPostById(ctx context.Context, id int) (*entity.Post, error) {
	if id == 0 {
		return nil, usecase_error.IdEmptyError
	}
	post, err := pu.repo.GetPostById(ctx, id)
	return post, err
}

func (pu *PostUsecase) UpdatePostOfId(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	if post.Id == 0 {
		return nil, usecase_error.IdEmptyError
	}
	if post.Content == "" {
		return nil, usecase_error.ContentEmptyError
	}
	if post.Img == nil {
		return nil, usecase_error.ImgEmptyError
	}
	if post.CommunityId == 0 {
		return nil, usecase_error.CommunityIdEmptyError
	}
	post, err := pu.repo.UpdatePostOfId(ctx, post)
	return post, err
}

func (pu *PostUsecase) DeletePostOfId(ctx context.Context, id int) error {
	if id == 0 {
		return usecase_error.IdEmptyError
	}
	err := pu.repo.DeletePostOfId(ctx, id)
	return err
}

func (pu *PostUsecase) CreateNewPost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	if post.Content == "" {
		return nil, usecase_error.ContentEmptyError
	}
	if post.Img == nil {
		return nil, usecase_error.ImgEmptyError
	}
	if post.CommunityId == 0 {
		return nil, usecase_error.CommunityIdEmptyError
	}

	post, err := pu.repo.CreateNewPost(ctx, post)
	return post, err
}
