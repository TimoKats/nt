package include

import (
  . "github.com/TimoKats/nt/include/shared"

  "time"
  "regexp"
  "strconv"
)

func hasOverlap(list1 []string, list2 []string) bool {
  for _, item1 := range list1 {
    for _, item2 := range list2 {
      if item1 == item2 {
        return true
      }
    }
  }
  return false
}

func removeIndex(s []*Note, index int) []*Note {
  return append(s[:index], s[index+1:]...)
}

func contains(slice []string, item string) bool {
  for _, v := range slice {
    if v == item {
      return true
    }
  }
  return false
}

func extractInts(input string) int {
  extractedInts := ""
  re := regexp.MustCompile(`\d+`)
  matches := re.FindAllString(input, -1)
  for _, match := range matches {
    extractedInts += match
  }
  num, convErr := strconv.Atoi(extractedInts)
  if convErr != nil {
    return -1
  }
  return num
}

func fromToday(note *Note) bool {
  current := time.Now()
  startOfToday := current.Truncate(24 * time.Hour)
  return note.Created.After(startOfToday)
}

