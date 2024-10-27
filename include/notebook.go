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
    Info.Println(note.Id, note.Text, note.Tags, note.Status)
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
