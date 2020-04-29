package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Albert221/mpra/graphql"
	"github.com/Albert221/mpra/puller"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/markbates/pkger"
)

var (
	host          string
	port          string
	datasetPath   string
	refreshPeriod time.Duration
)

func init() {
	flag.StringVar(&host, "host", "", "hostname to listen on")
	flag.StringVar(&port, "port", "8080", "port to listen on")
	flag.StringVar(&datasetPath, "dataset", "dataset.xml", "path to the file which the dataset will be cached to")
	flag.DurationVar(&refreshPeriod, "refresh", 1*time.Hour, "period every which the dataset will be refreshed")
}

func main() {
	flag.Parse()

	q := graphql.NewQuery()
	schema, err := q.CreateSchema()
	if err != nil {
		log.Fatal(err)
	}

	datasetPuller := puller.NewDatasetPuller(datasetPath, refreshPeriod, q)
	go (func() {
		if err := datasetPuller.Run(); err != nil {
			log.Fatal(err)
		}
	})()

	// GraphQL Playground
	http.Handle("/", http.FileServer(pkger.Dir("/public")))
	http.Handle("/graphql", &relay.Handler{Schema: schema})

	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("Listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
