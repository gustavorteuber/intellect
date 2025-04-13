package domain

import "time"

type Property struct {
	ID                   string
	LinkProperty         string
	Title                string
	Road                 string
	Geolocation          string
	Docs                 string
	Photos               string
	LandArea             string
	UsableArea           string
	NumberOfRooms       string
	NumberOfParkingSpaces string
	City                 string
	State                string
	District             string
	InclusionDate        time.Time
	Bank                 string
	PropertyType         string
	SaleType             string
	OriginalPrice        float64
	DiscountedPrice      float64
	Discount             float64
	EndDate              time.Time
	Details              string
	FirstAuctionPrice    float64
	FirstAuctionDate     time.Time
	SecondAuctionPrice   float64
	SecondAuctionDate    time.Time
	AuctioneerName       string
	AuctioneerLink       string
	Situation            string
	TypePayments         string
	Description          string
	CreciCode            string
	ResponsibleBroker    string
	BotAddAt             time.Time
	ProcessedByAI        bool
	IsSelled             bool
}
