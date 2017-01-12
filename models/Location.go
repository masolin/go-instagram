package models

import (
	"encoding/json"
	"strconv"

	"github.com/masolin/go-instagram/constants"
	"github.com/masolin/go-instagram/utils"
)

type Location struct {
	ExternalSource   string                `json:"external_source"`
	City             string                `json:"city"`
	Name             string                `json:"name"`
	FacebookPlacesID int64                 `json:"facebook_places_id"`
	Address          string                `json:"address"`
	Lat              float64               `json:"lat"`
	Lng              float64               `json:"lng"`
	Pk               int64                 `json:"pk"`
	AgentPool        *utils.SuperAgentPool `json:"-"`
}

type mediaResponse struct {
	RankedItems []*Media `json:"ranked_items"`
	Items       []*Media `json:"items"`
}

func (location Location) GetRankedMedia() []*Media {
	url := constants.ROUTES.LocationFeed + strconv.FormatInt(location.Pk, 10) + "/"

	agent := location.AgentPool.Get()
	defer location.AgentPool.Put(agent)

	_, body, _ := utils.WrapRequest(
		agent.Get(url).Query("rank_token=" + utils.GenerateUUID()))

	var resp mediaResponse
	json.Unmarshal([]byte(body), &resp)

	return resp.RankedItems
}

func (location Location) GetRecentMedia() []*Media {
	url := constants.ROUTES.LocationFeed + strconv.FormatInt(location.Pk, 10) + "/"

	agent := location.AgentPool.Get()
	defer location.AgentPool.Put(agent)

	_, body, _ := utils.WrapRequest(
		agent.Get(url).Query("rank_token=" + utils.GenerateUUID()))

	var resp mediaResponse
	json.Unmarshal([]byte(body), &resp)

	return resp.Items
}
