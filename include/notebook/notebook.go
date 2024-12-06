package include

import (
  . "github.com/TimoKats/nt/include/shared"

  "strings"
  "errors"
  "time"
)

func noteSelected(index int, note *Note, arguments Arguments) bool {
  if hasOverlap(arguments.Tags, note.Tags) {
    return true
  } else if containsInt(arguments.NoteIds, index) {
    return true
  } else if containsStr(arguments.Flags, "--done") && note.Done {
    return true
  } else if containsStr(arguments.Flags, "--old") && !fromToday(note) {
    return true
  } else if containsStr(arguments.Flags, "--today") && fromToday(note) {
    return true
  } else if containsStr(arguments.Flags, "--all") {
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
  var correctedIndex int = 0
  if argumentsEmpty(arguments) {
    NtNotes.Notes = []*Note{}
  } else {
    for index, note := range NtNotes.Notes {
      if noteSelected(index, note, arguments) {
        NtNotes.Notes = removeIndex(NtNotes.Notes, correctedIndex)
      } else {
        correctedIndex += 1
      }
    }
  }
  writeErr := WriteNotebook(NtNotes)
  return writeErr
}

func ListNotebook(arguments Arguments) error {
  if len(arguments.NoteIds) != 1 { formatSummaryHeader() }
  for index, note := range NtNotes.Notes {
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
  for _, note := range NtNotes.Notes {
    for _, tag := range note.Tags {
      tags[tag] += 1
    }
  }
  for tagname := range tags {
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
  NtNotes.Notes = append(NtNotes.Notes, &note)
  var writeErr error = WriteNotebook(NtNotes)
  return writeErr
}

func MoveNote(arguments Arguments) error {
  for index, note := range NtNotes.Notes {
    if containsInt(arguments.NoteIds, index) {
      NtNotes.Notes[index].Done = !note.Done
    }
  }
  var writeErr error = WriteNotebook(NtNotes)
  return writeErr
}

func SearchNote(arguments Arguments) error {
  arguments.Flags = []string{}
  for index, note := range NtNotes.Notes {
    for _, query := range strings.Fields(arguments.Text) {
      if strings.Contains(note.Text, query) {
        arguments.NoteIds = append(arguments.NoteIds, index); break
      }
    }
  }
  var listErr error = ListNotebook(arguments)
  return listErr
}

func ModifyNote(arguments Arguments) error {
  if len(arguments.NoteIds) == 0 {
    return errors.New("No NoteId provided")
  }
  if arguments.NoteIds[0] > len(NtNotes.Notes) || arguments.NoteIds[0] < 0 {
    return errors.New("NoteId does not exist.")
  }
  NtNotes.Notes[arguments.NoteIds[0]].Text = arguments.Text[2:]
  NtNotes.Notes[arguments.NoteIds[0]].Tags = arguments.Tags
  NtNotes.Notes[arguments.NoteIds[0]].Deadline = arguments.Deadline
  var writeErr error = WriteNotebook(NtNotes)
  return writeErr
}

func AddComment(arguments Arguments) error {
  for index, note := range NtNotes.Notes {
    if containsInt(arguments.NoteIds, index) {
      NtNotes.Notes[index].Comments = append(note.Comments, arguments.Text)
    }
  }
  var writeErr error = WriteNotebook(NtNotes)
  return writeErr
}

