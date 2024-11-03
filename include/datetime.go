package include

import (
  "time"
)

func fromToday(note *Note) bool {
  current := time.Now()
  startOfToday := current.Truncate(24 * time.Hour)
  return note.Created.After(startOfToday)
}

