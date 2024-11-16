package include

import (
  "time"
)

// parser

type CommandType int

const (
  None CommandType = iota
  Add
  List
  Move
  Tags
  Clear
  Server
  Comment

  Pull
  Push
)

type Arguments struct {
  Command CommandType

  Flags []string
  Tags []string
  Text string
  NoteId int
}

// notebook

type Note struct {
  Text string
  Done bool
  Created time.Time

  Due time.Time
  Tags []string
  Comments []string
}

type Notebook struct {
  Notes []*Note `json:"Notebook"`
}

// server

type Authentication struct {
  Username []byte
  Password []byte
}

// config

type NotebookConfig struct {
  Width int `toml:"width"`
}

type ServerConfig struct {
  Url string `toml:"url"`
  Port string `toml:"port"`
}

type Config struct {
  Notebook NotebookConfig `toml:"notebook"`
  Server ServerConfig `toml:"server"`
}

