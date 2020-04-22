package graphql

import (
	"strconv"

	"github.com/Albert221/medicinal-products-registry-api/domain"
)

type Package struct {
	pack *domain.Package
}

func (p *Package) Ean() string {
	return p.pack.Ean
}

func (p *Package) Size() (int32, error) {
	num, err := strconv.Atoi(p.pack.Size)
	return int32(num), err
}

func (p *Package) SizeUnit() string {
	return p.pack.SizeUnit
}
