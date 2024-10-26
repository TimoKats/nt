package include

import (
  "math/rand"
  "time"
)

var SeededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var IDCHARSET string = "abcdefghijklmnopqrstuvwxyz"

func randomString(length int) string {
  id := make([]byte, length)
  for i := range id {
    id[i] = IDCHARSET[SeededRand.Intn(len(IDCHARSET))]
  }
  return string(id)
}

func contains(slice []string, item string) bool {
  for _, v := range slice {
    if v == item {
      return true
    }
  }
  return false
}
