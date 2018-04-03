package link

import (
	"database/sql"
	"errors"

	"github.com/h00s/url-shortener-backend/db"
)

// Activity represents access to one of the links
type Activity struct {
	LinkID        int `json:"linkId"`
	clientAddress string
	AccessedAt    string `json:"AccessedAt"`
}

// Stats represent statistics on link
type Stats struct {
	Views int `json:"views"`
}

func insertActivity(db *db.Database, linkID int, clientAddress string) error {
	_, err := db.Conn.Exec(sqlInsertActivity, linkID, clientAddress, "NOW()")
	if err != nil {
		return errors.New("Error while inserting activity: " + err.Error())
	}
	return nil
}

func getLinkActivityStats(db *db.Database, linkID int) (*Stats, error) {
	s := &Stats{}

	err := db.Conn.QueryRow(sqlGetLinkActivityStats, linkID).Scan(&s.Views)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, errors.New("Error while getting link stats (" + err.Error() + ")")
	}
	return s, nil
}
