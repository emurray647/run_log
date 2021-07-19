package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	activityTableName = "activity_tbl"
)

type DbConnection struct {
	db *sql.DB
}

func (dbConnection *DbConnection) Open() error {
	var err error
	dbConnection.db, err = sql.Open("sqlite3", "test.db")
	return err
}

func (dbConnection *DbConnection) Close() error {
	return dbConnection.db.Close()
}

func (dbConnection *DbConnection) DropTables() error {
	stmt, err := dbConnection.db.Prepare(fmt.Sprintf("DROP TABLE IF EXISTS %s;", activityTableName))
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}

func (dbConnection *DbConnection) CreateTables() error {
	dbConnection.DropTables()
	stmt, err := dbConnection.db.Prepare(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ("+
		"activity_id INTEGER PRIMARY KEY,"+
		"user_id INTEGER,"+
		"data BLOB"+
		");",
		activityTableName))
	if err != nil {
		return err
	}
	_, err = stmt.Exec()

	return err
}

func (dbConnection *DbConnection) AddActivity(activity *Activity) error {
	stmt, err := dbConnection.db.Prepare(
		fmt.Sprintf("INSERT INTO %s (user_id, data) VALUES(?, ?)", activityTableName))
	if err != nil {
		return err
	}
	res, err := stmt.Exec(activity.User.ID, activity.DataGlob)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	activity.ID = uint(id)
	return err
}

func (dbConnection *DbConnection) GetActivity(id uint) ([]byte, error) {
	row := dbConnection.db.QueryRow(fmt.Sprintf("SELECT data FROM %s WHERE activity_id=%d", activityTableName, id))
	var result []byte
	err := row.Scan(&result)
	return result, err
}
