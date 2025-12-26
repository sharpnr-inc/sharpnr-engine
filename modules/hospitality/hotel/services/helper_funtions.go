package hotelservices

import "time"

// Helper functions
func calculateNights(checkIn, checkOut time.Time) float64 {
	days := checkOut.Sub(checkIn).Hours() / 24
	if days < 1 {
		return 1
	}
	return float64(int(days))
}
