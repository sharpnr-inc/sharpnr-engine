package sharpnreventbus

// ============================
// 🔷 GLOBAL / CORE EVENTS
// ============================
type CoreEvents struct {
	SystemInitialized   string
	OrganizationCreated string
	OrganizationUpdated string
	OrganizationDeleted string
}

var Core = CoreEvents{
	SystemInitialized:   "Core.System.Initialized",
	OrganizationCreated: "Core.Organization.Created",
	OrganizationUpdated: "Core.Organization.Updated",
	OrganizationDeleted: "Core.Organization.Deleted",
}

// ============================
// 🏨 HOSPITALITY EVENTS
// ============================
type HospitalityEvents struct {
	OrderCreated   string
	OrderUpdated   string
	OrderCompleted string
	BillGenerated  string
	RoomBooked     string
	RoomCheckedOut string
	MenuUpdated    string
}

var Hospitality = HospitalityEvents{
	OrderCreated:   "Hospitality.Order.Created",
	OrderUpdated:   "Hospitality.Order.Updated",
	OrderCompleted: "Hospitality.Order.Completed",
	BillGenerated:  "Hospitality.Bill.Generated",
	RoomBooked:     "Hospitality.Room.Booked",
	RoomCheckedOut: "Hospitality.Room.CheckedOut",
	MenuUpdated:    "Hospitality.Menu.Updated",
}

// ============================
// 💰 ACCOUNTING EVENTS
// ============================
type AccountingEvents struct {
	InvoiceCreated   string
	InvoicePaid      string
	PaymentReceived  string
	JournalPosted    string
	LedgerUpdated    string
	FiscalYearClosed string
}

var Accounting = AccountingEvents{
	InvoiceCreated:   "Accounting.Invoice.Created",
	InvoicePaid:      "Accounting.Invoice.Paid",
	PaymentReceived:  "Accounting.Payment.Received",
	JournalPosted:    "Accounting.Journal.Posted",
	LedgerUpdated:    "Accounting.Ledger.Updated",
	FiscalYearClosed: "Accounting.FiscalYear.Closed",
}

// ============================
// 👩‍💼 HR EVENTS
// ============================
type HREvents struct {
	EmployeeAdded    string
	EmployeeUpdated  string
	EmployeeRemoved  string
	PayrollProcessed string
	PolicyUpdated    string
}

var HR = HREvents{
	EmployeeAdded:    "HR.Employee.Added",
	EmployeeUpdated:  "HR.Employee.Updated",
	EmployeeRemoved:  "HR.Employee.Removed",
	PayrollProcessed: "HR.Payroll.Processed",
	PolicyUpdated:    "HR.Policy.Updated",
}

// ============================
// 📢 MARKETING EVENTS
// ============================
type MarketingEvents struct {
	LeadCreated   string
	LeadConverted string
	CampaignStart string
	CampaignEnd   string
}

var Marketing = MarketingEvents{
	LeadCreated:   "Marketing.Lead.Created",
	LeadConverted: "Marketing.Lead.Converted",
	CampaignStart: "Marketing.Campaign.Started",
	CampaignEnd:   "Marketing.Campaign.Ended",
}

// ============================
// 🏥 MEDICAL EVENTS
// ============================
type MedicalEvents struct {
	PatientAdmitted   string
	PatientDischarged string
	BillingGenerated  string
	PrescriptionAdded string
}

var Medical = MedicalEvents{
	PatientAdmitted:   "Medical.Patient.Admitted",
	PatientDischarged: "Medical.Patient.Discharged",
	BillingGenerated:  "Medical.Billing.Generated",
	PrescriptionAdded: "Medical.Prescription.Added",
}

// ============================
// 🎓 UNIVERSITY EVENTS
// ============================
type UniversityEvents struct {
	StudentEnrolled string
	CourseCompleted string
	ExamScheduled   string
	ResultPublished string
}

var University = UniversityEvents{
	StudentEnrolled: "University.Student.Enrolled",
	CourseCompleted: "University.Course.Completed",
	ExamScheduled:   "University.Exam.Scheduled",
	ResultPublished: "University.Result.Published",
}
