package mobilize

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ListEventsRequest struct {
	OrganizationID int
	TimeslotStart  string
	TimeslotEnd    string
	ZipCode        string
}

type ListEventsResponse struct {
	Count    int
	Next     string
	Previous string
	Data     []Event
}

func ListEvents(req ListEventsRequest) ([]Event, error) {
	listURL, _ := url.Parse("https://api.mobilize.us/v1/events")

	params := url.Values{}
	if req.OrganizationID > 0 {
		params.Add("organization_id", fmt.Sprintf("%d", req.OrganizationID))
	}
	if req.TimeslotStart != "" {
		params.Add("timeslot_start", req.TimeslotStart)
	} else {
		params.Add("timeslot_start", "gte_now")
	}
	if req.TimeslotEnd != "" {
		params.Add("timeslot_end", req.TimeslotEnd)
	}
	if req.ZipCode != "" {
		params.Add("zipcode", req.ZipCode)
	}
	listURL.RawQuery = params.Encode()

	response, err := http.Get(listURL.String())
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var listResponse ListEventsResponse
	json.Unmarshal(data, &listResponse)

	return listResponse.Data, nil
}
