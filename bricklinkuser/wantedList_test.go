package bricklinkuser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestCreateWantedList(t *testing.T) {
	t.SkipNow()
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Login(username, password); err != nil {
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

func TestEditWantedListItems(t *testing.T) {
	//t.SkipNow()
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Login(username, password); err != nil {
		t.Fatal(err)
	}
	//wantedMoreID: 2709071
	//wantedItemStr: [{"itemID":42854,"colorID":4,"wantedQty":-1,"wantedQtyFilled":0,"wantedNew":"X","wantedNotify":"N","wantedRemarks":null,"wantedPrice":-1}]

	t.Fatal(c.EditWantedListItems(2709071, []EditWantedItemRequest{}))
	/*if err != nil {
		t.Error(err)
	}
	if id == -1 {
		t.Error("Expected valid new id")
	}*/

}

func (c *Client) TestEditWantedListItems() {
	//https: //www.bricklink.com/ajax/clone/wanted/add.ajax
}

// EditWantedListItems is used to edit the items in a wanted list
func (c *Client) EditWantedListItems(wantedListID int, wantedItems []EditWantedItemRequest) error {
	resp, err := c.client.PostForm(cloneBase+"/wanted/add.ajax", getEditWantedListItemsQuery(wantedListID, wantedItems))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
	/*	r := &ManageWantedListReturn{}
		if err := json.NewDecoder(resp.Body).Decode(r); err != nil {
			return -1, err
		}
		if r.ReturnCode != 0 {
			return -1, fmt.Errorf("Error creating wanted list: %s", r.ReturnMessage)
		}
	*/
	return nil
}

func getEditWantedListItemsQuery(wantedListID int, wantedItems []EditWantedItemRequest) url.Values {
	data, err := json.Marshal(wantedItems)
	if err != nil {
		return nil
	}
	values := url.Values{}
	values.Add("wantedMoreID", strconv.Itoa(wantedListID))
	values.Add("wantedItemStr", string(data))
	return values
}
