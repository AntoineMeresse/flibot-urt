package api

import (
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func DownloadFile(filepath string, url string) (bytes int64, err error) {
	log.Debugf("Trying to download (%s) from (%s)", filepath, url)

	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("Can't download: %s", url)
		return 0, err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		log.Errorf("Can't download: %s", url)
		return 0, err
	}
	defer out.Close()

	return io.Copy(out, resp.Body)
}