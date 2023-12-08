package models

import (
	"fmt"
	"os"
	"strings"
)

type UrtPath struct {
	basePath string
	downloadPath string
	gotosPath string
}

func (urtPath *UrtPath) init() {
	urtPath.basePath = os.Getenv("urtPath")
	if urtPath.basePath != "" {
		path := strings.TrimSuffix(urtPath.basePath, "/")
		urtPath.downloadPath = fmt.Sprintf("%s/%s", path, "q3ut4/download")
		urtPath.gotosPath = fmt.Sprintf("%s/%s", path, "q3ut4/gotos")
	}
}