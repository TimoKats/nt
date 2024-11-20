package include

import (
  "time"
)

// parser

type CommandType int

const (
  None CommandType = iota

  // notebook
  Add
  List
  Move
  Tags
  Clear
  Server
  Search
  Comment

  // server
  Pull
  Push
  Ping
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
  DateFormats []string `toml:"date_format"`
}

type ServerConfig struct {
  Url string `toml:"url"`
  Port string `toml:"port"`
}

type Config struct {
  Notebook NotebookConfig `toml:"notebook"`
  Server ServerConfig `toml:"server"`
}

