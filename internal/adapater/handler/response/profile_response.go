package response

type ProfileResponse struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	Content      string `json:"content"`
}
