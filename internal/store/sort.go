package store

// NewSort Sort provider
func NewSort(name string, order int) *Sort {
	return &Sort{name: name, order: order}
}

// Sort definition
type Sort struct {
	name  string
	order int
}

// Name getter
func (s *Sort) Name() string {
	return s.name
}

// SetName setter
func (s *Sort) SetName(arg string) *Sort {
	s.name = arg
	return s
}

// Order getter
func (s *Sort) Order() int {
	return s.order
}

// SortOrder ascending or descending order
const (
	SortOrderAsc  = iota
	SortOrderDesc = iota
)
