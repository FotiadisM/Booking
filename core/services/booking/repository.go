package booking

import "context"

// Date sd
type Date struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

// CardInfo fd
type CardInfo struct {
	Type     string `json:"type"` // Mastercard, Visa etc
	Number   string `json:"number"`
	Name     string `json:"name"`
	ExpMonth string `json:"expire_month"`
	ExpYear  string `json:"expire_year"`
	CVV      string `json:"cvv"`
}

// Booking describes the properties of a booking
type Booking struct {
	ID        string   `json:"id"`
	UserID    string   `json:"user_id"`
	ListingID string   `json:"listing_id"`
	Days      []Date   `json:"days"`
	CardIndo  CardInfo `json:"card_info,omiyempty"`
}

// Repository sdf
type Repository interface {
	StoreBooking(context.Context, *Booking) error
}
