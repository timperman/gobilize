package handle

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/timperman/gobilize/mobilize"
)

var eventsTemplate = template.Must(template.ParseFiles("templates/event_list.html"))

func RenderEvents(w http.ResponseWriter, r *http.Request) {
	listReq := mobilize.ListEventsRequest{
		OrganizationID: 1767,
	}

	values := r.URL.Query()
	if org := values.Get("org_id"); org != "" {
		if orgID, err := strconv.ParseInt(org, 10, 0); err == nil {
			listReq.OrganizationID = int(orgID)
		}
	}
	days := values.Get("days")
	if days == "" {
		days = "7"
	}
	if numDays, err := strconv.ParseInt(days, 10, 0); err == nil {
		t := time.Now().Add(time.Duration(numDays) * 24 * time.Hour)
		listReq.TimeslotEnd = fmt.Sprintf("lte_%d", t.Unix())
	}
	if zip := values.Get("zip"); zip != "" {
		listReq.ZipCode = zip
	} else {
		listReq.ZipCode = "45140"
	}
	if maxDist := values.Get("max_dist"); maxDist != "" {
		if max, err := strconv.ParseInt(maxDist, 10, 0); err == nil {
			listReq.MaxDistance = int(max)
		}
	}

	events, err := mobilize.ListEventsByDate(listReq)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	if len(events) == 0 {
		w.WriteHeader(404)
		w.Write([]byte("No events found"))
		return
	}

	err = eventsTemplate.ExecuteTemplate(w, "event_list.html", events)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "text/html")
}
