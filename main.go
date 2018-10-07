package main

import (
	"github.com/Albert221/medicinal-products-registry-api/api"
	"github.com/graph-gophers/graphql-go/relay"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	downloader := Downloader{TargetFilename: "dataset.xml"}

	updatedChan := make(chan bool)
	go downloader.ScheduleDownloads(2 * time.Second, updatedChan)

	<-updatedChan
	schema := api.NewSchema(updatedChan)

	http.Handle("/query", &relay.Handler{Schema: schema.CreateGraphQLSchema()})
	if err := http.ListenAndServe(os.Getenv("MPR_ADDR"), nil); err != nil {
		log.Println(err)
	}
}