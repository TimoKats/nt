package include

import (
  "time"
)

type CommandType int
type StatusType int

const (
  None CommandType = iota
  Add
  Get
  Clear
)

const (
  Todo StatusType = iota
  InProgress
  Done
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
  Status StatusType
  Created time.Time

  // optional
  Due time.Time
  Tags []string
  Comments []string
}

type Notebook struct {
  Notes []*Note `json:"Notebook"`
}

