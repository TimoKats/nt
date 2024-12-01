package include

import (
  . "github.com/TimoKats/nt/include/shared"

  "strings"
  "errors"
  "time"
)

var Notes Notebook
var LoadErr error

func noteSelected(index int, note *Note, arguments Arguments) bool {
  if hasOverlap(arguments.Tags, note.Tags) {
    return true
  } else if containsInt(arguments.NoteIds, index) {
    return true
  } else if contains(arguments.Flags, "--done") && note.Done {
    return true
  } else if contains(arguments.Flags, "--old") && !fromToday(note) {
    return true
  } else if contains(arguments.Flags, "--today") && fromToday(note) {
    return true
  } else if contains(arguments.Flags, "--all") {
    return true
  }
  return false
}

func argumentsEmpty(arguments Arguments) bool {
  return (
    len(arguments.Text) == 0 &&
    len(arguments.Tags) == 0 &&
    len(arguments.Flags) == 0 &&
    len(arguments.NoteIds) == 0)
}

func ClearNotebook(arguments Arguments) error {
  correctedIndex := 0
  if argumentsEmpty(arguments) {
    Notes.Notes = []*Note{}
  } else {
    for index, note := range Notes.Notes {
      if noteSelected(index, note, arguments) {
        Notes.Notes = removeIndex(Notes.Notes, correctedIndex)
      } else {
        correctedIndex += 1
      }
    }
  }
  writeErr := WriteNotebook(Notes)
  return writeErr
}

func ListNotebook(arguments Arguments) error {
  if len(arguments.NoteIds) != 1 { formatSummaryHeader() }
  for index, note := range Notes.Notes {
    if noteSelected(index, note, arguments) || argumentsEmpty(arguments) {
      if len(arguments.NoteIds) != 1 {
        formatSummaryOutput(index, note)
      } else if containsInt(arguments.NoteIds, index) {
        formatSingleOutput(note)
      }
    }
  }
  return nil
}

func ReadTags(arguments Arguments) error {
  tags := make(map[string]int)
  for _, note := range Notes.Notes {
    for _, tag := range note.Tags {
      tags[tag] += 1
    }
  }
  for tagname, _ := range tags {
    Info.Println(tagname)
  }
  return nil
}

func AddNote(arguments Arguments) error {
  note := Note{
    Text: arguments.Text,
    Created: time.Now(),
    Tags: arguments.Tags,
    Deadline: arguments.Deadline,
  }
  Notes.Notes = append(Notes.Notes, &note)
  writeErr := WriteNotebook(Notes)
  return writeErr
}

func MoveNote(arguments Arguments) error {
  for index, note := range Notes.Notes {
    if containsInt(arguments.NoteIds, index) {
      Notes.Notes[index].Done = !note.Done
    }
  }
  writeErr := WriteNotebook(Notes)
  return writeErr
}

func SearchNote(arguments Arguments) error {
  arguments.Flags = []string{}
  var listErr error
  for index, note := range Notes.Notes {
    for _, query := range strings.Fields(arguments.Text) {
      if strings.Contains(note.Text, query) {
        arguments.NoteIds = append(arguments.NoteIds, index)
        break
      }
    }
  }
  listErr = ListNotebook(arguments)
  return listErr
}

func ModifyNote(arguments Arguments) error {
  if len(arguments.NoteIds) == 0 {
    return errors.New("No NoteId provided")
  }
  if arguments.NoteIds[0] > len(Notes.Notes) || arguments.NoteIds[0] < 0 {
    return errors.New("NoteId does not exist.")
  }
  Notes.Notes[arguments.NoteIds[0]].Text = arguments.Text[2:]
  Notes.Notes[arguments.NoteIds[0]].Tags = arguments.Tags
  Notes.Notes[arguments.NoteIds[0]].Deadline = arguments.Deadline
  writeErr := WriteNotebook(Notes)
  return writeErr
}

func AddComment(arguments Arguments) error {
  for index, note := range Notes.Notes {
    if containsInt(arguments.NoteIds, index) {
      Notes.Notes[index].Comments = append(note.Comments, arguments.Text)
    }
  }
  writeErr := WriteNotebook(Notes)
  return writeErr
}

func init()  {
  Notes, LoadErr = LoadNotebook() // NOTE: refactor this!
}

