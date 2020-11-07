package listing

import "context"

// Listing defines the properties of a Listing
type Listing struct {
	ID            string  `json:"id"`
	UserID        string  `json:"user_id"`
	UserName      string  `json:"user_name"`
	PriceDay      float32 `json:"price_day"`
	PriceWeek     float32 `json:"price_week"`
	PriceMonth    float32 `json:"price_month"`
	BedNum        int     `json:"bed_num"`
	PeopleNum     int     `json:"people_num"`
	RoomNum       int     `json:"room_num"`
	BathroomNum   int     `json:"bathroom_num"`
	HasTV         bool    `json:"has_tv"`
	HasWifi       bool    `json:"has_wifi"`
	HasKitchen    bool    `json:"has_kitcen"`
	HasLivingRoom bool    `json:"has_livingroom"`
	SquareMeters  int     `json:"square_meters"`
	Description   string  `json:"description"`
	Rules         string  `json:"rules"`
	ReviewNum     int     `json:"review_num"`
	ReviewAvrg    float32 `json:"review_avrg"`
	Street        string  `json:"street"`
	Number        int     `json:"number"`
	Neighbourhood string  `json:"neighbourhood"`
	City          string  `json:"city"`
	State         string  `json:"state"`
	Zipcode       int     `json:"zipcode"`
	Country       string  `json:"country"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longtitude"`
	Created       int64   `json:"-"`
}

// Repository describes the persistence on user model
type Repository interface {
	GetListings(context.Context) ([]*Listing, error)
	GetListingByID(context.Context, string) (*Listing, error)
	CreateListing(context.Context, *Listing) error
}
