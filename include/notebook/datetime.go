// Module for finding deadlines related to the notes. Some jargon: timeframes refer to
// periods. For example, 2d, 3w, 2m (from now). Timestamps are actual datetimes mentioned
// in the note. Multiple formats are supported. Lacking attributes are inferred based on
// the current datetime.

package include

import (
  . "github.com/TimoKats/nt/include/shared"

  "strconv"
  "errors"
  "time"
)

var current time.Time = time.Now() // NOTE: maybe don't use init then...

func inferYear(date time.Time) int {
  var year int = current.Year()
  if date.Month() < current.Month() { year += 1 }
  return year
}

func inferMonth(date time.Time) time.Month {
  var month time.Month = current.Month()
  if date.Day() < current.Day() { month += 1 }
  return month
}

func inferDay(datestring string) time.Time {
  var date time.Time = current // don't want to change current var
  var increments int = 0 // used to prevent eternal loop.
  for date.Weekday().String()[0:3] != datestring || increments > 7 {
    date = date.AddDate(0, 0, 1)
    increments += 1
  }
  return date
}

func inferDate(date time.Time, datestring string, dateformatType int) time.Time {
  var year int
  var month time.Month
  switch dateformatType {
    case 2: // only year not known
      year = inferYear(date)
      month = date.Month()
    case 3: // year and month not known
      year = inferYear(date)
      month = inferMonth(date)
    case 4: // only weekday is known
      return inferDay(datestring)
    default:
      return date
  }
  return time.Date(year, month, date.Day(), 0, 0, 0, 0, time.UTC)
}

func mapTimeframe(timeframe byte) int {
  switch timeframe {
    case 'd':
      return 1
    case 'w':
      return 7
    case 'm':
      return 30 // NOTE:: use current month to see how many days
    default:
      return 0
  }
}

func ParseDate(datestring string) (time.Time, error) {
  var date time.Time
  var dateErr error
  var dateformatType int = -1
  for index, dateFormat := range NtConfig.Notebook.DateFormats {
    date, dateErr = time.Parse(dateFormat, datestring)
    if dateErr == nil {
      dateformatType = index
      break
    }
  }
  return inferDate(date, datestring, dateformatType), dateErr
}

func ParseTimeframe(timestring string) (time.Time, error) {
  var date time.Time = current
  var convErr error
  var factor int = 0
  if len(timestring) != 2 {
    return date, errors.New("timestring not correct format (should e.g. be '2d').")
  } else if factor, convErr = strconv.Atoi(string(timestring[0])); convErr != nil {
    return date, convErr
  }
  return date.AddDate(0, 0, factor * mapTimeframe(timestring[1])), nil
}

