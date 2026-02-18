package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/AntoineMeresse/flibot-urt/src/models"
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
	slog.Debug("GetMapInformation", "url", api.UjmUrl, "mapname", mapname)
	url := fmt.Sprintf("%s/mapinfo/requestdata", api.UjmUrl)
	postBody, _ := json.Marshal(map[string]interface{}{
		"mapname": mapname,
		"apikey":  api.Apikey,
	})

	resp, err := api.Client.Post(url, "application/json", bytes.NewBuffer(postBody))

	if err == nil {
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res MapInfos
			if err := json.Unmarshal(body, &res); err == nil {
				slog.Debug("GetMapInformation result", "url", url, "result", res)
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
	slog.Debug("GetToprunsInformation", "url", api.UjmUrl, "mapname", mapname)
	url := fmt.Sprintf("%s/runs/requestdata", api.UjmUrl)
	postBody, _ := json.Marshal(map[string]interface{}{
		"mapname": mapname,
		"apikey":  api.Apikey,
	})

	resp, err := api.Client.Post(url, "application/json", bytes.NewBuffer(postBody))

	if err == nil {
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res ToprunsInfos
			if err := json.Unmarshal(body, &res); err == nil {
				slog.Debug("GetToprunsInformation result", "url", url, "result", res)
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
	slog.Debug("GetLatestRuns", "url", url)

	getBody, _ := json.Marshal(map[string]interface{}{
		"apikey": api.Apikey,
	})

	resp, err := api.UjmGetWithBody(url, bytes.NewBuffer(getBody))

	if err == nil {
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res []LatestRunElement
			if err := json.Unmarshal(body, &res); err == nil {
				slog.Debug("GetLatestRuns result", "url", url, "result", res)
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
	slog.Debug("GetLatestMaps", "url", url)

	getBody, _ := json.Marshal(map[string]interface{}{
		"apikey": api.Apikey,
	})

	resp, err := api.UjmGetWithBody(url, bytes.NewBuffer(getBody))

	if err == nil {
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res []LatestMapElement
			if err := json.Unmarshal(body, &res); err == nil {
				slog.Debug("GetLatestMaps result", "url", url, "result", res)
				return res, nil
			} else {
				return []LatestMapElement{}, err
			}
		}
	}

	return []LatestMapElement{}, err
}

type PersonalBestInfos struct {
	Id        int                   `json:"mapid"`
	Filename  string                `json:"mapfilename"`
	Mapname   string                `json:"mapname"`
	RunsInfos []PersonalBestElement `json:"runs"`
}

type PersonalBestElement struct {
	Run     RunPlayerInfos `json:"run"`
	Rank    string         `json:"rank"`
	Wayname string         `json:"wayname"`
}

func (api *Api) GetPersonalBestInformation(mapname string, guid string) (PersonalBestInfos, error) {
	slog.Debug("GetPersonalBestInformation", "url", api.UjmUrl, "mapname", mapname)
	url := fmt.Sprintf("%s/runs/getpb", api.UjmUrl)
	postBody, _ := json.Marshal(map[string]interface{}{
		"mapname":    mapname,
		"apikey":     api.Apikey,
		"playerguid": guid,
	})

	resp, err := api.UjmGetWithBody(url, bytes.NewBuffer(postBody))

	if err == nil {
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil {
			var res PersonalBestInfos
			if err := json.Unmarshal(body, &res); err == nil {
				slog.Debug("GetPersonalBestInformation result", "url", url, "result", res)
				return res, nil
			} else {
				return PersonalBestInfos{}, err
			}
		}
	}

	return PersonalBestInfos{}, err
}

type DemoBody struct {
	Playerguid  string `json:"playerguid"`
	Playername  string `json:"playername"`
	Serverip    string `json:"serverip"`
	Servername  string `json:"servername"`
	Serverfps   string `json:"serverfps"`
	Runtime     string `json:"runtime"`
	Mapfilename string `json:"mapfilename"`
	Waynumber   string `json:"waynumber"`
	Apikey      string `json:"apikey"`
	PlayerIp    string `json:"playerip"`
}

type SendDemoResponse struct {
	Added        int     `json:"added"`
	Improvement  string  `json:"improvement"`
	Wrdifference string  `json:"wrdifference"`
	Rank         *string `json:"rank"`
	Process      bool
}

func (api *Api) PostRunDemo(p models.PlayerRunInfo, demoDirectory string) (SendDemoResponse, error) {
	slog.Debug("PostRunDemo")

	d := &DemoBody{
		Playerguid:  p.Guid,
		Playername:  p.Playername,
		Serverip:    api.ServerUrl,
		Servername:  p.ServerName,
		Serverfps:   p.Fps,
		Runtime:     p.Time,
		Mapfilename: p.Mapname,
		Waynumber:   p.Way,
		Apikey:      api.Apikey,
		PlayerIp:    p.PlayerIp,
	}

	demoResponse, err := api.postRunWithDemo(d, p.GetDemoName(), demoDirectory)

	if err != nil {
		slog.Error("PostRunDemo: could not upload with demo file, retrying without", "err", err)
		demoResponse, err = api.PostRunWithoutDemo(d)
	}

	return demoResponse, err
}

func (api *Api) postRunWithDemo(demoBody *DemoBody, demoName string, demoDirectory string) (SendDemoResponse, error) {
	url := fmt.Sprintf("%s/runs/addrunwithdemo", api.UjmUrl)

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// File: "json"
	jsonBytes, err := json.Marshal(*demoBody)
	if err != nil {
		return SendDemoResponse{}, err
	}

	jsonPart, err := writer.CreateFormFile("json", "json")
	if err != nil {
		return SendDemoResponse{}, err
	}
	_, err = jsonPart.Write(jsonBytes)

	// File: "upload_file"
	path := fmt.Sprintf("%s/%s", demoDirectory, demoName)
	b, err := os.ReadFile(path)

	if err != nil {
		return SendDemoResponse{}, err
	}

	filePart, err := writer.CreateFormFile("upload_file", demoName)
	if err != nil {
		return SendDemoResponse{}, err
	}

	_, err = io.Copy(filePart, bytes.NewReader(b))
	if err != nil {
		return SendDemoResponse{}, err
	}

	err = writer.Close()
	if err != nil {
		return SendDemoResponse{}, err
	}

	resp, err := api.Client.Post(url, writer.FormDataContentType(), &requestBody)

	return handlePostDemoResponse(err, resp, url, "PostRunWithDemo")
}

func (api *Api) PostRunWithoutDemo(demoBody *DemoBody) (SendDemoResponse, error) {
	url := fmt.Sprintf("%s/runs/addrun", api.UjmUrl)

	j, err := json.Marshal(*demoBody)
	if err != nil {
		slog.Error("PostRunWithoutDemo: json marshal error", "err", err)
		return SendDemoResponse{}, fmt.Errorf("[PostRunWithoutDemo] Json marshal error: %w", err)
	}

	resp, err := api.Client.Post(url, "application/json", bytes.NewBuffer(j))

	return handlePostDemoResponse(err, resp, url, "PostRunWithoutDemo")
}

func handlePostDemoResponse(err error, resp *http.Response, url string, functionName string) (SendDemoResponse, error) {
	slog.Debug("handlePostDemoResponse", "function", functionName, "url", url)

	if err == nil {
		defer resp.Body.Close()
		slog.Debug("handlePostDemoResponse status", "function", functionName, "status", resp.StatusCode)
		if resp.StatusCode == 200 {
			if body, err := io.ReadAll(resp.Body); err == nil {
				slog.Debug("handlePostDemoResponse body", "function", functionName, "body", string(body))
				var res SendDemoResponse
				if err := json.Unmarshal(body, &res); err == nil {
					res.Process = true
					return res, nil
				}
			}
		}
		return SendDemoResponse{}, fmt.Errorf("[%s]: Send demo status: %d", functionName, resp.StatusCode)
	}
	return SendDemoResponse{}, fmt.Errorf("[%s] Demo body error: %v", functionName, err)
}
