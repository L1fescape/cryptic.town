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

func NewFilelog(filename string) (*log.Logger, error) {
  f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
  if err != nil {
    return nil, err
  }

  logger := log.New(f, "", log.Ldate|log.Ltime)

  return logger, nil
}

