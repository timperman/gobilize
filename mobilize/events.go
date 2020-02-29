package mobilize

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ListEventsRequest struct {
	OrganizationID  int
	TimeslotStart   string
	TimeslotEnd     string
	ZipCode         string
	GroupDateFormat string
}

type ListEventsResponse struct {
	Count    int
	Next     string
	Previous string
	Data     []Event
}

func ListEventsByDate(req ListEventsRequest) (map[string][]Event, error) {
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

	timeFormat := req.GroupDateFormat
	if timeFormat == "" {
		timeFormat = "Monday, January 2, 2006"
	}

	events := make(map[string][]Event, 0)
	next := listURL.String()
	for next != "" {
		response, err := http.Get(next)
		if err != nil {
			return events, err
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return events, err
		}

		var listResponse ListEventsResponse
		json.Unmarshal(data, &listResponse)

		for _, e := range listResponse.Data {
			date := e.Timeslots[0].StartDate.Time().Format(timeFormat)
			if dateEvents, ok := events[date]; ok {
				events[date] = append(dateEvents, e)
			} else {
				events[date] = []Event{e}
			}
		}

		next = listResponse.Next
	}
	return events, nil
}
