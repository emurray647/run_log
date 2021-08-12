package model

import (
	"time"

	"github.com/emurray647/run_log/core/format/fit"
	"github.com/emurray647/run_log/core/format/fit/generated"
)

type ActivitySummary struct {
	ID              uint    `json:"id"`
	StartTime       uint    `json:"start_time"`
	DurationSeconds float32 `json:"duration"`
	Distance        float32 `json:"distance"`
	AvgHeartRate    uint    `json:"avg_heart_rate"`
	AvgCadence      uint    `json:"avg_cadence"`
}

// It should be okay if this knows about the FitReader structs ...
// so we should be able to have conversion functions here
// Should it be passed the whole data glob?

func (a *ActivitySummary) Marshal(messages []fit.DataMessage) error {
	for i := range messages {
		if messages[i].GlobalMessageNum == generated.SessionDataMessageID {
			session := messages[i].Data.(*generated.SessionDataMessage)
			t, _ := time.Parse("2006-01-02 15:04:05", "1989-12-31 00:00:00")
			startTime := uint32(t.Unix()) + uint32(*session.StartTime)
			a.StartTime = uint(startTime)
			a.DurationSeconds = float32(generated.Convert(*session.TotalTimerTime, generated.SESSION_TOTAL_TIMER_TIME_DEFAULT_UNIT, generated.UNIT_SECOND))
			a.Distance = float32(generated.Convert(*session.TotalDistance, generated.SESSION_TOTAL_DISTANCE_DEFAULT_UNIT, generated.UNIT_METER))
			a.AvgHeartRate = uint(*session.AvgHeartRate)
			a.AvgCadence = uint(*session.AvgCadence)
		}
	}

	return nil
}

func correctTime(rawTime uint32) uint {
	t, _ := time.Parse("2006-01-02 15:04:05", "1989-12-31 00:00:00")
	correctedTime := uint32(t.Unix()) + uint32(rawTime)
	return uint(correctedTime)
}
