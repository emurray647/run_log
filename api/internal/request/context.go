package request

import (
	"database/sql"
	"fmt"
	"os"

	// "log"

	log "github.com/sirupsen/logrus"

	"github.com/go-sql-driver/mysql"
)

type RequestContext struct {
	db     *sql.DB
	logger *log.Logger

	userID *uint
	vars   map[string]string
}

func (rc *RequestContext) Open() {
	rc.logger = log.New()
	rc.logger.Out = os.Stdout
}

func (rc *RequestContext) Close() error {
	if rc.db != nil {
		err := rc.db.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (rc *RequestContext) GetLogger() *log.Logger {
	return rc.logger
}

func (rc *RequestContext) GetDBConnection() (*sql.DB, error) {

	// mysql -h db -uroot -proot

	if rc.db == nil {
		cfg := mysql.Config{
			User:                 "root", //os.Getenv("DBUSER"),
			Passwd:               "root", //os.Getenv("DBPASS"),
			Net:                  "tcp",
			Addr:                 "db", //"127.0.0.1:3306",
			DBName:               "run_log_db",
			AllowNativePasswords: true,
		}
		var err error
		rc.db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			return nil, fmt.Errorf("failed to open db: %w", err)
		}
	}

	err := rc.db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return rc.db, nil

}

func (rc *RequestContext) GetUserID() (uint, error) {
	if rc.userID == nil {
		return 0, fmt.Errorf("user ID was not set")
	}
	return *rc.userID, nil
}

func (rc *RequestContext) SetUserID(id uint) {
	rc.userID = &id
}

func (rc *RequestContext) SetPathVars(vars map[string]string) {
	rc.vars = vars
}

func (rc *RequestContext) GetPathVar(key string) (string, bool) {
	value, exists := rc.vars[key]
	return value, exists
}
