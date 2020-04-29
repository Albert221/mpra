package graphql

import (
	"github.com/Albert221/mpra/domain"
)

type Package struct {
	pack *domain.Package
}

func (p *Package) Ean() string {
	return p.pack.Ean
}

func (p *Package) Size() string {
	return p.pack.Size
}

func (p *Package) SizeUnit() string {
	return p.pack.SizeUnit
}
