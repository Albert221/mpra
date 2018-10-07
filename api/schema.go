package api

import (
	"encoding/xml"
	"github.com/Albert221/medicinal-products-registry-api/data"
	"github.com/graph-gophers/graphql-go"
	"io/ioutil"
	"log"
)

type Schema struct {
	medicalProducts *data.MedicalProducts
}

func NewSchema(updateChan chan bool) *Schema {
	schema := &Schema{getMedicalProducts()}

	go func() {
		for {
			<-updateChan
			schema.medicalProducts = getMedicalProducts()
		}
	}()

	return schema
}

func getMedicalProducts() *data.MedicalProducts {
	file, err := ioutil.ReadFile("downloads/dataset.xml")
	if err != nil {
		panic(err)
	}

	var products *data.MedicalProducts
	if err = xml.Unmarshal(file, &products); err != nil {
		log.Println(err)
	}

	return products
}

func (s *Schema) CreateGraphQLSchema() *graphql.Schema {
	file, err := ioutil.ReadFile("schema.graphql")
	if err != nil {
		panic(err)
	}

	return graphql.MustParseSchema(string(file), s)
}
