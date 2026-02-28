package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type GeoIPResult struct {
	Country    string `json:"country"`
	RegionName string `json:"regionName"`
	Timezone   string `json:"timezone"`
	Status     string `json:"status"`
}

func (api *Api) GetGeoIP(ip string) (GeoIPResult, error) {
	log.Debugf("[GetGeoIP] querying ip: %s", ip)
	body, _ := json.Marshal([]map[string]string{{"query": ip, "fields": "status,country,regionName,timezone"}})
	resp, err := http.Post(
		"http://ip-api.com/batch",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return GeoIPResult{}, err
	}
	defer resp.Body.Close()
	log.Debugf("[GetGeoIP] status: %d", resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return GeoIPResult{}, err
	}
	log.Debugf("[GetGeoIP] body: %s", string(data))

	var results []GeoIPResult
	if err := json.Unmarshal(data, &results); err != nil {
		return GeoIPResult{}, err
	}

	if len(results) == 0 {
		return GeoIPResult{}, fmt.Errorf("empty response")
	}

	return results[0], nil
}
