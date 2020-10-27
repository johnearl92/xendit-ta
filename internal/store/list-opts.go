package store

import "strings"

// NewListOpts provides ListOpts
func NewListOpts() ListOpts {
	opts := &BaseListOpts{}
	opts.Init()
	return opts
}

// ListOpts interface
type ListOpts interface {
	By() interface{}
	SetBy(arg interface{}) ListOpts
	Offset() *int
	SetOffset(*int) ListOpts
	SetOffsetVal(int) ListOpts
	Max() *int
	SetMax(*int) ListOpts
	SetMaxVal(int) ListOpts
	Sort() []*Sort
	SortString() string
	SetSort([]*Sort) ListOpts
	Preload() []string
	SetPreload(...string) ListOpts
}

// BaseListOpts represents listing options
type BaseListOpts struct {
	offset  *int
	max     *int
	sort    []*Sort
	by      interface{}
	preload []string
}

// Preload getter
func (b *BaseListOpts) Preload() []string {
	return b.preload
}

// SetPreload setter
func (b *BaseListOpts) SetPreload(arg ...string) ListOpts {
	b.preload = append(b.preload, arg...)

	return b
}

// Init initializer
func (b *BaseListOpts) Init() {
	b.offset = nil
	b.max = nil
}

// By gettter
func (b *BaseListOpts) By() interface{} {
	return b.by
}

// SetBy setter
func (b *BaseListOpts) SetBy(by interface{}) ListOpts {
	b.by = by
	return b
}

// Offset getter
func (b *BaseListOpts) Offset() *int {
	return b.offset
}

// SetOffset setter
func (b *BaseListOpts) SetOffset(arg *int) ListOpts {
	b.offset = arg
	return b
}

// SetOffsetVal setter
func (b *BaseListOpts) SetOffsetVal(arg int) ListOpts {
	b.offset = &arg
	return b
}

// Max getter
func (b *BaseListOpts) Max() *int {
	return b.max
}

// SetMax setter
func (b *BaseListOpts) SetMax(arg *int) ListOpts {
	b.max = arg
	return b
}

// SetMaxVal setter
func (b *BaseListOpts) SetMaxVal(arg int) ListOpts {
	b.max = &arg
	return b
}

// Sort getter
func (b *BaseListOpts) Sort() []*Sort {
	return b.sort
}

// SortString sorting
func (b *BaseListOpts) SortString() string {
	orderBy := []string{}

	for _, sort := range b.Sort() {
		order := "asc"

		if sort.Order() == SortOrderDesc {
			order = "desc"
		}

		orderBy = append(orderBy, sort.Name()+" "+order)
	}

	return strings.Join(orderBy, ", ")
}

// SetSort setter
func (b *BaseListOpts) SetSort(arg []*Sort) ListOpts {
	b.sort = arg
	return b
}
