package hotelservices

import (
	"context"
	"fmt"
	"sharpnr-engine/modules/hospitality/hospitalityconstants"
	hoteldtos "sharpnr-engine/modules/hospitality/hotel/hotelDtos"
	hotel_models "sharpnr-engine/modules/hospitality/hotel/models"
	"strings"
	"time"
)

func (s *hotelService) buildInvoiceDTO(
	ctx context.Context,
	booking *hotel_models.HotelRoomBooking,
	bill *hotel_models.HotelBill,
) (*hoteldtos.HotelInvoiceDTO, error) {

	hotel, _ := s.hotelRepo.GetHotelByID(ctx, booking.HotelID)
	guest, _ := s.hotelRepo.GetHotelGuestById(ctx, booking.PrimaryGuestID)
	charges, _ := s.hotelRepo.GetChargesByBookingID(ctx, booking.ID)
	taxes, _ := s.hotelRepo.GetBillTaxes(ctx, bill.ID)
	payments, _ := s.hotelRepo.GetPaymentsByBookingId(ctx, booking.ID)

	items := mapChargesToInvoiceItems(charges)
	taxDTOs := mapBillTaxesToDTO(taxes)
	paymentDTOs, paidAmount := mapPaymentsToDTO(payments)

	return &hoteldtos.HotelInvoiceDTO{
		InvoiceNumber: generateInvoiceNo(bill),
		InvoiceDate:   *bill.FinalizedAt,

		Hotel: hoteldtos.HotelInvoiceHotelDTO{
			Name:    hotel.DisplayName,
			GSTIN:   hotel.GstinNumber,
			Address: hotel.Address,
			City:    hotel.City,
			State:   hotel.State,
		},
		Guest: hoteldtos.HotelInvoiceGuestDTO{
			Name:   guest.Name,
			Mobile: guest.Mobile,
		},
		Booking: hoteldtos.HotelInvoiceBookingDTO{
			RoomName: "Room",
			CheckIn:  booking.CheckInDate,
			CheckOut: booking.CheckOutDate,
			Nights:   int(calculateNights(booking.CheckInDate, booking.CheckOutDate)),
		},

		Items: items,
		Taxes: taxDTOs,

		BaseAmount:  bill.BaseAmount,
		TaxAmount:   bill.TaxAmount,
		TotalAmount: bill.TotalAmount,

		Payments:   paymentDTOs,
		PaidAmount: paidAmount,
		BalanceDue: bill.TotalAmount - paidAmount,

		Status: bill.Status,
	}, nil
}

func mapChargesToInvoiceItems(
	charges []hotel_models.HotelBookingCharge,
) []hoteldtos.HotelInvoiceItemDTO {

	items := make([]hoteldtos.HotelInvoiceItemDTO, 0, len(charges))

	for _, c := range charges {

		desc := c.Description
		if desc == "" {
			desc = humanizeChargeType(c.ChargeType)
		}

		item := hoteldtos.HotelInvoiceItemDTO{
			Description: desc,
			Quantity:    1,
			Rate:        c.Amount,
			Amount:      c.Amount,
			Type:        c.ChargeType,
		}

		items = append(items, item)
	}

	return items
}

func humanizeChargeType(ct hospitalityconstants.HotelChargeType) string {
	switch ct {
	case hospitalityconstants.RoomRent:
		return "Room Rent"
	case hospitalityconstants.FoodOrder:
		return "Food Order"
	case hospitalityconstants.BarOrder:
		return "Bar Order"
	case hospitalityconstants.LaundryService:
		return "Laundry Service"
	case hospitalityconstants.SpaService:
		return "Spa Service"
	case hospitalityconstants.Minibar:
		return "Minibar"
	case hospitalityconstants.ServiceCharge:
		return "Service Charge"
	default:
		return string(ct)
	}
}

func calculateNightsForInvoice(booking *hotel_models.HotelRoomBooking) int {

	checkIn := booking.CheckInDate.Truncate(24 * time.Hour)
	checkOut := booking.CheckOutDate.Truncate(24 * time.Hour)

	nights := int(checkOut.Sub(checkIn).Hours() / 24)

	if nights < 1 {
		return 1
	}

	return nights
}

func generateInvoiceNo(bill *hotel_models.HotelBill) string {

	datePart := bill.CreatedAt.Format("20060102")

	// Use last 6 chars of UUID for uniqueness
	uuidPart := bill.ID.String()
	uuidPart = uuidPart[len(uuidPart)-6:]

	return fmt.Sprintf("INV-%s-%s", datePart, strings.ToUpper(uuidPart))
}

func mapPaymentsToDTO(
	payments []hotel_models.HotelRoomBookingPayment,
) ([]hoteldtos.HotelInvoicePaymentDTO, float64) {

	result := make([]hoteldtos.HotelInvoicePaymentDTO, 0, len(payments))
	totalPaid := 0.0

	for _, p := range payments {
		result = append(result, hoteldtos.HotelInvoicePaymentDTO{
			Method:    string(p.Mode),
			Amount:    p.Amount,
			Reference: &p.Reference,
			Date:      p.CreatedAt,
		})

		totalPaid += p.Amount
	}

	return result, totalPaid
}

func mapBillTaxesToDTO(
	taxes []hotel_models.HotelBillTax,
) []hoteldtos.HotelInvoiceTaxDTO {

	result := make([]hoteldtos.HotelInvoiceTaxDTO, 0, len(taxes))

	for _, t := range taxes {
		result = append(result, hoteldtos.HotelInvoiceTaxDTO{
			Component: t.ComponentType,
			Rate:      t.Rate,
			Amount:    t.Amount,
		})
	}

	return result
}
