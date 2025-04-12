package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

type MapInfos struct {
	Id       int      `json:"id"`
	Filename string   `json:"filename"`
	Mapname  string   `json:"mapname"`
	Jumps    string   `json:"jnumber"`
	Level    int      `json:"level"`
	Release  string   `json:"releasedate"`
	Mappers  []string `json:"mappers"`
	Notes    []string `json:"notes"`
	Types    []string `json:"types"`
}

func (api *Api) GetMapInformation(mapname string) (MapInfos, error) {
	logrus.Debugf("[GetMapInformation] Url: %s, mapname: %s", api.UjmUrl, mapname)
	url := fmt.Sprintf("%s/mapinfo/requestdata", api.UjmUrl)
	postBody, _ := json.Marshal(map[string]interface{}{
		"mapname": mapname,
		"apikey":  api.Apikey,
	})

	resp, err := api.Client.Post(url, "application/json", bytes.NewBuffer(postBody))

	if err == nil {
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res MapInfos
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
	Id        int                         `json:"mapid"`
	Filename  string                      `json:"mapfilename"`
	Mapname   string                      `json:"mapname"`
	RunsInfos map[string][]RunPlayerInfos `json:"runs"`
}

type RunPlayerInfos struct {
	PlayerName string `json:"playername"`
	RunDate    string `json:"rundate"`
	RunTime    string `json:"time"`
}

func (api *Api) GetToprunsInformation(mapname string) (ToprunsInfos, error) {
	logrus.Debugf("[GetToprunsInformation] Url: %s, mapname: %s", api.UjmUrl, mapname)
	url := fmt.Sprintf("%s/runs/requestdata", api.UjmUrl)
	postBody, _ := json.Marshal(map[string]interface{}{
		"mapname": mapname,
		"apikey":  api.Apikey,
	})

	resp, err := api.Client.Post(url, "application/json", bytes.NewBuffer(postBody))

	if err == nil {
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res ToprunsInfos
			if err := json.Unmarshal(body, &res); err == nil {
				logrus.Debugf("[GetToprunsInformation] (%s): %v", url, res)
				return res, nil
			} else {
				return ToprunsInfos{}, err
			}
		}
	}

	return ToprunsInfos{}, err
}

type LatestRunElement struct {
	Mapname    string `json:"mapname"`
	PlayerName string `json:"playername"`
	Rank       string `json:"rank"`
	RunDate    string `json:"rundate"`
	RunTime    string `json:"runtime"`
	Way        int    `json:"waynumber"`
}

func (api *Api) GetLatestRuns() ([]LatestRunElement, error) {
	url := fmt.Sprintf("%s/runs/latestruns", api.UjmUrl)
	logrus.Debugf("[GetLatestRuns] Url: %s", url)

	getBody, _ := json.Marshal(map[string]interface{}{
		"apikey": api.Apikey,
	})

	resp, err := api.UjmGetWithBody(url, bytes.NewBuffer(getBody))

	if err == nil {
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res []LatestRunElement
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

type LatestMapElement struct {
	Date     string   `json:"dateadded"`
	Filename string   `json:"filename"`
	Mapname  string   `json:"mapname"`
	Mappers  []string `json:"mapper"`
	Types    []string `json:"types"`
}

func (api *Api) GetLatestMaps() ([]LatestMapElement, error) {
	url := fmt.Sprintf("%s/mapinfo/latestmaps", api.UjmUrl)
	logrus.Debugf("[GetLatestMaps] Url: %s", url)

	getBody, _ := json.Marshal(map[string]interface{}{
		"apikey": api.Apikey,
	})

	resp, err := api.UjmGetWithBody(url, bytes.NewBuffer(getBody))

	if err == nil {
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res []LatestMapElement
			// logrus.Debug(string(body))
			if err := json.Unmarshal(body, &res); err == nil {
				logrus.Tracef("[GetLatestMaps] (%s): %v", url, res)
				return res, nil
			} else {
				return []LatestMapElement{}, err
			}
		}
	}

	return []LatestMapElement{}, err
}

type PersonnalBestInfos struct {
	Id        int                    `json:"mapid"`
	Filename  string                 `json:"mapfilename"`
	Mapname   string                 `json:"mapname"`
	RunsInfos []PersonnalBestElement `json:"runs"`
}

type PersonnalBestElement struct {
	Run     RunPlayerInfos `json:"run"`
	Rank    string         `json:"rank"`
	Wayname string         `json:"wayname"`
}

func (api *Api) GetPersonnalBestInformation(mapname string, guid string) (PersonnalBestInfos, error) {
	logrus.Debugf("[GetPersonnalBestInformation] Url: %s, mapname: %s", api.UjmUrl, mapname)
	url := fmt.Sprintf("%s/runs/getpb", api.UjmUrl)
	postBody, _ := json.Marshal(map[string]interface{}{
		"mapname":    mapname,
		"apikey":     api.Apikey,
		"playerguid": guid,
	})

	resp, err := api.UjmGetWithBody(url, bytes.NewBuffer(postBody))

	if err == nil {
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res PersonnalBestInfos
			if err := json.Unmarshal(body, &res); err == nil {
				logrus.Tracef("[GetPersonnalBestInformation] (%s): %v", url, res)
				return res, nil
			} else {
				return PersonnalBestInfos{}, err
			}
		}
	}

	return PersonnalBestInfos{}, err
}
