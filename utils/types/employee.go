package types

type Employee struct {
	Name        string   `json:"name"`         // The name of the employee
	Position    string   `json:"position"`     // The position or job title of the employee
	Email       string   `json:"email"`        // The email address of the employee
	Phone       string   `json:"phone"`        // The phone number of the employee
	Department  string   `json:"department"`   // The department where the employee works
	JoiningDate string   `json:"joining_date"` // The date when the employee joined the organisation
	Salary      int      `json:"salary"`       // The salary of the employee
	Address     string   `json:"address"`      // The residential address of the employee
	Skills      []string `json:"skills"`       // List of skills possessed by the employee
	Projects    []string `json:"projects"`     // List of projects the employee is involved in
}