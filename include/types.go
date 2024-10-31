package include

import (
  "time"
)

type CommandType int

const (
  None CommandType = iota
  Add
  List
  Move
  Clear
)

type Arguments struct {
  Command CommandType

  // optional fields
  Tags []string
  Text string
}

type Note struct {
  Id int
  Text string
  Done bool
  Created time.Time

  // optional
  Due time.Time
  Tags []string
  Comments []string
}

type Notebook struct {
  Notes []*Note `json:"Notebook"`
}

