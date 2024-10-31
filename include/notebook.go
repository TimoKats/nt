package include

import (
  "time"
)

var notebook Notebook

func ClearNotebook(arguments Arguments) error {
  notebook.Notes = []*Note{}
  writeErr := WriteNotebook()
  return writeErr
}

func ReadNotebook(arguments Arguments) error {
  for _, note := range notebook.Notes {
    if (hasOverlap(arguments.Tags, note.Tags) || len(arguments.Tags) == 0) {
      done := formatDone(note.Done)
      text := formatText(note.Text)
      tags := formatTags(note.Tags)
      Info.Println(note.Id, "-", done, text, "-", tags)
    }
  }
  return nil
}

func AddNote(arguments Arguments) error {
  note := Note{
    Id: len(notebook.Notes),
    Text: arguments.Text,
    Created: time.Now(),
    Tags: arguments.Tags,
  }
  notebook.Notes = append(notebook.Notes, &note)
  writeErr := WriteNotebook()
  return writeErr
}

func MoveNote(arguments Arguments) error {
  noteId, extractErr := extractInts(arguments.Text)
  if extractErr != nil {
    return extractErr
  }
  for index, note := range notebook.Notes {
    if note.Id == noteId {
      notebook.Notes[index].Done = !note.Done
    }
  }
  writeErr := WriteNotebook()
  return writeErr
}

