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
  Tags
  Help
  Clear
  Comment
)

type Arguments struct {
  Command CommandType

  // optional fields
  Flags []string
  Tags []string
  Text string
  NoteId int
}

type Note struct {
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

