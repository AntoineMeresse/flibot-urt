package models

import (
	"fmt"
	"os"
	"strings"
)

type UrtPath struct {
	BasePath string
	DownloadPath string
	GotosPath string
}

func (urtPath *UrtPath) init() {
	urtPath.BasePath = os.Getenv("urtPath")
	if urtPath.BasePath != "" {
		path := strings.TrimSuffix(urtPath.BasePath, "/")
		urtPath.DownloadPath = fmt.Sprintf("%s/%s", path, "q3ut4/download")
		urtPath.GotosPath = fmt.Sprintf("%s/%s", path, "q3ut4/gotos")
	}
}