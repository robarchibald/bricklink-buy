package bricklinkuser

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// ManageWantedListReturn is the return information provided after managing a wanted list
type ManageWantedListReturn struct {
	WantedListID   int    `json:"wantedMoreID"`
	ReturnCode     int    `json:"returnCode"`
	ReturnMessage  string `json:"returnMessage"`
	ErrorTicket    int    `json:"errorTicket"`
	ProcessingTime int    `json:"procssingTime"`
}

// CreateWantedList is used to create a Bricklink wanted list
func (c *Client) CreateWantedList(name, description string) (int, error) {
	resp, err := c.client.PostForm(cloneBase+"/wanted/editList.ajax", getManageWantedListQuery(name, description, "C"))
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	r := &ManageWantedListReturn{}
	if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
		return -1, err
	}
	if r.ReturnCode != 0 {
		return -1, fmt.Errorf("Error creating wanted list: %s", r.ReturnMessage)
	}
	return r.WantedListID, nil
}

func getManageWantedListQuery(name, description, action string) url.Values {
	values := url.Values{}
	values.Add("wantedMoreName", name)
	values.Add("description", description)
	values.Add("action", action)
	return values
}