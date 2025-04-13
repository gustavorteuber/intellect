package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yourusername/intellect/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yourusername/intellect/config"
)

type MySQLPropertyRepository struct {
	DB *sql.DB
}

func NewMySQLPropertyRepository() *MySQLPropertyRepository {
	config.LoadEnvVariables()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.GetDBUser(), config.GetDBPassword(), config.GetDBHost(),
		config.GetDBPort(), config.GetDBName())

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	return &MySQLPropertyRepository{DB: db}
}

// Salva a propriedade no banco de dados
func (repo *MySQLPropertyRepository) Save(property domain.Property) error {
	_, err := repo.DB.Exec(`
		INSERT INTO auctions.properties 
		(id, link_property, title, road, geolocation, docs, photos, land_area, usable_area, 
		number_of_rooms, number_of_parking_spaces, city, state, district, inclusion_date, 
		bank, property_type, sale_type, original_price, discounted_price, discount, end_date, 
		details, first_auction_price, first_auction_date, second_auction_price, second_auction_date, 
		auctioneer_name, auctioneer_link, situation, type_payments, description, creci_code, 
		responsible_broker, bot_add_at, processed_by_ai, is_selled) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		property.ID, property.LinkProperty, property.Title, property.Road, property.Geolocation,
		property.Docs, property.Photos, property.LandArea, property.UsableArea, property.NumberOfRooms,
		property.NumberOfParkingSpaces, property.City, property.State, property.District, property.InclusionDate,
		property.Bank, property.PropertyType, property.SaleType, property.OriginalPrice, property.DiscountedPrice,
		property.Discount, property.EndDate, property.Details, property.FirstAuctionPrice, property.FirstAuctionDate,
		property.SecondAuctionPrice, property.SecondAuctionDate, property.AuctioneerName, property.AuctioneerLink,
		property.Situation, property.TypePayments, property.Description, property.CreciCode, property.ResponsibleBroker,
		property.BotAddAt, property.ProcessedByAI, property.IsSelled)

	return err
}
