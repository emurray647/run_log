package accounts

import (
	"github.com/emurray647/run_log/internal/request"
)

type Account struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (a *Account) DBWrite(rc request.RequestContext) (int64, error) {
	// var err error
	db, err := rc.GetDBConnection()

	queryString := "INSERT INTO users (id, name) VALUES (?, ?)"
	result, err := db.Exec(queryString, a.ID, a.Name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	return id, err
}

func (a *Account) DBRead(rc request.RequestContext, username string) error {

	db, err := rc.GetDBConnection()
	if err != nil {
		return err

	}

	queryString := "SELECT id, name FROM users WHERE name = ?"
	row := db.QueryRow(queryString, username)

	return row.Scan(&a.ID, &a.Name)

}
