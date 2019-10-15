package ipinfo

// For full documentation, please see ipinfo's documentation at:
// https://ipinfo.io/developers/responses#full-response

// FullResponse is the full response that can be obtained from the authorized API.
type FullResponse struct {
	Location
	Hostname string `json:"hostname"`
	// Org will only be included on a free or basic plan.
	// This combines the AS number + name.
	Org string `json:"org"`
	Asn ASN    `json:"asn"`
	// Company will only be included if on a Pro plan.
	Company Company `json:"company"`
	// Carrier will only be included if on a Pro plan, and if the API determines
	// that the IP is used for mainly mobile carrier traffic.
	Carrier Carrier `json:"carrier"`
}

// Location represents the geolocation info about an IP.
// This will be included in most requests as an embedded struct field.
type Location struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

// ASN represents the AS number.
// Will only be included if on a Pro plan.
type ASN struct {
	Asn    string `json:"asn"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Route  string `json:"route"`
	Type   string `json:"type"`
}

// Company represents the company or organization that actually uses the IP address.
// The type field will be one of "business", "education", "hosting" or "isp"
// Will only be included if on a Pro plan.
type Company struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
}

// Carrier represents the carrier operating the IP, and will only be included if
// there is a high liklihood that this IP address is used exclusively for mobile
// carrier traffic.
// Will only be included if on a Pro plan.
type Carrier struct {
	Name string `json:"name"`
	Mcc  string `json:"mcc"`
	Mnc  string `json:"mnc"`
}

// Response is the response obtained from the unauthorized API.
type Response struct {
	Location
	Hostname string `json:"hostname"`
	// Org is a string consisting of the combined AS number and name.
	Org string `json:"org"`
	// Readme will always be a link to ipinfo's marketing page telling you to
	// get the paid or otherwise registered version of the API.
	// https://ipinfo.io/missingauth
	Readme string `json:"readme"`
}
