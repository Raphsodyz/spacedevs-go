package models

type SearchLaunchFilters struct {
	MissionIds  []int64
	RocketIds   []int64
	LocationIds []int64
	PadIds      []int64
	LaunchIds   []int64
	Page        int
}

type SearchLaunchRequest struct {
	Mission  *string `json:"mission" validate:"omitempty,lte=360"`
	Rocket   *string `json:"rocket" validate:"omitempty,lte=360"`
	Location *string `json:"location" validate:"omitempty,lte=360"`
	Pad      *string `json:"pad" validate:"omitempty,lte=360"`
	Launch   *string `json:"launch" validate:"omitempty,lte=360"`
	Page     *int    `json:"page" validate:"omitempty,gte=1"`
}

type SearchLaunchResponse struct {
	Entities         []LaunchView
	NumberOfPages    int
	CurrentPage      int
	NumberOfEntities int
}
