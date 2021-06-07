package middlewares

import (
	"github.com/apex/log"
	"io/ioutil"
	"net/http"
	"runtime/debug"
)

func Recover(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				all, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Warn("failed to read request body while handling panic recovery")
					all = []byte{}
				}
				defer r.Body.Close()
				log.WithFields(log.Fields{
					"stack": string(debug.Stack()),
					"path": r.RequestURI,
					"method": r.Method,
					"payload": string(all),
				}).Error("panic in http handler")

				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
