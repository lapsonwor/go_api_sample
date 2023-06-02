package sqlc

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriverMain = "postgres"
	dbSource     = "postgresql://root:polka@localhost:5432/game_backend_db?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	db, err := sql.Open(dbDriverMain, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to the db", err)
	}
	testQueries = New(db)
	os.Exit(m.Run())
}

func TestGetBuilding(t *testing.T) {
	var building_id int32 = 1
	building, err := testQueries.GetBuilding(context.Background(), building_id)
	if err != nil {
		log.Fatal("Cannot get the building", err)
	}
	log.Fatal(building)
}

func TestGetCardByTokenID(t *testing.T) {
	var card_id int32 = 1
	card, err := testQueries.GetCardByTokenID(context.Background(), card_id)
	if err != nil {
		log.Fatal("Cannot get the card", err)
	}
	log.Fatal(&card)
}

func TestCreateRarity(t *testing.T) {
	var rarities = []string{
		"Common", "Rare", "Epic", "Legendary", "Mythical",
	}
	for _, rarity := range rarities {
		_, err := testQueries.CreateRarity(context.Background(), rarity)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestGetWhitelistByWallet(t *testing.T) {
	var wallet_address string = "0x01096b8C74AeDfE6D0F4abbEc397b6A39334649F"
	building, err := testQueries.GetWhitelistByWallet(context.Background(), wallet_address)
	if err != nil {
		log.Fatal("Your wallet is not in the whitelist", err)
	}
	log.Fatal(building)
}
