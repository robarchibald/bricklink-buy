package bricklinkuser

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/andrewarchi/brick-apis/credentials"
)

func TestCreateWantedList(t *testing.T) {
	t.SkipNow()
	c, err := NewClient(&credentials.BrickLinkUser{Username: username, Password: password})
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Login(); err != nil {
		t.Fatal(err)
	}
	id, err := c.CreateWantedList("test2", "test2 description")
	if err != nil {
		t.Error(err)
	}
	if id == -1 {
		t.Error("Expected valid new id")
	}
}

func TestDecodeWantedListReturn(t *testing.T) {
	jsonText := `{"wantedMoreID":2711505,"returnCode":5,"returnMessage":"OK","errorTicket":9,"procssingTime":75}`
	r := &ManageWantedListReturn{}
	if err := json.NewDecoder(strings.NewReader(jsonText)).Decode(r); err != nil {
		t.Fatal(err)
	}
	if r.WantedListID != 2711505 || r.ReturnCode != 5 || r.ReturnMessage != "OK" || r.ErrorTicket != 9 || r.ProcessingTime != 75 {
		t.Error("Expected correct values", r)
	}
}
