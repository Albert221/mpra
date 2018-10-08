package api

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/Albert221/medicinal-products-registry-api/data"
	graphql "github.com/graph-gophers/graphql-go"
)

const (
	medProductsDataSetPath string = "downloads/dataset.xml"
	graphqlSchemaPath      string = "schema.graphql"
)

// Schema represents API Schema
type Schema struct {
	medicalProducts *data.MedicalProducts
}

// NewSchema return a pointer to API Schema
func NewSchema() (*Schema, error) {
	pp, err := getMedicalProducts()
	if err != nil {
		return nil, err
	}
	return &Schema{medicalProducts: pp}, nil
}

// CreateGraphQLSchema returns graphql API schema
func (s *Schema) CreateGraphQLSchema() *graphql.Schema {
	file, err := ioutil.ReadFile(graphqlSchemaPath)
	if err != nil {
		panic(err)
	}

	return graphql.MustParseSchema(string(file), s)
}

// RefershMedicalProducts updates schema's medical products
func (s *Schema) RefershMedicalProducts() error {
	pp, err := getMedicalProducts()
	if err != nil {
		return err
	}

	s.medicalProducts = pp
	return nil
}

func getMedicalProducts() (*data.MedicalProducts, error) {
	file, err := ioutil.ReadFile(medProductsDataSetPath)
	if err != nil {
		return nil, err
	}

	var products data.MedicalProducts
	err = xml.Unmarshal(file, &products)
	return &products, err
}
