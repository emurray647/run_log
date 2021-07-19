package network

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/emurray647/run_log/core/db"
	"github.com/emurray647/run_log/core/format"
	"github.com/emurray647/run_log/core/format/fit"
	"github.com/emurray647/run_log/core/format/json"
	"github.com/gorilla/mux"
)

type info struct {
	db            *db.DbConnection
	fr            *fit.FitReader
	outFilterOpts []format.FilterOption
}

func (info info) allActivities(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: /activities")
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

func (info info) uploadActivity(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	// TODO: any sort of validation?
	info.db.AddActivity(&db.Activity{
		User:     db.User{ID: 0},
		DataGlob: reqBody,
	})

	messages := info.fr.Read(bytes.NewReader(reqBody))

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
	router.HandleFunc("/", info.allActivities)
	router.HandleFunc("/activities", info.allActivities)
	router.HandleFunc("/activities/upload", info.uploadActivity)
	router.HandleFunc("/activities/{id}", info.activity)
	log.Fatal(http.ListenAndServe(":8080", router))
}
