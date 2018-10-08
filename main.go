package main

import (
	"net/http"
	"os"
	"time"

	"github.com/Albert221/medicinal-products-registry-api/api"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	downloader := Downloader{TargetFilename: "dataset.xml"}

	updatedChan := make(chan bool)
	go downloader.ScheduleDownloads(20*time.Minute, updatedChan)

	schema, err := api.NewSchema()
	if err != nil {
		panic(err)
	}

	<-updatedChan
	err = schema.RefershMedicalProducts()
	if err != nil {
		panic(err)
	}

	http.Handle("/query", &relay.Handler{Schema: schema.CreateGraphQLSchema()})
	if err := http.ListenAndServe(os.Getenv("MPR_ADDR"), nil); err != nil {
		panic(err)
	}
}
