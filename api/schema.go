package api

import (
	"encoding/xml"
	"io/ioutil"
	"log"

	"github.com/Albert221/medicinal-products-registry-api/data"
	graphql "github.com/graph-gophers/graphql-go"
)

const (
	medProductsDataSetPath string = "downloads/dataset.xml"
	graphqlSchemaPath      string = "schema.graphql"
)

// Schema represents API Schema
type Schema struct {
	updateChan      chan bool
	medicalProducts *data.MedicalProducts
}

// NewSchema return a pointer to API Schema
func NewSchema(updateChan chan bool) (*Schema, error) {
	products, err := getMedicalProducts()
	if err != nil {
		return nil, err
	}
	schema := Schema{medicalProducts: products, updateChan: updateChan}

	// run service for schema update
	go schema.runUpdateService()
	return &schema, nil
}

// CreateGraphQLSchema returns graphql API schema
func (s *Schema) CreateGraphQLSchema() *graphql.Schema {
	file, err := ioutil.ReadFile(graphqlSchemaPath)
	if err != nil {
		panic(err)
	}

	return graphql.MustParseSchema(string(file), s)
}

func (s *Schema) runUpdateService() {
	for {
		select {
		case <-s.updateChan:
			products, err := getMedicalProducts()
			if err != nil {
				log.Println(err)
			}

			s.medicalProducts = products
		}
	}
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
