package employee_constants

type EmployeeOnboardingStatus string

const (
	EmployeeOnboardingCreated          EmployeeOnboardingStatus = "CREATED"
	EmployeeOnboardingEmailVerified    EmployeeOnboardingStatus = "EMAIL_VERIFIED"
	EmployeeOnboardingPasswordSet      EmployeeOnboardingStatus = "PASSWORD_SET"
	EmployeeOnboardingProfileSubmitted EmployeeOnboardingStatus = "PROFILE_SUBMITTED"
	EmployeeOnboardingProfileApproved  EmployeeOnboardingStatus = "PROFILE_APPROVED"
	EmployeeOnboardingPolicyAssigned   EmployeeOnboardingStatus = "POLICY_ASSIGNED"
	EmployeeOnboardingActive           EmployeeOnboardingStatus = "ACTIVE"
	EmployeeOnboardingRejected         EmployeeOnboardingStatus = "REJECTED"
)

type EmployeeStatus string

const (
	EmployeeStatusPending  EmployeeStatus = "PENDING"
	EmployeeStatusActive   EmployeeStatus = "ACTIVE"
	EmployeeStatusDisabled EmployeeStatus = "DISABLED"
	EmployeeStatusExited   EmployeeStatus = "EXITED"
)

type EmployeeEmployementType string

const (
	EmployeeEmploymentTypeFullTime EmployeeEmployementType = "FULL_TIME"
	EmployeeEmploymentTypePartTime EmployeeEmployementType = "PART_TIME"
	EmployeeEmploymentTypeContract EmployeeEmployementType = "CONTRACT"
	EmployeeEmploymentTypeIntern   EmployeeEmployementType = "INTERN"
)
