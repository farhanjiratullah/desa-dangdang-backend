package request

type StatisticRequest struct {
	Name  string `json:"name" validate:"required"`
	Total int64  `json:"total" validate:"required"`
	Icon  string `json:"icon" validate:"required"`
}
