package model

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/emurray647/run_log/internal/request"
)

type ActivitySummary struct {
	ID            int64   `json:"id,omitempty"`
	UserID        int64   `json:"user_id,omitempty"`
	StartTime     int     `json:"start_time,omitempty"`
	TotalTime     float32 `json:"duration,omitempty"`
	TotalDistance float32 `json:"distance,omitempty"`
	AvgHeartRate  uint8   `json:"avg_heart_rate,omitempty"`
	AvgCadence    uint8   `json:"avg_cadence,omitempty"`
	Ascent        float32 `json:"ascent,omitempty"`
	Descent       float32 `json:"descent,omitempty"`
}

type Activity struct {
	ActivitySummary `json:"summary,omitempty"`
	Records         ActivityRecords `json:"records,omitempty"`
	Laps            ActivityLaps    `json:"laps,omitempty"`
}

type Position struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type ActivityRecord struct {
	// timestamp
	Timestamp *float32  `json:"timestamp"`
	Position  *Position `json:"position"`
	Elevation *float32  `json:"elevation"`
	HeartRate *uint16   `json:"heartrate"`
	Cadence   *uint16   `json:"cadence"`
	Distance  *float32  `json:"distance"`
	Speed     *float32  `json:"speed"`
	// ...
}

type ActivityLap struct {
	StartTime *float32 `json:"start_time"`
	EndTime   *float32 `json:"end_time"`
	Duration  *float32 `json:"duration"`
	Distance  *float32 `json:"distance"`
	HeartRate *uint16  `json:"heartrate"`
	Cadence   *uint16  `json:"cadence"`
}

type ActivitySummaries []ActivitySummary

type ActivityRecords []ActivityRecord

type ActivityLaps []ActivityLap

func (a *Activity) DBWrite(rc request.RequestContext) (int64, error) {

	db, err := rc.GetDBConnection()
	if err != nil {
		return 0, err
	}

	queryString := fmt.Sprintf("INSERT INTO %s "+
		"(user_id, start_time, total_time, total_distance, avg_heart_rate, avg_cadence, ascent, descent, records, laps) "+
		"VALUES(?,?,?,?,?,?,?,?,?,?)", "activities")
	stmt, err := db.Prepare(queryString)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(a.UserID, a.StartTime, a.TotalTime, a.TotalDistance, a.AvgHeartRate, a.AvgCadence,
		a.Ascent, a.Descent, a.Records.Marshal(), a.Laps.Marshal())
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	a.ID = id
	return id, err
}

func (a *ActivitySummaries) DBReadAllActivities(rc request.RequestContext) error {

	*a = (*a)[:0]

	db, err := rc.GetDBConnection()
	if err != nil {
		return err
	}

	userID, err := rc.GetUserID()
	if err != nil {
		return err
	}

	queryString := fmt.Sprintf(`
		SELECT activity_id, start_time, total_time, total_distance, avg_heart_rate, avg_cadence
		FROM %s
		WHERE user_id=%d
	`, "activities", userID)
	stmt, err := db.Prepare(queryString)
	if err != nil {
		return err
	}

	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		summary := ActivitySummary{}
		err := rows.Scan(&summary.ID, &summary.StartTime, &summary.TotalTime, &summary.TotalDistance, &summary.AvgHeartRate, &summary.AvgCadence)
		if err != nil {
			return err
		}

		*a = append(*a, summary)
	}

	return nil
}

func (a *Activity) DBReadActivity(rc request.RequestContext, activityID uint) error {

	// return nil
	db, err := rc.GetDBConnection()
	if err != nil {
		return err
	}

	userID, err := rc.GetUserID()
	if err != nil {
		return err
	}

	queryString := fmt.Sprintf(`
		SELECT activity_id, start_time, total_time, total_distance, avg_heart_rate, avg_cadence, records, laps 
		FROM %s 
		WHERE activity_id=%d and user_id=%d
	`, "activities", activityID, userID)

	stmt, err := db.Prepare(queryString)
	if err != nil {
		return err
	}

	row := stmt.QueryRow()
	// probably need len stored in db too
	var recordsBuffer []byte
	var lapsBuffer []byte
	err = row.Scan(&a.ID, &a.StartTime, &a.TotalTime, &a.TotalDistance, &a.AvgHeartRate, &a.AvgCadence, &recordsBuffer, &lapsBuffer)
	if err != nil {
		return fmt.Errorf("could not scan for row with id=%d: %w", activityID, err)
	}

	err = a.Records.Unmarshal(recordsBuffer)
	if err != nil {
		return fmt.Errorf("error unmarshallnig records: %w", err)
	}

	err = a.Laps.Unmarshal(lapsBuffer)
	if err != nil {
		return fmt.Errorf("error unmarshalling laps: %w", err)
	}

	fmt.Printf("num laps: %d\n", len(a.Laps))

	return nil
}

func (ar ActivityRecords) Marshal() []byte {

	// todo: better way to get size
	totalSize := uint(len(ar)) * (4 + 4 + 4 + 4 + 2 + 2 + 4 + 4)

	buffer := make([]byte, totalSize)
	offset := 0

	for _, record := range ar {

		if record.Timestamp != nil {
			binary.BigEndian.PutUint32(buffer[offset:], math.Float32bits(*record.Timestamp))
		} else {
			binary.BigEndian.PutUint32(buffer[offset:], math.MaxUint32)
		}
		offset += 4

		binary.BigEndian.PutUint32(buffer[offset:], math.Float32bits(record.Position.Latitude))
		offset += 4
		binary.BigEndian.PutUint32(buffer[offset:], math.Float32bits(record.Position.Longitude))
		offset += 4

		if record.Elevation != nil {
			binary.BigEndian.PutUint32(buffer[offset:], math.Float32bits(*record.Elevation))
		} else {
			binary.BigEndian.PutUint32(buffer[offset:], math.MaxUint32)
		}
		offset += 4

		if record.HeartRate != nil {
			binary.BigEndian.PutUint16(buffer[offset:], *record.HeartRate)
		} else {
			binary.BigEndian.PutUint16(buffer[offset:], math.MaxUint16)
		}
		offset += 2

		if record.Cadence != nil {
			binary.BigEndian.PutUint16(buffer[offset:], *record.Cadence)
		} else {
			binary.BigEndian.PutUint16(buffer[offset:], math.MaxUint16)
		}
		offset += 2

		if record.Distance != nil {
			binary.BigEndian.PutUint32(buffer[offset:], math.Float32bits(*record.Distance))
		} else {
			binary.BigEndian.PutUint32(buffer[offset:], math.MaxUint32)
		}
		offset += 4

		if record.Speed != nil {
			binary.BigEndian.PutUint32(buffer[offset:], math.Float32bits(*record.Speed))
		} else {
			binary.BigEndian.PutUint32(buffer[offset:], math.MaxUint32)
		}
		offset += 4

	}

	return buffer
}

func (ar *ActivityRecords) Unmarshal(b []byte) error {

	*ar = (*ar)[:0]

	buffer := bytes.NewBuffer(b)
	var err error

	for buffer.Len() != 0 {

		record := ActivityRecord{}

		var timestamp uint32
		err = binary.Read(buffer, binary.BigEndian, &timestamp)
		if err != nil {
			return fmt.Errorf("error reading timestamp on %dth sample: %w", len(*ar), err)
		}
		if timestamp != math.MaxUint32 {
			record.Timestamp = new(float32)
			*record.Timestamp = math.Float32frombits(timestamp)
		} else {
			record.Timestamp = nil
		}

		record.Position = new(Position)
		err = binary.Read(buffer, binary.BigEndian, &record.Position.Latitude)
		if err != nil {
			return fmt.Errorf("error reading latitude on %dth sample: %w", len(*ar), err)
		}
		err = binary.Read(buffer, binary.BigEndian, &record.Position.Longitude)
		if err != nil {
			return fmt.Errorf("error reading longitude on %dth sample: %w", len(*ar), err)
		}

		var elevation uint32
		err = binary.Read(buffer, binary.BigEndian, &elevation)
		if err != nil {
			return fmt.Errorf("error reading elevation on %dth sample: %w", len(*ar), err)
		}
		if elevation != math.MaxUint32 {
			record.Elevation = new(float32)
			*record.Elevation = math.Float32frombits(elevation)
		} else {
			record.Elevation = nil
		}

		var heartRate uint16
		err = binary.Read(buffer, binary.BigEndian, &heartRate)
		if err != nil {
			return fmt.Errorf("error reading heartrate on %dth sample: %w", len(*ar), err)
		}
		if heartRate != math.MaxUint16 {
			record.HeartRate = &heartRate
		} else {
			record.HeartRate = nil
		}

		var cadence uint16
		binary.Read(buffer, binary.BigEndian, &cadence)
		if cadence != math.MaxUint16 {
			record.Cadence = &cadence
		} else {
			record.Cadence = nil
		}

		var distance uint32
		binary.Read(buffer, binary.BigEndian, &distance)
		if distance != math.MaxUint32 {
			record.Distance = new(float32)
			*record.Distance = math.Float32frombits(distance)
		}

		var speed uint32
		err = binary.Read(buffer, binary.BigEndian, &speed)
		if speed != math.MaxUint32 {
			record.Speed = new(float32)
			*record.Speed = math.Float32frombits(speed)
		} else {
			record.Speed = nil
		}

		*ar = append(*ar, record)
	}

	return nil

}

func (al ActivityLaps) Marshal() []byte {

	// todo: better way to get size
	totalSize := uint(len(al)) * (4 + 4 + 4 + 4 + 2 + 2)

	buffer := make([]byte, totalSize)
	offset := 0

	for _, lap := range al {

		if lap.StartTime != nil {
			binary.BigEndian.PutUint32(buffer[offset:], math.Float32bits(*lap.StartTime))
		} else {
			binary.BigEndian.PutUint32(buffer[offset:], math.MaxUint32)
		}
		offset += 4

		if lap.EndTime != nil {
			binary.BigEndian.PutUint32(buffer[offset:], math.Float32bits(*lap.EndTime))
		} else {
			binary.BigEndian.PutUint32(buffer[offset:], math.MaxUint32)
		}
		offset += 4

		if lap.Duration != nil {
			binary.BigEndian.PutUint32(buffer[offset:], math.Float32bits(*lap.Duration))
		} else {
			binary.BigEndian.PutUint32(buffer[offset:], math.MaxUint32)
		}
		offset += 4

		if lap.Distance != nil {
			binary.BigEndian.PutUint32(buffer[offset:], math.Float32bits(*lap.Distance))
		} else {
			binary.BigEndian.PutUint32(buffer[offset:], math.MaxUint32)
		}
		offset += 4

		if lap.HeartRate != nil {
			binary.BigEndian.PutUint16(buffer[offset:], *lap.HeartRate)
		} else {
			binary.BigEndian.PutUint16(buffer[offset:], math.MaxUint16)
		}
		offset += 2

		if lap.Cadence != nil {
			binary.BigEndian.PutUint16(buffer[offset:], *lap.Cadence)
		} else {
			binary.BigEndian.PutUint16(buffer[offset:], math.MaxUint16)
		}
		offset += 2
	}

	return buffer
}

func (al *ActivityLaps) Unmarshal(b []byte) error {

	*al = (*al)[:0]

	buffer := bytes.NewBuffer(b)
	var err error

	for buffer.Len() != 0 {

		lap := ActivityLap{}

		var startTime uint32
		err = binary.Read(buffer, binary.BigEndian, &startTime)
		if err != nil {
			return err
		}
		if startTime != math.MaxUint32 {
			lap.StartTime = new(float32)
			*lap.StartTime = math.Float32frombits(startTime)
		} else {
			lap.StartTime = nil
		}

		var endTime uint32
		err = binary.Read(buffer, binary.BigEndian, &endTime)
		if err != nil {
			return err
		}
		if endTime != math.MaxUint32 {
			lap.EndTime = new(float32)
			*lap.EndTime = math.Float32frombits(endTime)
		} else {
			lap.EndTime = nil
		}

		var duration uint32
		err = binary.Read(buffer, binary.BigEndian, &duration)
		if err != nil {
			return err
		}
		if duration != math.MaxUint32 {
			lap.Duration = new(float32)
			*lap.Duration = math.Float32frombits(duration)
		} else {
			lap.Duration = nil
		}

		var distance uint32
		err = binary.Read(buffer, binary.BigEndian, &distance)
		if err != nil {
			return err
		}
		if distance != math.MaxUint32 {
			lap.Distance = new(float32)
			*lap.Distance = math.Float32frombits(distance)
		} else {
			lap.Distance = nil
		}

		var heartRate uint16
		err = binary.Read(buffer, binary.BigEndian, &heartRate)
		if err != nil {
			return err
		}
		if heartRate != math.MaxUint16 {
			lap.HeartRate = &heartRate
		} else {
			lap.HeartRate = nil
		}

		var cadence uint16
		err = binary.Read(buffer, binary.BigEndian, &cadence)
		if err != nil {
			return err
		}
		if cadence != math.MaxUint16 {
			lap.Cadence = &cadence
		} else {
			lap.Cadence = nil
		}

		*al = append(*al, lap)
	}

	return nil
}
