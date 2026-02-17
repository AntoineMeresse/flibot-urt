package actionslist

import (
	"encoding/json"
	"strings"

	appcontext "github.com/AntoineMeresse/flibot-urt/src/context"
	log "github.com/sirupsen/logrus"
)

type savedPositionInfo struct {
	Mapname  string  `json:"mapname"`
	Location string  `json:"location"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Z        float64 `json:"z"`
	Angle    float64 `json:"angle"`
}

func SavedPosition(actionParams []string, c *appcontext.AppContext) {
	posJson := strings.Join(actionParams, "")
	posJson = strings.Replace(posJson, "'", "\"", -1)
	log.Debugf("SavedPosition: %v", posJson)

	var posInfo savedPositionInfo
	if err := json.Unmarshal([]byte(posJson), &posInfo); err != nil {
		log.Errorf("SavedPosition: Error unmarshalling json: %v", err)
		return
	}

	if err := c.DB.PositionSave(posInfo.Mapname, posInfo.Location, posInfo.X, posInfo.Y, posInfo.Z, posInfo.Angle); err != nil {
		log.Errorf("SavedPosition: Error saving position to db: %v", err)
	}
}
