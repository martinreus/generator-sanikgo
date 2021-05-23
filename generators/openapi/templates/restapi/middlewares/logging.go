package middlewares

import (
	"github.com/apex/log"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type formatter struct {
}

type logEntry struct {
	r *http.Request
}

func (e logEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	log.WithFields(log.Fields{
		"status": status,
		"payloadSize": bytes,
		"duration": elapsed,
		"path": e.r.RequestURI,
		"method": e.r.Method,
	}).Debug("Request handled")
}

func (e logEntry) Panic(v interface{}, stack []byte) {
	log.WithFields(log.Fields{
		"panic": v,
		"stack": stack,
	}).Error("Panic in http handler")
}

func (c formatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	return logEntry{r}
}

func Logger() func(next http.Handler) http.Handler {
	return middleware.RequestLogger(formatter{})
}