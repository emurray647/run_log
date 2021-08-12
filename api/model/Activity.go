package model

import (
	"github.com/emurray647/run_log/core/format/fit"
	"github.com/emurray647/run_log/core/format/fit/generated"
	"github.com/emurray647/run_log/util"
)

type Activity struct {
	Summary ActivitySummary  `json:"summary,omitempty"`
	Records []ActivityRecord `json:"records,omitempty"`
	Laps    []ActivityLap    `json:"laps,omitempty"`
}

type ActivityRecord struct {
	Timestamp   *float64 `json:"timestamp,omitempty"`
	PositionLat *float32 `json:"position_latitude,omitempty"`
	PositionLon *float32 `json:"position_longitude,omitempty"`
	Distance    *float32 `json:"distance,omitempty"`
	Speed       *float32 `json:"speed,omitempty"`
	Elevation   *float32 `json:"elevation,omitempty"`
	HeartRate   *uint    `json:"heartrate,omitempty"`
	Cadence     *uint    `json:"cadence,omitempty"`
}

type ActivityLap struct {
	// Timestamp *float64 `json:"timestamp,omitempty"`
	StartTime *float64 `json:"start_time,omitempty"`
	EndTime   *float64 `json:"end_time,omitempty"`
	Duration  *float32 `json:"duration,omitempty"`
	Distance  *float32 `json:"distance,omitempty"`
	HeartRate *uint    `json:"heartrate,omitempty"`
	Cadence   *uint    `json:"cadence,omitempty"`
}

func newActivityLap() ActivityLap {
	lap := ActivityLap{}
	// lap.Timestamp = new(float64)
	lap.StartTime = new(float64)
	lap.EndTime = new(float64)
	lap.Duration = new(float32)
	lap.Distance = new(float32)
	lap.HeartRate = new(uint)
	lap.Cadence = new(uint)

	return lap
}

func (a *Activity) Marshal(messages []fit.DataMessage) error {

	a.Summary.Marshal(messages)

	points := make([]util.LatLongPoint, 0)

	for i := range messages {
		if messages[i].GlobalMessageNum == generated.RecordDataMessageID {
			record := messages[i].Data.(*generated.RecordDataMessage)
			activityRecord := ActivityRecord{}
			activityRecord.Timestamp = new(float64)
			activityRecord.Distance = new(float32)
			// activityRecord.Elevation = new(float32)
			activityRecord.HeartRate = new(uint)
			activityRecord.Cadence = new(uint)
			activityRecord.PositionLat = new(float32)
			activityRecord.PositionLon = new(float32)
			activityRecord.Speed = new(float32)
			// *activityRecord.Timestamp = float32(generated.Convert(*record.Timestamp, generated.RECORD_TIMESTAMP_DEFAULT_UNIT, generated.UNIT_SECOND))
			// *activityRecord.Timestamp = float64(*record.Timestamp)
			*activityRecord.Timestamp = float64(correctTime(uint32(*record.Timestamp)) - a.Summary.StartTime)
			*activityRecord.Distance = float32(generated.Convert(*record.Distance, generated.RECORD_DISTANCE_DEFAULT_UNIT, generated.UNIT_METER))
			// TODO: need some error check for if fields aren't set
			// *activityRecord.Elevation = float32(generated.Convert(*record.Altitude, generated.RECORD_ALTITUDE_DEFAULT_UNIT, generated.UNIT_METER))
			*activityRecord.HeartRate = uint(generated.Convert(*record.HeartRate, generated.RECORD_HEART_RATE_DEFAULT_UNIT, generated.UNIT_COUNTRATE))
			*activityRecord.Cadence = uint(generated.Convert(*record.Cadence, generated.RECORD_CADENCE_DEFAULT_UNIT, generated.UNIT_COUNTRATE))
			*activityRecord.PositionLat = float32(generated.Convert(*record.PositionLat, generated.RECORD_POSITION_LAT_DEFAULT_UNIT, generated.UNIT_DEGREES))
			*activityRecord.PositionLon = float32(generated.Convert(*record.PositionLong, generated.RECORD_POSITION_LONG_DEFAULT_UNIT, generated.UNIT_DEGREES))
			*activityRecord.Speed = float32(generated.Convert(*record.Speed, generated.RECORD_SPEED_DEFAULT_UNIT, generated.UNIT_METERPERSECOND))

			a.Records = append(a.Records, activityRecord)

			points = append(points, util.LatLongPoint{Lat: *activityRecord.PositionLat, Long: *activityRecord.PositionLon})
		}

		if messages[i].GlobalMessageNum == generated.LapDataMessageID {
			lap := messages[i].Data.(*generated.LapDataMessage)
			activityLap := newActivityLap()
			// *activityLap.Timestamp = float32(generated.Convert(*lap.Timestamp, generated.LAP_TIMESTAMP_DEFAULT_UNIT, generated.UNIT_SECOND))
			// *activityLap.Timestamp = float64(*lap.Timestamp)
			// *activityLap.StartTime = float64(*lap.StartTime)
			// *activityLap.EndTime = float64(*lap.Timestamp)
			*activityLap.StartTime = float64(correctTime(uint32(*lap.StartTime)) - a.Summary.StartTime)
			*activityLap.EndTime = float64(correctTime(uint32(*lap.Timestamp)) - a.Summary.StartTime)
			*activityLap.Duration = float32(generated.Convert(*lap.TotalElapsedTime, generated.LAP_TOTAL_ELAPSED_TIME_DEFAULT_UNIT, generated.UNIT_SECOND))
			*activityLap.Distance = float32(generated.Convert(*lap.TotalDistance, generated.LAP_TOTAL_DISTANCE_DEFAULT_UNIT, generated.UNIT_METER))
			*activityLap.HeartRate = uint(generated.Convert(*lap.AvgHeartRate, generated.LAP_AVG_HEART_RATE_DEFAULT_UNIT, generated.UNIT_COUNTRATE))
			*activityLap.Cadence = uint(generated.Convert(*lap.AvgCadence, generated.LAP_AVG_CADENCE_DEFAULT_UNIT, generated.UNIT_COUNTRATE))

			a.Laps = append(a.Laps, activityLap)

		}
	}

	// fmt.Println(points)
	elevations, _ := util.GetElevations(points)
	for i := range elevations {
		a.Records[i].Elevation = &elevations[i]
	}

	return nil
}
