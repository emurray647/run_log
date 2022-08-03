package activities

// import (
// 	"fmt"

// 	"github.com/emurray647/run_log/internal/request"
// )

// type ActivitySummary struct {
// 	ID            int64   `json:"id,omitempty"`
// 	UserID        int64   `json:"user_id,omitempty"`
// 	StartTime     int     `json:"start_time,omitempty"`
// 	TotalTime     float32 `json:"duration,omitempty"`
// 	TotalDistance float32 `json:"distance,omitempty"`
// 	AvgHeartRate  uint8   `json:"avg_heart_rate,omitempty"`
// 	AvgCadence    uint8   `json:"avg_cadence,omitempty"`
// }

// type Activity struct {
// 	ActivitySummary `json:"summary,omitempty"`
// 	Records         []ActivityRecord `json:"records,omitempty"`
// }

// type Position struct {
// 	Latitude  float32
// 	Longitude float32
// }

// type ActivityRecord struct {
// 	// timestamp
// 	Position  *Position
// 	Elevation *float32
// 	HeartRate *uint
// 	Cadence   *uint
// 	Distance  *float32
// 	// ...
// }

// type ActivitySummaries []ActivitySummary

// func (a *Activity) DBWrite(rc request.RequestContext) (int64, error) {

// 	db, err := rc.GetDBConnection()
// 	if err != nil {
// 		return 0, err
// 	}

// 	queryString := fmt.Sprintf("INSERT INTO %s (user_id, start_time, total_time, total_distance, avg_heart_rate, avg_cadence, data) "+
// 		"VALUES(?,?,?,?,?,?,?)", "activities")
// 	stmt, err := db.Prepare(queryString)
// 	if err != nil {
// 		return 0, err
// 	}

// 	res, err := stmt.Exec(a.UserID, a.StartTime, a.TotalTime, a.TotalDistance, a.AvgHeartRate, a.AvgCadence, []byte(""))
// 	if err != nil {
// 		return 0, err
// 	}

// 	id, err := res.LastInsertId()
// 	a.ID = id
// 	return id, err
// }

// func (a *ActivitySummaries) DBReadAllActivities(rc request.RequestContext) error {

// 	*a = (*a)[:0]

// 	db, err := rc.GetDBConnection()
// 	if err != nil {
// 		return err
// 	}

// 	userID, err := rc.GetUserID()
// 	if err != nil {
// 		return err
// 	}

// 	queryString := fmt.Sprintf(`
// 		SELECT activity_id, start_time, total_time, total_distance, avg_heart_rate, avg_cadence
// 		FROM %s
// 		WHERE user_id=%d
// 	`, "activities", userID)
// 	stmt, err := db.Prepare(queryString)
// 	if err != nil {
// 		return err
// 	}

// 	rows, err := stmt.Query()
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		summary := ActivitySummary{}
// 		err := rows.Scan(&summary.ID, &summary.StartTime, &summary.TotalTime, &summary.TotalDistance, &summary.AvgHeartRate, &summary.AvgCadence)
// 		if err != nil {
// 			return err
// 		}

// 		*a = append(*a, summary)
// 	}

// 	return nil
// }

// func (a *Activity) DBReadActivity(rc request.RequestContext, activityID uint) error {

// 	return nil
// 	// db, err := rc.GetDBConnection()
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// userID, err := rc.GetUserID()
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// queryString := fmt.Sprintf("SELECT data FROM %s WHERE activity_id=%d", "activities", activityID)
// 	// stmt, err := db.Prepare(queryString)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// row := stmt.QueryRow()

// }
