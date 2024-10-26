package include

import (
  "errors"
  "time"
)

func indexNote(note Note) error {
  if notebook.TagIndex == nil {
    notebook.TagIndex = make(map[string][]*Note)
  }
  if len(note.Tags) == 0 {
    notebook.TagIndex["main"] = append(notebook.TagIndex["main"], &note)
  } else {
    for _, tag := range note.Tags {
      notebook.TagIndex[tag] = append(notebook.TagIndex[tag], &note)
    }
  }
  return nil
}

func WriteNote(arguments Arguments) error {
  note := Note{
    Id: randomString(5),
    Text: arguments.Text,
    Created: time.Now(),
    Tags: arguments.Tags,
  }
  indexErr := indexNote(note)
  writeErr := WriteNotebook()
  return errors.Join(indexErr, writeErr)
}
