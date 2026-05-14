package models

type Pagination[T any] struct {
	Entities         []T
	NumberOfPages    int
	CurrentPage      int
	NumberOfEntities int
}
