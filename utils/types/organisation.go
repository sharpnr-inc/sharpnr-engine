package types

type Organization struct {
	UniqueID       string   `json:"id"`              // Unique identifier for the organisation
	TenentID       string   `json:"tenant_id"`       // Tenant ID for multi-tenancy support
	Name           string   `json:"name"`            // The name of the organisation
	Email          string   `json:"email"`           // The official email address of the organisation
	ContactNumber  string   `json:"phone"`           // The official phone number of the organisation
	Address        string   `json:"address"`         // The physical address of the organisation
	Description    string   `json:"description"`     // A brief description of the organisation
	Website        string   `json:"website"`         // The official website of the organisation
	LogoURL        string   `json:"logo_url"`        // URL to the logo of the organisation
	Location       string   `json:"location"`        // The physical location of the organisation
	FoundedYear    int      `json:"founded_year"`    // The year the organisation was founded
	Employees      []string `json:"members"`         // List of members in the organisation
	ContactInfo    string   `json:"contact_info"`    // Contact information for the organisation
	Departments    []string `json:"departments"`     // List of departments within the organisation
	PaymentDetails string   `json:"payment_details"` // Payment details for the organisation
	BanksInfo      []string `json:"banks_info"`      // List of bank information associated with the organisation

}