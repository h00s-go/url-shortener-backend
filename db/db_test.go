package db

import (
	"testing"

	"github.com/h00s/url-shortener-backend/config"
)

func TestDB(t *testing.T) {
	c, err := config.Load("../configuration_test.json")
	if err != nil {
		t.Error("Unable to load configuration")
	}

	db, err := Connect(c)
	if err != nil {
		t.Error("Unable to connect to DB", err)
	}

	db.Conn.Query("DROP TABLE schema; DROP TABLE links;")

	err = db.Migrate()
	if err != nil {
		t.Error("Unable to migrate DB", err)
	}
}
