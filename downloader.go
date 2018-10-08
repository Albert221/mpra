package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	ApiTimeout             = 10 * time.Second
	ApiEndpoint            = "https://api.dane.gov.pl/resources/1851"
	ApiDownloadUrlJsonPath = "data.attributes.file_url"
	DownloadsDir           = "downloads"
	DownloadTimeout        = 20 * time.Minute
)

type Downloader struct {
	TargetFilename string
}

func (d *Downloader) ScheduleDownloads(frequency time.Duration, updated chan bool) {
	for {
		if err := d.Download(); err == nil {
			updated <- true
		} else {
			log.Println(err)
		}
		time.Sleep(frequency)
	}
}

func (d *Downloader) Download() error {
	d.ensureDownloadDirectory()

	downloadUrl, err := d.fetchDownloadUrl()
	if err != nil {
		return err
	}

	return d.downloadFile(downloadUrl)
}

func (d *Downloader) ensureDownloadDirectory() {
	os.Mkdir(DownloadsDir, 0755)
}

func (d *Downloader) fetchDownloadUrl() (string, error) {
	client := &http.Client{Timeout: ApiTimeout}

	resp, err := client.Get(ApiEndpoint)
	if err != nil {
		return "", err
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	val := gjson.Get(string(response), ApiDownloadUrlJsonPath)

	return val.String(), nil
}

func (d *Downloader) downloadFile(url string) error {
	client := &http.Client{Timeout: DownloadTimeout}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tempPath := fmt.Sprintf("%s/temp_%s", DownloadsDir, d.TargetFilename)
	file, err := os.Create(tempPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	file.Close()

	return os.Rename(tempPath, fmt.Sprintf("%s/%s", DownloadsDir, d.TargetFilename))
}
