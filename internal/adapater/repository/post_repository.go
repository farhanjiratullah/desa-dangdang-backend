package repository

import (
	"context"
	"desadangdang/internal/core/domain/entity"
	"desadangdang/internal/core/domain/model"
	"fmt"

	"github.com/gosimple/slug"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PostInterface interface {
	CreatePost(ctx context.Context, req entity.PostEntity) error
	FetchAllPosts(ctx context.Context) ([]entity.PostEntity, error)
	FetchByIDPost(ctx context.Context, id int64) (*entity.PostEntity, error)
	FetchBySlugPost(ctx context.Context, slug string) (*entity.PostEntity, error)
	EditByIDPost(ctx context.Context, req entity.PostEntity) error
	DeleteByIDPost(ctx context.Context, id int64) error
	CheckSlugUnique(slug string, id int64) bool
}

type post struct {
	DB *gorm.DB
}

func (p *post) CheckSlugUnique(slug string, id int64) bool {
	var count int64
	err := p.DB.Model(&model.Post{}).Where("slug = ? AND id != ?", slug, id).Count(&count).Error
	if err != nil {
		log.Errorf("[REPOSITORY] CheckSlugUnique - 1: %v", err)
		return false
	}
	return count == 0
}

// CreatePost implements PostInterface.
func (p *post) CreatePost(ctx context.Context, req entity.PostEntity) error {
	// Generate slug from title if it's empty or invalid
	if req.Slug == "" {
		req.Slug = slug.Make(req.Title)
	}

	// Check if the slug is unique
	if !p.CheckSlugUnique(req.Slug, 0) { // Passing 0 for the ID because it's a new post
		log.Errorf("[REPOSITORY] CreatePost - Slug '%s' already exists", req.Slug)
		return fmt.Errorf("slug '%s' already exists", req.Slug)
	}

	modelPost := model.Post{
		Title:         req.Title,
		Slug:          req.Slug,
		Author:        req.Author,
		FeaturedImage: req.FeaturedImage,
		Content:       req.Content,
		PublishedAt:   req.PublishedAt,
	}

	if err := p.DB.Create(&modelPost).Error; err != nil {
		log.Errorf("[REPOSITORY] CreatePost - 1: %v", err)
		return err
	}
	return nil
}

// DeleteByIDPost implements PostInterface.
func (p *post) DeleteByIDPost(ctx context.Context, id int64) error {
	modelPost := model.Post{}

	err := p.DB.Where("id = ?", id).First(&modelPost).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPost - 1: %v", err)
		return err
	}

	err = p.DB.Delete(&modelPost).Error
	if err != nil {
		log.Errorf("[REPOSITORY] DeleteByIDPost - 2: %v", err)
		return err
	}
	return nil
}

// EditByIDPost implements PostInterface.
func (p *post) EditByIDPost(ctx context.Context, req entity.PostEntity) error {
	// Generate slug from title if it's empty or invalid
	if req.Slug == "" {
		req.Slug = slug.Make(req.Title)
	}

	// Check if the slug is unique (except the post with the same ID)
	if !p.CheckSlugUnique(req.Slug, req.ID) { // Passing the actual ID of the post being edited
		log.Errorf("[REPOSITORY] EditByIDPost - Slug '%s' already exists", req.Slug)
		return fmt.Errorf("slug '%s' already exists", req.Slug)
	}

	modelPost := model.Post{}

	err := p.DB.Where("id = ?", req.ID).First(&modelPost).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDPost - 1: %v", err)
		return err
	}

	// Update the post fields
	modelPost.Title = req.Title
	modelPost.Slug = req.Slug
	modelPost.Author = req.Author
	modelPost.FeaturedImage = req.FeaturedImage
	modelPost.Content = req.Content
	modelPost.PublishedAt = req.PublishedAt

	// Save the updated post
	err = p.DB.Save(&modelPost).Error
	if err != nil {
		log.Errorf("[REPOSITORY] EditByIDPost - 2: %v", err)
		return err
	}
	return nil
}

// FetchAllPosts implements PostInterface.
func (p *post) FetchAllPosts(ctx context.Context) ([]entity.PostEntity, error) {
	modelPosts := []model.Post{}
	err := p.DB.Select("id", "title", "slug", "author", "featured_image", "content", "published_at").Find(&modelPosts).Order("created_at DESC").Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchAllPosts - 1: %v", err)
		return nil, err
	}

	var postEntities []entity.PostEntity
	for _, v := range modelPosts {
		postEntities = append(postEntities, entity.PostEntity{
			ID:           v.ID,
			Title:        v.Title,
			Slug:         v.Slug,
			Author:       v.Author,
			FeaturedImage: v.FeaturedImage,
			Content:      v.Content,
			PublishedAt:  v.PublishedAt, // Include PublishedAt field
		})
	}

	return postEntities, nil
}

// FetchByIDPost implements PostInterface.
func (p *post) FetchByIDPost(ctx context.Context, id int64) (*entity.PostEntity, error) {
	modelPost := model.Post{}
	err := p.DB.Where("id = ?", id).First(&modelPost).Error
	if err != nil {
		log.Errorf("[REPOSITORY] FetchByIDPost - 1: %v", err)
		return nil, err
	}

	return &entity.PostEntity{
		ID:           modelPost.ID,
		Title:        modelPost.Title,
		Slug:         modelPost.Slug,
		Author:       modelPost.Author,
		FeaturedImage: modelPost.FeaturedImage,
		Content:      modelPost.Content,
		PublishedAt:  modelPost.PublishedAt, // Include PublishedAt field
	}, nil
}

func (p *post) FetchBySlugPost(ctx context.Context, slug string) (*entity.PostEntity, error) {
    modelPost := model.Post{}
    err := p.DB.Where("slug = ?", slug).First(&modelPost).Error
    if err != nil {
        log.Errorf("[REPOSITORY] FetchBySlugPost - 1: %v", err)
        return nil, err
    }

    return &entity.PostEntity{
        ID:            modelPost.ID,
        Title:         modelPost.Title,
        Slug:          modelPost.Slug,
        Author:        modelPost.Author,
        FeaturedImage: modelPost.FeaturedImage,
        Content:       modelPost.Content,
        PublishedAt:   modelPost.PublishedAt,
    }, nil
}

func NewPostRepository(DB *gorm.DB) PostInterface {
	return &post{
		DB: DB,
	}
}
