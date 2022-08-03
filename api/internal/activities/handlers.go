package activities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	// "github.com/emurray647/run_log/core/format/fit"
	"github.com/emurray647/run_log/internal/model"

	"github.com/emurray647/run_log/internal/fit"
	"github.com/emurray647/run_log/internal/request"
)

func UploadActivity() request.HandlerFunc {

	fitReader := fit.FitReader{}

	return func(rc request.RequestContext, w http.ResponseWriter, r *http.Request) {

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		fmt.Println("Call to reader")
		activities := fitReader.Read(bytes.NewReader(reqBody))
		activity := activities[0]
		// messages := fitReader.Read(bytes.NewReader(reqBody))

		// var session *generated.SessionDataMessage
		// for i := range messages {
		// 	if messages[i].GlobalMessageNum == generated.SessionDataMessageID {
		// 		session = messages[i].Data.(*generated.SessionDataMessage)
		// 		break
		// 	}
		// }

		// t, _ := time.Parse("2006-01-02 15:04:05", "1989-12-31 00:00:00")
		// startTime := uint32(t.Unix()) + uint32(*session.StartTime)

		// fmt.Println(generated.Convert(*session.TotalDistance, generated.SESSION_TOTAL_DISTANCE_DEFAULT_UNIT, generated.UNIT_METER))
		userId, err := rc.GetUserID()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		activity.UserID = int64(userId)

		// activity := &model.Activity{
		// 	ActivitySummary: model.ActivitySummary{
		// 		UserID:        int64(userId),
		// 		StartTime:     int(startTime),
		// 		TotalTime:     float32(generated.Convert(*session.TotalTimerTime, generated.SESSION_TOTAL_TIMER_TIME_DEFAULT_UNIT, generated.UNIT_SECOND)),
		// 		TotalDistance: float32(generated.Convert(*session.TotalDistance, generated.SESSION_TOTAL_DISTANCE_DEFAULT_UNIT, generated.UNIT_METER)),
		// 		AvgHeartRate:  uint8(*session.AvgHeartRate),
		// 		AvgCadence:    uint8(*session.AvgCadence),
		// 	},
		// }

		_, err = activity.DBWrite(rc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to write to storage"))
			fmt.Println(err)
			return
		}

		// err = info.db.AddActivity(&db.Activity{
		// 	User:         db.User{ID: 0},
		// 	StartTime:    uint(startTime),
		// 	TotalTime:    float32(generated.Convert(*session.TotalTimerTime, generated.SESSION_TOTAL_TIMER_TIME_DEFAULT_UNIT, generated.UNIT_SECOND)),
		// 	Distance:     float32(generated.Convert(*session.TotalDistance, generated.SESSION_TOTAL_DISTANCE_DEFAULT_UNIT, generated.UNIT_METER)),
		// 	AvgHeartRate: uint(*session.AvgHeartRate),
		// 	AvgCadence:   uint(*session.AvgCadence),
		// 	DataGlob:     reqBody,
		// })
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		// writer := json.JsonWriter{}
		// writer.Write(w, &fit.Activity{Messages: messages}, info.outFilterOpts...)

	}
}

func GetActivities() request.HandlerFunc {
	return func(rc request.RequestContext, w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(http.StatusNotImplemented)

		summaries := model.ActivitySummaries{}

		summaries.DBReadAllActivities(rc)

		fmt.Println(summaries)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		b, _ := json.Marshal(summaries)
		w.Write(b)

	}
}

func GetActivity() request.HandlerFunc {
	return func(rc request.RequestContext, w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(http.StatusNotImplemented)

		activity := model.Activity{}

		activity_id, ok := rc.GetPathVar("activity_id")
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("did not provide activity ID"))
		}

		aid, err := strconv.ParseUint(activity_id, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("could not convert %s to integer", activity_id)))
		}

		err = activity.DBReadActivity(rc, uint(aid))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			rc.GetLogger().Errorf("error reading: %w", err)
			fmt.Println(err.Error())
		}

		b, err := json.Marshal(activity)
		if err != nil {
			rc.GetLogger().Errorf("failed to output json: %w", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)

	}
}
