package models

type SearchLaunchRequest struct {
	Mission  *string `json:"mission" validate:"omitempty,lte=360"`
	Rocket   *string `json:"rocket" validate:"omitempty,lte=360"`
	Location *string `json:"location" validate:"omitempty,lte=360"`
	Pad      *string `json:"pad" validate:"omitempty,lte=360"`
	Page     *int    `json:"page" validate:"omitempty,gte=0"`
}
