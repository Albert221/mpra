package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Albert221/medicinal-products-registry-api/api"
	"github.com/graph-gophers/graphql-go/relay"
)

const (
	graphiqlDir string = "./graphiql"
)

func main() {
	downloader := Downloader{TargetFilename: "dataset.xml"}

	updatedChan := make(chan bool)
	go downloader.ScheduleDownloads(20*time.Minute, updatedChan)

	<-updatedChan
	schema, err := api.NewSchema(updatedChan)
	if err != nil {
		log.Println(err)
	}

	http.Handle("/query", &relay.Handler{Schema: schema.CreateGraphQLSchema()})
	http.Handle("/", http.FileServer(http.Dir(graphiqlDir)))
	if err := http.ListenAndServe(os.Getenv("MPR_ADDR"), nil); err != nil {
		log.Println(err)
	}
}
