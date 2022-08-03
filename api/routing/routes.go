package routing

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/emurray647/run_log/internal/accounts"
	"github.com/emurray647/run_log/internal/activities"
	"github.com/emurray647/run_log/internal/request"
)

var (
	apiPrefix = "/api/v1"
)

type route struct {
	path string
	fun  request.HandlerFunc
}

func BuildRoutes() (http.Handler, error) {
	routes := []route{
		{
			path: "/test",
			fun: func(ctx request.RequestContext, w http.ResponseWriter, r *http.Request) {
				fmt.Println("Handling /")
			},
		},
		{
			path: "/users/create",
			fun:  accounts.CreateAccount(),
		},
		{
			path: "/users/{username}",
			fun:  accounts.GetAccount(),
		},

		{
			path: "/users/{user_id}/activities/upload",
			fun:  activities.UploadActivity(),
		},
		{
			path: "/users/{user_id}/activities",
			fun:  activities.GetActivities(),
		},
		{
			//
			path: "/users/{user_id}/activities/{activity_id}",
			fun:  activities.GetActivity(),
		},
	}

	fmt.Println(routes)
	return buildRoutes(routes)
}

func buildRoutes(routes []route) (http.Handler, error) {
	router := mux.NewRouter().StrictSlash(false)
	for _, r := range routes {
		fmt.Printf("building handler for %s\n", r.path)
		// router.HandleFunc(r.path, r.fun)
		router.HandleFunc(apiPrefix+r.path, request.RouteRequest(r.fun))
	}

	return router, nil
}
