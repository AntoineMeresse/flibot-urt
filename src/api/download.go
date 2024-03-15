package api

import (
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func DownloadFile(filepath string, url string) error {
	log.Debugf("Trying to download (%s) from (%s)", filepath, url)

	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("Can't download: %s", url)
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		log.Errorf("Can't download: %s", url)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}