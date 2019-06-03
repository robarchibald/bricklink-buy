package bricklinkuser

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestGetError(t *testing.T) {
	jsonText := `{"returnCode":-3,"returnMessage":"Invalid Parameter!","errorTicket":0,"procssingTime":0}`
	r := &http.Response{Body: ioutil.NopCloser(strings.NewReader(jsonText))}
	if err := getError(r, nil); err != nil && err.Error() != "Invalid Parameter!" {
		t.Fatal(err)
	}
}

func TestGetWantedList(t *testing.T) {
	t.SkipNow()
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Login(username, password); err != nil {
		t.Fatal(err)
	}
	t.Error(c.GetWantedList(2005021))
}
