package request

type ClientSectionRquest struct {
	Name     string `json:"name" validate:"required"`
	PathIcon string `json:"path_icon" validate:"required"`
}
