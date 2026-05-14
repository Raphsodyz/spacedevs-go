package models

type SearchLaunchFilters struct {
	MissionIds  []int64
	RocketIds   []int64
	LocationIds []int64
	PadIds      []int64
	LaunchIds   []int64
	Page        int
}
