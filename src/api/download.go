package api

import (
	"io"
	"log/slog"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string) (bytes int64, err error) {
	slog.Debug("Trying to download", "filepath", filepath, "url", url)

	resp, err := http.Get(url)
	if err != nil {
		slog.Error("Can't download", "url", url, "err", err)
		return 0, err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		slog.Error("Can't create file", "filepath", filepath, "err", err)
		return 0, err
	}
	defer out.Close()

	return io.Copy(out, resp.Body)
}
