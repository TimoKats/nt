package include

import (
  "github.com/atotto/clipboard"

  . "github.com/TimoKats/nt/include/shared"

  "strconv"
  "strings"
  "time"
)

func formatDone(taskDone bool) string {
  if taskDone {
    return "[X]"
  }
  return "[ ]"
}

func formatSummaryText(text string) string {
  text = strings.Replace(text, "\n", " ", -1)
  maxLength := NtConfig.Notebook.Width + 1
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

func formatId(index int) string {
  var strId string = strconv.Itoa(index)
  for len(strId) < 3 {
    strId += " "
  }
  return strId
}

func formatDeadline(deadline time.Time) string {
  if !deadline.IsZero() {
    return deadline.Format("2006-01-02")
  }
  return "          "
}

func formatSummaryOutput(index int, note *Note) {
  id := formatId(index)
  done := formatDone(note.Done)
  text := formatSummaryText(note.Text)
  tags := formatTags(note.Tags)
  deadline := formatDeadline(note.Deadline)
  Info.Println(id, "  ", done, "  ", text, "  ", deadline, " ", tags)
}

func formatSummaryHeader() {
  headerWidth := strings.Repeat(" ", NtConfig.Notebook.Width)
  seperatorWidth := strings.Repeat("-", NtConfig.Notebook.Width)
  Info.Printf("Id   Done   Text%s Deadline   Tags ", headerWidth)
  Info.Printf("---- ------ ----%s ---------- -----------", seperatorWidth)
}

func formatSingleOutput(note *Note) {
  clipboard.WriteAll(note.Text)
  Info.Println(note.Text)
  Info.Println("\n---------\nComments:")
  for _, comment := range note.Comments {
    Info.Println("-", comment[2:])
  }
}

