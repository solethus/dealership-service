package model

import (
	"time"

	"github.com/google/uuid"
)

type Dealership struct {
	ID              uuid.UUID        `json:"id"`
	Name            string           `json:"name"`
	Address         Address          `json:"address"`
	Phone           string           `json:"phone"`
	Email           string           `json:"email"`
	Website         string           `json:"website"`
	OperatingHours  []OperatingHours `json:"operating_hours"`
	Services        []string         `json:"services"`
	Brands          []string         `json:"brands"`
	InventoryCount  int              `json:"inventory_count"`
	SalesPersonnel  []Employee       `json:"sales_personnel"`
	Rating          float64          `json:"rating"`
	ReviewCount     int              `json:"review_count"`
	EstablishedDate time.Time        `json:"established_date"`
	LastUpdated     time.Time        `json:"last_updated"`
	IsActive        bool             `json:"is_active"`
}

type Address struct {
	Street      string      `json:"street"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	PostalCode  string      `json:"postal_code"`
	Country     string      `json:"country"`
	Coordinates GeoLocation `json:"coordinates"`
}

type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type OperatingHours struct {
	DayOfWeek string `json:"day_of_week"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
	IsClosed  bool   `json:"is_closed"`
}

type Employee struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
}

type DealershipFilter struct {
	Name     string
	City     string
	State    string
	Brand    string
	Service  string
	IsActive *bool
}

type DealershipStats struct {
	TotalDealerships   int
	ActiveDealerships  int
	AverageRating      float64
	TotalInventory     int
	TopRatedDealership string
	MostCommonBrand    string
	AverageEmployees   float64
}

type Reservation struct {
}
