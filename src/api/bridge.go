package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type MapSearchResult struct {
	Matching []string `json:"matching"`
}

func (api *Api) GetMapsWithPattern(criteria string) []string {
	url := fmt.Sprintf("%s/maps/download/%s", api.BridgeLocalUrl, criteria)
	resp, err := api.Client.Get(url)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []string{}
	}

	var res MapSearchResult
	if err := json.Unmarshal(body, &res); err != nil {
		return []string{}
	}

	return res.Matching
}

func (api *Api) MapSync() error {
	log.Debug("MapSync called.")
	// TODO: Implement mapsync via Bridge
	return fmt.Errorf("mapsync method not implemented yet")
}

type BridgePlayer struct {
	Name    string `json:"name"`
	Ingame  bool   `json:"ingame"`
	Running bool   `json:"running"`
}

func (api *Api) SendServerInfo(mapname string, players []BridgePlayer) error {
	url := fmt.Sprintf("%s/server", api.BridgeLocalUrl)
	payload, err := json.Marshal(map[string]any{
		"serverAddress": api.ServerUrl,
		"mapname":       mapname,
		"playersList":   players,
	})
	if err != nil {
		return err
	}
	resp, err := api.Client.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bridge returned status %d", resp.StatusCode)
	}
	return nil
}

func (api *Api) SendGlobalMessage(playerName, message string) error {
	url := fmt.Sprintf("%s/message/all", api.BridgeUrl)
	payload, err := json.Marshal(map[string]string{
		"message":       message,
		"serverAddress": api.ServerUrl,
		"name":          playerName,
		"apikey":        api.BridgeApiKey,
	})
	if err != nil {
		return err
	}
	resp, err := api.Client.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bridge returned status %d", resp.StatusCode)
	}
	return nil
}

type ServersListStatus []map[string]ServerStatus

type ServerStatus struct {
	Mapname   string   `json:"mapname"`
	NbPlayers int      `json:"nbPlayers"`
	Ingame    []string `json:"ingame"`
	Spec      []string `json:"spec"`
}

func (api *Api) GetServerStatus() (ServersListStatus, error) {
	url := fmt.Sprintf("%s/status", api.BridgeLocalUrl)
	resp, err := api.Client.Get(url)
	if err != nil {
		return ServersListStatus{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ServersListStatus{}, err
	}

	log.Debugf("Status: %s", string(body))

	var res ServersListStatus
	if err := json.Unmarshal(body, &res); err != nil {
		log.Error(err.Error())
		return ServersListStatus{}, err
	}

	log.Debugf("GetServerStatus (%s): %v", url, res)
	return res, nil
}
