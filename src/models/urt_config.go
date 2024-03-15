package models

import (
	"fmt"
	"os"
	"strings"
)

type UrtConfig struct {
	BasePath string
	DownloadPath string
	GotosPath string
	MapRepository string
}

func (u *UrtConfig) init() {
	u.BasePath = os.Getenv("urtPath")
	if u.BasePath != "" {
		path := strings.TrimSuffix(u.BasePath, "/")
		u.DownloadPath = fmt.Sprintf("%s/%s", path, "q3ut4/download")
		u.GotosPath = fmt.Sprintf("%s/%s", path, "q3ut4/gotos")
	}
	u.MapRepository = os.Getenv("urtRepo")
}