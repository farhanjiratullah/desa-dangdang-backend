package request

type PostRequest struct {
	Title         string `json:"title" validate:"required"`
	Slug          string `json:"slug"`
	Author       string `json:"author" validate:"required"`
	FeaturedImage string `json:"featured_image" validate:"required"`
	Content      string `json:"content" validate:"required"`
	PublishedAt  string `json:"published_at" validate:"required"`
}
