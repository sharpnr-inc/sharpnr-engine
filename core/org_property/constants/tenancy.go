package property_constants

/* -------------------------------------------------------------------------- */
/*                          Property Listing & Search                         */
/* -------------------------------------------------------------------------- */

// PropertyVisibility controls whether an outside organization can discover this
// property in search and raise a join request against it. PRIVATE properties are
// invisible to everyone except the owning org — a single-occupant property (one
// org running its own hotel) stays PRIVATE forever.
type PropertyVisibility string

const (
	PropertyVisibilityPublic  PropertyVisibility = "PUBLIC"
	PropertyVisibilityPrivate PropertyVisibility = "PRIVATE"
)

type PropertyType string

const (
	PropertyTypeStandalone PropertyType = "STANDALONE" // one building, usually one occupant
	PropertyTypeMall       PropertyType = "MALL"
	PropertyTypeComplex    PropertyType = "COMPLEX"
	PropertyTypeTower      PropertyType = "TOWER"
	PropertyTypeCampus     PropertyType = "CAMPUS"
)

/* ---------------------------------- <***> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                                    Units                                   */
/* -------------------------------------------------------------------------- */

type PropertyUnitType string

const (
	UnitTypeShop   PropertyUnitType = "SHOP"
	UnitTypeKiosk  PropertyUnitType = "KIOSK"
	UnitTypeOffice PropertyUnitType = "OFFICE"
	UnitTypeFloor  PropertyUnitType = "FLOOR"
	UnitTypeAnchor PropertyUnitType = "ANCHOR" // large anchor tenant space
	UnitTypePad    PropertyUnitType = "PAD"
)

type PropertyUnitStatus string

const (
	UnitStatusAvailable   PropertyUnitStatus = "AVAILABLE"
	UnitStatusReserved    PropertyUnitStatus = "RESERVED"
	UnitStatusOccupied    PropertyUnitStatus = "OCCUPIED"
	UnitStatusUnderFitout PropertyUnitStatus = "UNDER_FITOUT"
	UnitStatusBlocked     PropertyUnitStatus = "BLOCKED"
)

/* ---------------------------------- <***> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                                  Occupancy                                 */
/* -------------------------------------------------------------------------- */

// OccupancyTenure is how the occupant holds the space. OWNED means the unit was
// sold outright — the occupancy never expires and no rent schedule applies, but
// CAM and shared-service charges still do.
type OccupancyTenure string

const (
	TenureLease         OccupancyTenure = "LEASE"
	TenureOwned         OccupancyTenure = "OWNED"
	TenureLicense       OccupancyTenure = "LICENSE" // short-term, e.g. seasonal kiosk
	TenureOwnerOperated OccupancyTenure = "OWNER_OPERATED"
)

type OccupancyStatus string

const (
	OccupancyScheduled  OccupancyStatus = "SCHEDULED" // signed, not yet started
	OccupancyActive     OccupancyStatus = "ACTIVE"
	OccupancyExpired    OccupancyStatus = "EXPIRED"
	OccupancyTerminated OccupancyStatus = "TERMINATED"
)

/* ---------------------------------- <***> --------------------------------- */

type JoinRequestStatus string

const (
	JoinRequestPending   JoinRequestStatus = "PENDING"
	JoinRequestApproved  JoinRequestStatus = "APPROVED"
	JoinRequestRejected  JoinRequestStatus = "REJECTED"
	JoinRequestWithdrawn JoinRequestStatus = "WITHDRAWN"
)

/* ---------------------------------- <***> --------------------------------- */

/* -------------------------------------------------------------------------- */
/*                          Shared Services & Billing                         */
/* -------------------------------------------------------------------------- */

type SharedServiceType string

const (
	SharedServiceSecurityDesk  SharedServiceType = "SECURITY_DESK"
	SharedServiceReception     SharedServiceType = "RECEPTION"
	SharedServiceParking       SharedServiceType = "PARKING"
	SharedServiceHousekeeping  SharedServiceType = "HOUSEKEEPING"
	SharedServiceMaintenance   SharedServiceType = "MAINTENANCE"
	SharedServiceFireSafety    SharedServiceType = "FIRE_SAFETY"
	SharedServiceWasteDisposal SharedServiceType = "WASTE_DISPOSAL"
	SharedServicePowerBackup   SharedServiceType = "POWER_BACKUP"
)

type ChargeType string

const (
	ChargeTypeRent          ChargeType = "RENT"
	ChargeTypeCAM           ChargeType = "CAM" // common area maintenance
	ChargeTypeUtility       ChargeType = "UTILITY"
	ChargeTypeSharedService ChargeType = "SHARED_SERVICE"
	ChargeTypeDeposit       ChargeType = "DEPOSIT"
	ChargeTypeOneTime       ChargeType = "ONE_TIME"
)

// ChargeBasis decides how PropertyChargeSchedule.Amount is read: a flat figure,
// a rate multiplied by the unit's chargeable area, or a rate multiplied by a
// metered reading supplied at run time.
type ChargeBasis string

const (
	ChargeBasisFixed   ChargeBasis = "FIXED"
	ChargeBasisPerSqFt ChargeBasis = "PER_SQ_FT"
	ChargeBasisMetered ChargeBasis = "METERED"
)

type ChargeFrequency string

const (
	ChargeFrequencyMonthly   ChargeFrequency = "MONTHLY"
	ChargeFrequencyQuarterly ChargeFrequency = "QUARTERLY"
	ChargeFrequencyAnnually  ChargeFrequency = "ANNUALLY"
	ChargeFrequencyOneTime   ChargeFrequency = "ONE_TIME"
)

type ChargeLineStatus string

const (
	ChargeLineDraft     ChargeLineStatus = "DRAFT"
	ChargeLineInvoiced  ChargeLineStatus = "INVOICED"
	ChargeLineCancelled ChargeLineStatus = "CANCELLED"
)
