package service

import (
	"context"
	"desadangdang/internal/adapater/repository"
	"desadangdang/internal/core/domain/entity"
)

type PostServiceInterface interface {
	CreatePost(ctx context.Context, req entity.PostEntity) error
	FetchAllPosts(ctx context.Context) ([]entity.PostEntity, error)
	FetchByIDPost(ctx context.Context, id int64) (*entity.PostEntity, error)
	FetchBySlugPost(ctx context.Context, slug string) (*entity.PostEntity, error)
	EditByIDPost(ctx context.Context, req entity.PostEntity) error
	DeleteByIDPost(ctx context.Context, id int64) error
}

type postService struct {
	postRepo repository.PostInterface
}

// CreatePost implements PostServiceInterface.
func (p *postService) CreatePost(ctx context.Context, req entity.PostEntity) error {
	return p.postRepo.CreatePost(ctx, req)
}

// DeleteByIDPost implements PostServiceInterface.
func (p *postService) DeleteByIDPost(ctx context.Context, id int64) error {
	return p.postRepo.DeleteByIDPost(ctx, id)
}

// EditByIDPost implements PostServiceInterface.
func (p *postService) EditByIDPost(ctx context.Context, req entity.PostEntity) error {
	return p.postRepo.EditByIDPost(ctx, req)
}

// FetchAllPosts implements PostServiceInterface.
func (p *postService) FetchAllPosts(ctx context.Context) ([]entity.PostEntity, error) {
	return p.postRepo.FetchAllPosts(ctx)
}

// FetchByIDPost implements PostServiceInterface.
func (p *postService) FetchByIDPost(ctx context.Context, id int64) (*entity.PostEntity, error) {
	return p.postRepo.FetchByIDPost(ctx, id)
}

func (p *postService) FetchBySlugPost(ctx context.Context, slug string) (*entity.PostEntity, error) {
	return p.postRepo.FetchBySlugPost(ctx, slug)
}

func NewPostService(postRepo repository.PostInterface) PostServiceInterface {
	return &postService{
		postRepo: postRepo,
	}
}
