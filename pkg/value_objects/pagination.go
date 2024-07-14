package value_objects

type PaginationParams[T any] struct {
	Page    int
	PerPage int
	Filter  T
}

type PaginationResponse[T any] struct {
	Data        []T
	Total       int64
	PerPage     int
	CurrentPage int
}
