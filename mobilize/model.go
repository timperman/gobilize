package mobilize

type Event struct {
	ID                     int64
	Title                  string
	Summary                string
	Description            string
	FeaturedImageURL       string `json:"featured_image_url"`
	HighPriority           bool   `json:"high_priority"`
	Sponsor                Organization
	Timeslots              []Timeslot
	Location               Location
	Timezone               string
	EventType              string `json:"event_type"`
	BrowserURL             string `json:"browser_url"`
	CreatedDate            Time   `json:"created_date"`
	ModifiedDate           Time   `json:"modified_date"`
	Visibility             string
	AddressVisibility      string `json:"address_visibility"`
	CreatedByVolunteerHost bool   `json:"created_by_volunteer_host"`
	VirtualActionURL       string `json:"virtual_action_url"`
	Contact                Contact
	AccessibilityStatus    string `json:"accessibility_status"`
	AccessibilityNotes     string `json:"accessibility_notes"`
	Tags                   []Tag
	EventCampaign          EventCampaign `json:"event_campaign"`
}

type Organization struct {
	ID                int64
	Name              string
	Slug              string
	IsCoordinated     bool   `json:"is_coordinated"`
	IsIndependent     bool   `json:"is_independent"`
	RaceType          string `json:"race_type"`
	IsPrimaryCampaign bool   `json:"is_primary_campaign"`
	State             string
	District          string
	CandidateName     string `json:"candidate_name"`
	OrgType           string `json:"org_type"`
}

type Timeslot struct {
	ID           int64
	StartDate    Time `json:"start_date"`
	EndDate      Time `json:"end_date"`
	MaxAttendees int  `json:"max_attendees"`
	IsFull       bool `json:"is_full"`
}

type Location struct {
	Venue                 string
	AddressLines          []string `json:"address_lines"`
	Locality              string
	Region                string
	Country               string
	PostalCode            string `json:"postal_code"`
	Location              LatLong
	CongressionalDistrict string `json:"congressional_district"`
	StateLegDistrict      string `json:"state_leg_district"`
	StateSenateDistrict   string `json:"state_senate_district"`
}

type LatLong struct {
	Latitude  float64
	Longitude float64
}

type Contact struct {
	Name         string
	EmailAddress string `json:"email_address"`
	PhoneNumber  string `json:"phone_number"`
	OwnerUserID  int64  `json:"owner_user_id"`
}

type Tag struct {
	ID   int64
	Name string
}

type EventCampaign struct {
	ID                 int64
	Slug               string
	EventCreatePageURL string `json:"event_create_page_url"`
}
