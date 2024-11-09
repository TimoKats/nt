package include

import (
  . "github.com/TimoKats/nt/include/shared"

  "github.com/atotto/clipboard"
  "strings"
)

func formatDone(taskDone bool) string {
  if taskDone {
    return "[X]"
  }
  return "[ ]"
}

func formatSummaryText(text string) string {
  text = strings.Replace(text, "\n", " ", -1)
  maxLength := 30
  if len(text) > maxLength {
    return text[:maxLength]
  }
  spaces := (maxLength - len(text))
  for i := 0; i < spaces; i++ {
    text += " "
  }
  return text
}

func formatTags(tags []string) string {
  formattedTags := ""
  for index, tag := range tags {
    if index < len(tags) - 1 {
      formattedTags += tag + ", "
    } else {
      formattedTags += tag
    }
  }
  return formattedTags
}

func formatSummaryOutput(index int, note *Note) {
  done := formatDone(note.Done)
  text := formatSummaryText(note.Text)
  tags := formatTags(note.Tags)
  Info.Println(index, "  ", done, "  ", text, "  ", tags)
}

func formatSummaryHeader() {
  Info.Println("Id   Done   Text                              Tags")
  Info.Println("---- ------ --------------------------------- ---------")
}

func formatSingleOutput(note *Note) {
  clipboard.WriteAll(note.Text)
  Info.Println(note.Text)
  Info.Println("\n---------\nComments:")
  for _, comment := range note.Comments {
    Info.Println("-", comment[2:])
  }
}

func FormatInfo() error {
  Info.Println(`Commands:
    - test
    - test
  `)
  return nil
}

