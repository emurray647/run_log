package request

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HandlerFunc func(RequestContext, http.ResponseWriter, *http.Request)

func RouteRequest(fun HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := RequestContext{}
		ctx.Open()
		defer ctx.Close()

		vars := mux.Vars(r)
		if user_id := vars["user_id"]; user_id != "" {
			uid, err := strconv.ParseUint(user_id, 10, 32)
			if err != nil {
				ctx.GetLogger().Errorf("unable to parse user_id {%s}: %w", user_id, err)
			}
			ctx.SetUserID(uint(uid))
		}

		ctx.SetPathVars(vars)

		defer func() {
			ctx.GetLogger().Warn("setting header")
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}()

		w.Header().Set("Access-Control-Allow-Origin", "*")

		fun(ctx, w, r)
	}
}
