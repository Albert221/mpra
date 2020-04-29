package graphql

import (
	"io/ioutil"
	"strings"
	"sync"

	"github.com/Albert221/mpra/domain"
	gql "github.com/graph-gophers/graphql-go"
	"github.com/markbates/pkger"
	"github.com/pkg/errors"
)

type Query struct {
	lock     sync.RWMutex
	products []*domain.Product
}

func NewQuery() *Query {
	return &Query{}
}

func (q *Query) CreateSchema() (*gql.Schema, error) {
	f, err := pkger.Open("/schema.graphql")
	if err != nil {
		return nil, errors.Wrap(err, "opening graphql schema file")
	}
	defer f.Close()

	schema, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "reading graphql schema file")
	}

	return gql.MustParseSchema(string(schema), q), nil
}

func (q *Query) Populate(products []*domain.Product) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.products = products
}

func (q *Query) Products(args struct{ Ean string }) []*Product {
	q.lock.RLock()
	defer q.lock.RUnlock()

	var results []*Product

	for _, product := range q.products {
		for _, pack := range product.Packages {
			if strings.Contains(pack.Ean, args.Ean) {
				results = append(results, &Product{product})
			}
		}
	}

	return results
}
