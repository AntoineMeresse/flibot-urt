package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type MapInfos struct {
	Id int `json:"id"`
	Filename string `json:"filename"`
	Mapname string `json:"mapname"`
	Jumps string `json:"jnumber"`
	Level int `json:"level"`
	Release string `json:"releasedate"`
	Mappers []string `json:"mappers"`
	Notes []string `json:"notes"`
	Types []string `json:"types"`
}

func (api *Api) GetMapInformation(mapname string) (MapInfos, error) {
	logrus.Debugf("[GetMapInformation] Url: %s, mapname: %s", api.UjmUrl, mapname)
	url := fmt.Sprintf("%s/mapinfo/requestdata", api.UjmUrl)
	postBody, _ := json.Marshal(map[string]interface{}{
		"mapname": mapname,
		"apikey": api.Apikey,
	})

	resp, err := api.Client.Post(url, "application/json", bytes.NewBuffer(postBody))

	if err == nil {
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res MapInfos;
			if err := json.Unmarshal(body, &res); err == nil {
				logrus.Tracef("GetMapInformation (%s): %v", url, res)
				return res, nil
			} else {
				return MapInfos{}, err
			}
		} 
	}

	return MapInfos{}, err
}

type ToprunsInfos struct {
	Id int `json:"mapid"`
	Filename string `json:"mapfilename"`
	Mapname string `json:"mapname"`
	RunsInfos map[string][]RunPlayerInfos `json:"runs"`
}

type RunPlayerInfos struct {
	PlayerName string `json:"playername"`
	RunDate string `json:"rundate"`
	RunTime string `json:"time"`
}

func (api *Api) GetToprunsInformation(mapname string) (ToprunsInfos, error) {
	logrus.Debugf("[GetToprunsInformation] Url: %s, mapname: %s", api.UjmUrl, mapname)
	url := fmt.Sprintf("%s/runs/requestdata", api.UjmUrl)
	postBody, _ := json.Marshal(map[string]interface{}{
		"mapname": mapname,
		"apikey": api.Apikey,
	})

	resp, err := api.Client.Post(url, "application/json", bytes.NewBuffer(postBody))

	if err == nil {
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res ToprunsInfos;
			if err := json.Unmarshal(body, &res); err == nil {
				logrus.Tracef("[GetToprunsInformation] (%s): %v", url, res)
				return res, nil
			} else {
				return ToprunsInfos{}, err
			}
		} 
	}

	return ToprunsInfos{}, err
}
type LatestRunElement struct {
	Mapname string `json:"mapname"`
	PlayerName string `json:"playername"`
	Rank string `json:"rank"`
	RunDate string `json:"rundate"`
	RunTime string `json:"runtime"`
	Way int `json:"waynumber"`
}

func (api *Api) GetLatestRuns() ([]LatestRunElement, error) {
	url := fmt.Sprintf("%s/runs/latestruns", api.UjmUrl)
	logrus.Debugf("[GetLatestRuns] Url: %s", url)
	
	getBody, _ := json.Marshal(map[string]interface{}{
		"apikey": api.Apikey,
	})

	request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(getBody))
	request.Header.Set("Content-Type", "application/json")

	if (err != nil) {
		return []LatestRunElement{}, err
	}

	resp, err := http.DefaultClient.Do(request)

	if err == nil {
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res []LatestRunElement;
			if err := json.Unmarshal(body, &res); err == nil {
				logrus.Tracef("[GetLatestRuns] (%s): %v", url, res)
				return res, nil
			} else {
				return []LatestRunElement{}, err
			}
		} 
	}

	return []LatestRunElement{}, err
}