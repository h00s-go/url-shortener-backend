package link

import (
	"errors"

	"github.com/h00s/url-shortener-backend/db"
)

// Activity represents access to one of the links
type Activity struct {
	LinkID        int `json:"linkId"`
	clientAddress string
	AccessedAt    string `json:"AccessedAt"`
}

func insertActivity(db *db.Database, linkID int, clientAddress string) error {
	_, err := db.Conn.Exec(sqlInsertActivity, linkID, clientAddress, "NOW()")
	if err != nil {
		return errors.New("Error while inserting activity")
	}
	return nil
}
