package main

import (
  notebook "github.com/TimoKats/nt/include/notebook"
  . "github.com/TimoKats/nt/include/shared"

  "testing"
  "time"
)

func TestAddNote(t *testing.T) {
  arguments := Arguments{ Command: 1, Text: "hello" }
  runErr := notebook.AddNote(arguments)
  if runErr != nil && NtNotes.Notes[0].Text == "hello" {
    t.Errorf("Failed to add note.")
  }
}

func TestAddTaggedNote(t *testing.T) {
  arguments := ParseArgs([]string{"nt add what :due:2010-01-02 :tag:hello"})
  runErr := notebook.AddNote(arguments)
  testNote := Note{
    Text: "What",
    Deadline: time.Date(2010, 1, 2, 0, 0, 0, 0, time.UTC),
    Tags: []string{ "hello" },
  }
  if runErr != nil && NtNotes.Notes[0] == &testNote {
    t.Errorf("Failed to add note.")
  }
}

func TestListNotes(t *testing.T) {
  arguments := Arguments{ Command: 2 }
  runErr := notebook.ListNotebook(arguments)
  if runErr != nil {
    t.Errorf("Failed to list notes.")
  }
}

func TestClearNotes(t *testing.T) {
  arguments := Arguments{ Command: 5 }
  runErr := notebook.ClearNotebook(arguments)
  if runErr != nil && len(NtNotes.Notes) == 0 {
    t.Errorf("Failed to remove notes.")
  }
}

func init() {
  TestMode = true
  NtConfig = DefaultConfig
}
