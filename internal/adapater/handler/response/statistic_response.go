package response

type StatisticResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Total int64  `json:"total"`
	Icon  string `json:"icon"`
}
