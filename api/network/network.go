package network

import (
	"bytes"
	json_base "encoding/json"
	"fmt"
	"io/ioutil"

	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/emurray647/run_log/core/db"
	"github.com/emurray647/run_log/core/format"
	"github.com/emurray647/run_log/core/format/fit"
	"github.com/emurray647/run_log/core/format/fit/generated"
	"github.com/emurray647/run_log/core/format/json"
	"github.com/gorilla/mux"
)

const (
	urlPrefix = "/api/v1"

	defaultPageQuery  = uint(1)
	defaultNumPerPage = uint(20)
)

type info struct {
	db            *db.DbConnection
	fr            *fit.FitReader
	outFilterOpts []format.FilterOption
}

func (info info) allActivities(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: /activities")

	// var err error

	pageQuery := defaultPageQuery
	numPerPageQuery := defaultNumPerPage

	pageQueryString := r.URL.Query().Get("page")
	if pageQueryString != "" {
		tmp, err := strconv.ParseInt(pageQueryString, 0, 32)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			pageQuery = uint(tmp)
		}
	}

	numPerPageQueryString := r.URL.Query().Get("numPerPage")
	if numPerPageQueryString != "" {
		tmp, err := strconv.ParseUint(numPerPageQueryString, 0, 32)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			numPerPageQuery = uint(tmp)
		}
	}

	fmt.Println(pageQuery, " ", numPerPageQuery)

	start := (pageQuery - 1) * numPerPageQuery
	end := start + numPerPageQuery
	activities := info.db.GetActivitySummaries(start, end)

	w.Header().Set("Content-Type", "application/json")
	fmt.Println(r.Header.Get("origin"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json_base.NewEncoder(w).Encode(activities)

}

func (info info) activity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	b, err := info.db.GetActivity(uint(key))
	if err != nil {
		panic(err.Error())
	}
	messages := info.fr.Read(bytes.NewReader(b))

	writer := json.JsonWriter{}
	err = writer.Write(w, &fit.Activity{Messages: messages}, info.outFilterOpts...)
	if err != nil {
		panic(err.Error())
	}

}

// func (info info) helper(b []byte) {

// 	messages := info.fr.Read(bytes.NewReader(b))

// 	w := &bytes.Buffer{}
// 	writer := json.JsonWriter{}
// 	err := writer.Write(w, &fit.Activity{Messages: messages}, info.outFilterOpts...)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

func (info info) uploadActivity(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	// TODO: any sort of validation?
	messages := info.fr.Read(bytes.NewReader(reqBody))
	var session *generated.SessionDataMessage
	for i := range messages {
		if messages[i].GlobalMessageNum == generated.SessionDataMessageID {
			session = messages[i].Data.(*generated.SessionDataMessage)
			break
		}
	}

	t, _ := time.Parse("2006-01-02 15:04:05", "1989-12-31 00:00:00")
	startTime := uint32(t.Unix()) + uint32(*session.StartTime)

	fmt.Println(generated.Convert(*session.TotalDistance, generated.SESSION_TOTAL_DISTANCE_DEFAULT_UNIT, generated.UNIT_METER))

	err := info.db.AddActivity(&db.Activity{
		User:         db.User{ID: 0},
		StartTime:    uint(startTime),
		TotalTime:    float32(generated.Convert(*session.TotalTimerTime, generated.SESSION_TOTAL_TIMER_TIME_DEFAULT_UNIT, generated.UNIT_SECOND)),
		Distance:     float32(generated.Convert(*session.TotalDistance, generated.SESSION_TOTAL_DISTANCE_DEFAULT_UNIT, generated.UNIT_METER)),
		AvgHeartRate: uint(*session.AvgHeartRate),
		AvgCadence:   uint(*session.AvgCadence),
		DataGlob:     reqBody,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	// messages := info.fr.Read(bytes.NewReader(reqBody))

	w.Header().Set("Content-Type", "application/json")
	writer := json.JsonWriter{}
	writer.Write(w, &fit.Activity{Messages: messages}, info.outFilterOpts...)
}

func HandleRequests(dbConnection *db.DbConnection) {
	info := info{
		db:            dbConnection,
		fr:            &fit.FitReader{},
		outFilterOpts: json.DefaultJSONFilterOptions,
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(urlPrefix+"/", info.allActivities)
	router.HandleFunc(urlPrefix+"/activities", info.allActivities)
	router.HandleFunc(urlPrefix+"/activities/upload", info.uploadActivity)
	router.HandleFunc(urlPrefix+"/activities/{id}", info.activity)

	log.Fatal(http.ListenAndServe(":8080", router))
}
