package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emurray647/run_log/internal/request"
	"github.com/gorilla/mux"
)

func CreateAccount() request.HandlerFunc {

	return func(rc request.RequestContext, w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := r.Form.Get("name")
		if name == "" {
			// TODO: error message
			// w.Write()
			fmt.Println(r.Form)
			rc.GetLogger().Error("name parameter required to create account")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// TODO: verify that username does not already exist

		account := Account{
			Name: name,
		}

		id, err := account.DBWrite(rc)
		if err != nil {
			fmt.Println(err.Error())
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("insert id %d", id)))

	}
}

func GetAccount() request.HandlerFunc {
	return func(rc request.RequestContext, w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

		account := Account{}
		err := account.DBRead(rc, username)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(account)
		w.Write(b)

	}
}
