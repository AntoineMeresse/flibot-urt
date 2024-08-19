package api

import (
	"encoding/json"
	"fmt"
	"io"
)


type MapSearchResult struct {
	Matching []string `json:"matching"`
}

func (api *Api) GetMapsWithPattern(criteria string) []string {
	url := fmt.Sprintf("%s/maps/download/%s", api.BridgeLocalUrl, criteria)
	resp, err := api.Client.Get(url)

	if err == nil {
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res MapSearchResult;
			if err := json.Unmarshal(body, &res); err == nil {
				// logrus.Debugf("GetMapWithPattern (%s): %v", url, res)
				return res.Matching
			} 
		} 
	} 
	return []string{""}
}