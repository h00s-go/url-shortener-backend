package db

import (
	"testing"

	"github.com/h00s/url-shortener-backend/config"
)

func TestDB(t *testing.T) {
	c, err := config.LoadConfiguration("../configuration_test.json")
	if err != nil {
		t.Error("Unable to load configuration")
	}

	db, err := NewDatabase(c)
	if err != nil {
		t.Error("Unable to connect to DB", err)
	}

	db.conn.Query("DROP TABLE schema; DROP TABLE links;")

	err = db.Init()
	if err != nil {
		t.Error("Unable to initialize DB", err)
	}
}
