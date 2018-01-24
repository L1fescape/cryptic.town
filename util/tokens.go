package util

import (
  "crypto/rand"
  "fmt"
)

func randomString(n int) string {
  b := make([]byte, 4)
  rand.Read(b)
  return fmt.Sprintf("%x", b)
}

func GenToken() string {
  return randomString(24)
}
