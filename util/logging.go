package util

import (
  "log"
  "net/http"
  "os"
)

var Log = log.New(os.Stdout, "", 0)

type loggingResponseWriter struct {
  http.ResponseWriter
  statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
  return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
  lrw.statusCode = code
  lrw.ResponseWriter.WriteHeader(code)
}

func WrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
    lrw := NewLoggingResponseWriter(w)
    wrappedHandler.ServeHTTP(lrw, req)

    statusCode := lrw.statusCode
    Log.Printf("%s %s %d %s", req.Host, req.Method, statusCode, req.URL.Path)
  })
}

