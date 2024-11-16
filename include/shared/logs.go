package include

import (
  "log"
  "os"
)

var (
  Action *log.Logger
  Warn *log.Logger
  Info *log.Logger
  Error *log.Logger
  Fatal *log.Logger
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Yellow = "\033[33m"
var Magenta = "\033[35m"

func init() {
  Info = log.New(os.Stdout, "", 0)
  Warn = log.New(os.Stdout, Yellow + "warn:  " + Reset, 0)
  Error = log.New(os.Stdout, Red + "error: " + Reset, 0)
  Fatal = log.New(os.Stdout, Magenta + "fatal: " + Reset, 0)
}
