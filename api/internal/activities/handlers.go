package activities

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/emurray647/run_log/internal/elevation"
	"github.com/emurray647/run_log/internal/model"
	"github.com/emurray647/run_log/internal/response"

	"github.com/emurray647/run_log/internal/fit"
	"github.com/emurray647/run_log/internal/request"
)

func UploadActivity() request.HandlerFunc {

	fitReader := fit.FitReader{}

	return func(rc request.RequestContext, w http.ResponseWriter, r *http.Request) {

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			response.WriteJSONError(w, http.StatusInternalServerError, "failed to read request body")
			return
		}

		activities := fitReader.Read(bytes.NewReader(reqBody))
		activity := activities[0]

		err = elevation.FixElevations(&activity)
		if err != nil {
			response.WriteJSONError(w, http.StatusInternalServerError, "failed adding elevation")
			return
		}

		userId, err := rc.GetUserID()
		if err != nil {
			response.WriteJSONError(w, http.StatusBadRequest, "could not find user ID")
			return
		}
		activity.UserID = int64(userId)

		_, err = activity.DBWrite(rc)
		if err != nil {
			response.WriteJSONError(w, http.StatusInternalServerError, "failed to write activity to storage")
			rc.GetLogger().Errorf("failed to write activity to storage: %w", err)
			return
		}

		response.WriteJSON(w, http.StatusOK, nil)

	}
}

func GetActivities() request.HandlerFunc {
	return func(rc request.RequestContext, w http.ResponseWriter, r *http.Request) {

		summaries := model.ActivitySummaries{}

		summaries.DBReadAllActivities(rc)

		response.WriteJSON(w, http.StatusOK, summaries)

	}
}

func GetActivity() request.HandlerFunc {
	return func(rc request.RequestContext, w http.ResponseWriter, r *http.Request) {
		activity := model.Activity{}

		activity_id, ok := rc.GetPathVar("activity_id")
		if !ok {
			response.WriteJSONError(w, http.StatusBadRequest, "did not provide activity ID")
			return
		}

		aid, err := strconv.ParseUint(activity_id, 10, 64)
		if err != nil {
			response.WriteJSONError(w, http.StatusBadRequest, fmt.Sprintf("could not convert %s to integer", activity_id))
			return
		}

		err = activity.DBReadActivity(rc, uint(aid))
		if err != nil {
			response.WriteJSONError(w, http.StatusBadRequest, fmt.Sprintf("failed to read activity %d", aid))
			return
		}

		response.WriteJSON(w, http.StatusOK, activity)

	}
}
