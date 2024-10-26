package include

func ReadNotebook(arguments Arguments) error {
  for tag, notes := range notebook.TagIndex {
    if contains(arguments.Tags, tag) || len(arguments.Tags) == 0 {
      Info.Println(tag)
      for _, note := range notes {
        Info.Printf("  %s", note.Text)
      }
    }
  }
  return nil
}

