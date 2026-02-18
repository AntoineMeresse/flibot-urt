package api

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
)

func (api *Api) SendFileToWebhook(filePath string, message string) error {
	path := "/demos/" + filePath
	slog.Debug("Sending file to webhook", "webhook", api.DiscordWebhook, "file", filePath, "path", path, "message", message)

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
