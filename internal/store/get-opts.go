package store

// NewGetOpts provides getOpts
func NewGetOpts() GetOpts {
	return &getOpts{}
}

// GetOpts getOpts interface
type GetOpts interface {
	Preload() []string
	SetPreload(arg ...string) GetOpts
}

// getOpts get options for gorm
type getOpts struct {
	preload []string
}

// Preload getter
func (g *getOpts) Preload() []string {
	return g.preload
}

// SetPreload setter
func (g *getOpts) SetPreload(arg ...string) GetOpts {
	g.preload = append(g.preload, arg...)

	return g
}
