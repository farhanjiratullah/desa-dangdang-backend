package request

type ProfileRequest struct {
	Title         string `json:"title" validate:"required"`
	Content      string `json:"content" validate:"required"`
}
