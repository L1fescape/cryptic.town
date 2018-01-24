package util

import (
  "log"
  "net/http"
  "os"
)

var Log = log.New(os.Stdout, "", 0)

type responseWriter struct {
  http.ResponseWriter
  statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
  rw.statusCode = code
  rw.ResponseWriter.WriteHeader(code)
}

func WrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    rw := &responseWriter{w, http.StatusOK}
    wrappedHandler.ServeHTTP(rw, req)

    Log.Printf("%s %s %d %s", req.Host, req.Method, rw.statusCode, req.URL.Path)
  })
}

