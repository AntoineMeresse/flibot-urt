package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func (api *Api) SendFileToWebhook(demoDir string, filePath string, message string) error {
	path := demoDir + "/" + filePath
	logrus.Debugf("Sending file to webhook: %s | File: %s (path: %s) | Message: %s", api.DiscordWebhook, filePath, path, message)

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", path)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	writer.WriteField("content", message)

	writer.Close()

	resp, err := http.Post(api.DiscordWebhook, writer.FormDataContentType(), &body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("webhook returned status: %d", resp.StatusCode)
	}

	return nil
}
