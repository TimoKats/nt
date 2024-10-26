package include

import (
  "time"
)

type ArgType int

const (
  None ArgType = iota
  Write
  Read
  Tags
)

type Arguments struct {
  Command string
  CurrentType ArgType

  // optional fields
  Tags []string
  Text string
}

type Note struct {
  Id string
  Text string
  Created time.Time

  // optional
  Tags []string
  Expires time.Time
}

// add quick notes in seperate index that I can clean
type Notebook struct {
  TagIndex map[string][]*Note `json:"Notebook"`
}

