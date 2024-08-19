package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

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
	logrus.Debugf("GetMapInformation. Url: %s, mapname: %s", api.UjmUrl, mapname)
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
				logrus.Debugf("GetMapInformation (%s): %v", url, res)
				return res, nil
			} 
		} 
	} 

	return MapInfos{}, err
}