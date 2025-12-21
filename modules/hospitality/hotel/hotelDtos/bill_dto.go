package hoteldtos

import (
	"time"

	"github.com/google/uuid"
)

type BillPreviewDTO struct {
	BookingID uuid.UUID

	BaseAmount  float64
	TaxAmount   float64
	TotalAmount float64

	TaxBreakup []TaxComponentPreviewDTO

	PaidAmount float64
	Balance    float64
}

type HotelInvoiceHotelDTO struct {
	Name    string `json:"name"`
	GSTIN   string `json:"gstin"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
}

type TaxComponentPreviewDTO struct {
	ComponentType string // CGST / SGST / IGST
	Rate          float64
	Amount        float64
}

type HotelInvoicePaymentDTO struct {
	Method    string    `json:"method"`
	Amount    float64   `json:"amount"`
	Reference *string   `json:"reference"`
	Date      time.Time `json:"date"`
}

type HotelInvoiceTaxDTO struct {
	Component string  `json:"component"` // CGST, SGST, IGST
	Rate      float64 `json:"rate"`
	Amount    float64 `json:"amount"`
}

type HotelInvoiceItemDTO struct {
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	Rate        float64 `json:"rate"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"` // ROOM_RENT, FOOD, SERVICE
}

type HotelInvoiceGuestDTO struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

type HotelInvoiceBookingDTO struct {
	RoomName string    `json:"room_name"`
	CheckIn  time.Time `json:"check_in"`
	CheckOut time.Time `json:"check_out"`
	Nights   int       `json:"nights"`
}

type HotelInvoiceDTO struct {
	InvoiceNumber string    `json:"invoice_number"`
	InvoiceDate   time.Time `json:"invoice_date"`

	Hotel   HotelInvoiceHotelDTO   `json:"hotel"`
	Guest   HotelInvoiceGuestDTO   `json:"guest"`
	Booking HotelInvoiceBookingDTO `json:"booking"`

	Items []HotelInvoiceItemDTO `json:"items"`
	Taxes []HotelInvoiceTaxDTO  `json:"taxes"`

	BaseAmount  float64 `json:"base_amount"`
	TaxAmount   float64 `json:"tax_amount"`
	TotalAmount float64 `json:"total_amount"`

	Payments   []HotelInvoicePaymentDTO `json:"payments"`
	PaidAmount float64                  `json:"paid_amount"`
	BalanceDue float64                  `json:"balance_due"`

	Status string `json:"status"`
}
