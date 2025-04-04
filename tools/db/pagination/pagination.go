package pagination

type Pagination interface {
	GetPageNumber() int32
	GetPageSize() int32
}
