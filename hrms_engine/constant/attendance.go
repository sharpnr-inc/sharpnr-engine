package hrmsEngineConstant

// Weekday bits for Shift.WorkingDaysMask. A shift that runs Monday to Friday
// carries Mon|Tue|Wed|Thu|Fri. Stored as a smallint rather than a child table
// because the set is fixed at seven values and is always read whole.
const (
	WeekDaySunday    int16 = 1 << 0
	WeekDayMonday    int16 = 1 << 1
	WeekDayTuesday   int16 = 1 << 2
	WeekDayWednesday int16 = 1 << 3
	WeekDayThursday  int16 = 1 << 4
	WeekDayFriday    int16 = 1 << 5
	WeekDaySaturday  int16 = 1 << 6

	// WeekDaysMonToFri and WeekDaysMonToSat are the two common defaults.
	WeekDaysMonToFri = WeekDayMonday | WeekDayTuesday | WeekDayWednesday | WeekDayThursday | WeekDayFriday
	WeekDaysMonToSat = WeekDaysMonToFri | WeekDaySaturday
)

// HasWeekDay reports whether mask includes the given weekday bit.
func HasWeekDay(mask, day int16) bool {
	return mask&day != 0
}
