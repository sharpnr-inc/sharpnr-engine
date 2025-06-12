package types

type Organisation struct {
	Name        string   `json:"name"`         // The name of the organisation
	Description string   `json:"description"`  // A brief description of the organisation
	Website     string   `json:"website"`      // The official website of the organisation
	LogoURL     string   `json:"logo_url"`     // URL to the logo of the organisation
	Location    string   `json:"location"`     // The physical location of the organisation
	FoundedYear int      `json:"founded_year"` // The year the organisation was founded
	Members     []string `json:"members"`      // List of members in the organisation
	ContactInfo string   `json:"contact_info"` // Contact information for the organisation
}