package middleware

import (
	"errors"
	"net/http"

	"github.com/ArcherLee25040/Go_api/go_api/api"
	"github.com/chanxuehong/wechat/mch/tools"
	"github.com/opentracing/opentracing-go/log"
)

// import(	"errors"
// "net/http"
// "github.com/ArcherLee25040/go_api/internal/tools"
// log "github.com/sirupsen/logrus")

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("Username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetail
		loginDetails = (*database).GetUserLoginDatails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return

		}

		next.ServeHTTP(w, r)
	})
}
