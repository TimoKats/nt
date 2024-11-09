package include

import (
  . "github.com/TimoKats/nt/include/shared"

  "time"
)

var notebook Notebook

func noteSelected(index int, note *Note, arguments Arguments) bool {
  if hasOverlap(arguments.Tags, note.Tags) {
    return true
  } else if arguments.NoteId == index {
    return true
  } else if contains(arguments.Flags, "--done") && note.Done {
    return true
  } else if contains(arguments.Flags, "--old") && !fromToday(note) {
    return true
  }
  return false
}

func argumentsEmpty(arguments Arguments) bool {
  return (
    len(arguments.Text) == 0 &&
    len(arguments.Tags) == 0 &&
    len(arguments.Flags) == 0 &&
    arguments.NoteId == -1)
}

func ClearNotebook(arguments Arguments) error {
  correctedIndex := 0
  if argumentsEmpty(arguments) {
    notebook.Notes = []*Note{}
  } else {
    for index, note := range notebook.Notes {
      if noteSelected(index, note, arguments) {
        notebook.Notes = removeIndex(notebook.Notes, correctedIndex)
      } else {
        correctedIndex += 1
      }
    }
  }
  writeErr := WriteNotebook()
  return writeErr
}

func ReadNotebook(arguments Arguments) error {
  if argumentsEmpty(arguments) { formatSummaryHeader() }
  for index, note := range notebook.Notes {
    if noteSelected(index, note, arguments) || argumentsEmpty(arguments) {
      if arguments.NoteId == -1 {
        formatSummaryOutput(index, note)
      } else if arguments.NoteId == index {
        formatSingleOutput(note)
      }
    }
  }
  return nil
}

func ReadTags(arguments Arguments) error {
  tags := make(map[string]int)
  for _, note := range notebook.Notes {
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
  }
  notebook.Notes = append(notebook.Notes, &note)
  writeErr := WriteNotebook()
  return writeErr
}

func MoveNote(arguments Arguments) error {
  for index, note := range notebook.Notes {
    if arguments.NoteId == index {
      notebook.Notes[index].Done = !note.Done
    }
  }
  writeErr := WriteNotebook()
  return writeErr
}

func AddComment(arguments Arguments) error {
  for index, note := range notebook.Notes {
    if arguments.NoteId == index {
      notebook.Notes[index].Comments = append(note.Comments, arguments.Text)
    }
  }
  writeErr := WriteNotebook()
  return writeErr
}

