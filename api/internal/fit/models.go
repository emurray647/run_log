package fit

import (
	"time"

	"github.com/emurray647/run_log/internal/fit/generated"
	"github.com/emurray647/run_log/internal/model"
)

// type Activity struct {
// 	Messages []DataMessage
// }

func convertToActivityRecord(msg *generated.RecordDataMessage, summary *model.ActivitySummary) *model.ActivityRecord {
	timestamp := float32(correctTime(uint32(*msg.Timestamp)) - uint(summary.StartTime))
	heartRate := uint16(generated.Convert(*msg.HeartRate, generated.RECORD_HEART_RATE_DEFAULT_UNIT, generated.UNIT_COUNTRATE))
	cadence := uint16(generated.Convert(*msg.Cadence, generated.RECORD_CADENCE_DEFAULT_UNIT, generated.UNIT_COUNTRATE))
	distance := float32(generated.Convert(*msg.Distance, generated.RECORD_DISTANCE_DEFAULT_UNIT, generated.UNIT_METER))
	speed := float32(generated.Convert(*msg.Speed, generated.RECORD_SPEED_DEFAULT_UNIT, generated.UNIT_METERPERSECOND))

	record := &model.ActivityRecord{
		Timestamp: &timestamp,
		Position: &model.Position{
			Latitude:  float32(generated.Convert(*msg.PositionLat, generated.RECORD_POSITION_LAT_DEFAULT_UNIT, generated.UNIT_DEGREES)),
			Longitude: float32(generated.Convert(*msg.PositionLong, generated.RECORD_POSITION_LONG_DEFAULT_UNIT, generated.UNIT_DEGREES)),
		},
		HeartRate: &heartRate,
		Cadence:   &cadence,
		Distance:  &distance,
		Speed:     &speed,
	}

	return record
}

func convertToActivityLap(msg *generated.LapDataMessage, summary *model.ActivitySummary) *model.ActivityLap {

	startTime := float32(correctTime(uint32(*msg.StartTime)) - uint(summary.StartTime))
	endTime := float32(correctTime(uint32(*msg.Timestamp)) - uint(summary.StartTime))
	duration := float32(generated.Convert(*msg.TotalElapsedTime, generated.LAP_TOTAL_ELAPSED_TIME_DEFAULT_UNIT, generated.UNIT_SECOND))
	distance := float32(generated.Convert(*msg.TotalDistance, generated.LAP_TOTAL_DISTANCE_DEFAULT_UNIT, generated.UNIT_METER))
	heartRate := uint16(generated.Convert(*msg.AvgHeartRate, generated.LAP_AVG_HEART_RATE_DEFAULT_UNIT, generated.UNIT_COUNTRATE))
	cadence := uint16(generated.Convert(*msg.AvgCadence, generated.LAP_AVG_CADENCE_DEFAULT_UNIT, generated.UNIT_COUNTRATE))

	activityLap := &model.ActivityLap{
		StartTime: &startTime,
		EndTime:   &endTime,
		Duration:  &duration,
		Distance:  &distance,
		HeartRate: &heartRate,
		Cadence:   &cadence,
	}

	return activityLap
}

func correctTime(rawTime uint32) uint {
	t, _ := time.Parse("2006-01-02 15:04:05", "1989-12-31 00:00:00")
	correctedTime := uint32(t.Unix()) + uint32(rawTime)
	return uint(correctedTime)
}
