package graphql

import "github.com/Albert221/mpra/domain"

type Product struct {
	product *domain.Product
}

func (p *Product) Name() string {
	return p.product.ProductName
}

func (p *Product) CommonName() *string {
	return &p.product.CommonProductName
}

func (p *Product) Type() string {
	return p.product.Kind
}

func (p *Product) Strength() string {
	return p.product.Strength
}

func (p *Product) Form() string {
	return p.product.Form
}

func (p *Product) ResponsibleEntity() string {
	return p.product.ResponsibleEntity
}

func (p *Product) ActiveSubstances() []string {
	return p.product.ActiveSubstances
}

func (p *Product) Packages() []*Package {
	var packages []*Package
	for _, pack := range p.product.Packages {
		packages = append(packages, &Package{pack})
	}

	return packages
}
