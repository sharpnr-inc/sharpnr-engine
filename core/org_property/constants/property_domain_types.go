package property_constants

type PropertyDomainTypes string

const (
	Hospitality    PropertyDomainTypes = "HOSPITALITY"
	Education      PropertyDomainTypes = "EDUCATION"
	MedicalScience PropertyDomainTypes = "MEDICAL_SCIENCE"
)

type PropertyDomain struct {
	ID          string
	Name        PropertyDomainTypes
	Description string
	IsAvailable bool
}

var PropertyDomains = []PropertyDomain{
	{
		ID:   "23dr9Q",
		Name: Hospitality,
		Description: "Hospitality domain is designed for managing hotels, restaurants, bars, cafés, pubs, clubs, and cloud kitchens. " +
			"It provides end-to-end operational capabilities including reservations, table and floor management, menu configuration, " +
			"order processing, kitchen and service point coordination, inventory tracking, staff operations, billing, and real-time reporting. " +
			"This domain supports both dine-in and takeaway workflows and is optimized for high-volume, service-driven environments.",
		IsAvailable: true,
	},
	{
		ID:   "EDU-001",
		Name: Education,
		Description: "Education domain is intended for managing academic institutions such as schools, colleges, universities, and training centers. " +
			"It will include features like student enrollment, course and curriculum management, attendance tracking, examinations, grading, " +
			"faculty administration, scheduling, and academic reporting. This domain is currently under planning and not yet available.",
		IsAvailable: false,
	},
	{
		ID:   "MED-001",
		Name: MedicalScience,
		Description: "Medical Science domain is planned for healthcare and clinical operations including hospitals, clinics, diagnostic centers, " +
			"and laboratories. It will cover patient management, appointments, electronic medical records (EMR), billing, insurance workflows, " +
			"staff scheduling, compliance, and analytics. This domain is under design and not available at this time.",
		IsAvailable: false,
	},
}
