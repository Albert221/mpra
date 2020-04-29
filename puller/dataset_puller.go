package puller

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Albert221/mpra/domain"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

const (
	apiEndpoint            = "https://api.dane.gov.pl/resources/23421"
	apiDownloadUrlJsonPath = "data.attributes.link"
)

type DatasetPuller struct {
	cachedDatasetPath string
	refresh           time.Duration
	populator         domain.Populator
}

func NewDatasetPuller(path string, refresh time.Duration, populator domain.Populator) *DatasetPuller {
	return &DatasetPuller{
		populator:         populator,
		refresh:           refresh,
		cachedDatasetPath: path,
	}
}

func (d *DatasetPuller) Run() error {
	dir := filepath.Dir(d.cachedDatasetPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.Wrap(err, "creating directories for cached dataset file")
	}

	f, err := os.Open(d.cachedDatasetPath)
	if err == nil {
		data, _ := ioutil.ReadAll(f)
		d.populate(data)
	}
	f.Close()

	for {
		url, err := d.fetchDownloadUrl()
		if err != nil {
			log.Println(err)
		}

		client := &http.Client{Timeout: 1 * time.Hour}
		resp, err := client.Get(url)
		if err != nil {
			return errors.Wrap(err, "fetching dataset from api")
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}
		if err := d.saveCachedDataset(data); err != nil {
			log.Println(err)
			continue
		}
		if err := d.populate(data); err != nil {
			log.Println(err)
			continue
		}

		resp.Body.Close()
		time.Sleep(d.refresh)
	}
}

func (d *DatasetPuller) populate(data []byte) error {
	products, err := domain.UnmarshallProducts(data)
	if err != nil {
		return errors.Wrap(err, "unmarshalling products")
	}

	d.populator.Populate(products)
	return nil
}

func (d *DatasetPuller) saveCachedDataset(data []byte) error {
	f, err := os.Create(d.cachedDatasetPath)
	if err != nil {
		return errors.Wrap(err, "creating cached dataset file")
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return errors.Wrap(err, "writing dataset to cache file")
	}

	return nil
}

func (d *DatasetPuller) fetchDownloadUrl() (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(apiEndpoint)
	if err != nil {
		return "", errors.Wrap(err, "fetching api endpoint")
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "reading api endpoint response")
	}

	return gjson.Get(string(response), apiDownloadUrlJsonPath).String(), nil
}
