// Various util functions for the notebook. Mainly centers around slices.

package include

import (
  . "github.com/TimoKats/nt/include/shared"

  "time"
)

func hasOverlap(list1 []string, list2 []string) bool {
  for _, item1 := range list1 {
    for _, item2 := range list2 {
      if item1 == item2 { return true }
    }
  }
  return false
}

func removeIndex(s []*Note, index int) []*Note {
  return append(s[:index], s[index+1:]...)
}

func containsInt(slice []int, item int) bool {
  for _, v := range slice {
    if v == item { return true }
  }
  return false
}

func containsStr(slice []string, item string) bool {
  for _, v := range slice {
    if v == item { return true }
  }
  return false
}

func fromToday(note *Note) bool {
  current := time.Now()
  startOfToday := current.Truncate(24 * time.Hour)
  return note.Created.After(startOfToday)
}

