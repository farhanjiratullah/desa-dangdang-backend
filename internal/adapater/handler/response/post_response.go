package response

type PostResponse struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	Slug          string `json:"slug"`
	Author       string `json:"author"`
	FeaturedImage string `json:"featured_image"`
	Content      string `json:"content"`
	PublishedAt  string `json:"published_at"`
}
