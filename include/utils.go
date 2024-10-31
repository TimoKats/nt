package include

import (
  "regexp"
  "strconv"
  "strings"
)

func hasOverlap(list1 []string, list2 []string) bool {
  for _, item1 := range list1 {
    for _, item2 := range list2 {
      if item1 == item2 {
        return true
      }
    }
  }
  return false
}

func contains(slice []string, item string) bool {
  for _, v := range slice {
    if v == item {
      return true
    }
  }
  return false
}

func extractInts(input string) (int, error) {
  extractedInts := ""
  re := regexp.MustCompile(`\d+`)
  matches := re.FindAllString(input, -1)
  for _, match := range matches {
    extractedInts += match
  }
  num, convErr := strconv.Atoi(extractedInts)
  if convErr != nil {
    return 0, convErr
  }
  return num, nil
}

func responsiveWhitespace(text string, columnwidth int) string {
  if len(text) > columnwidth {
    text = text[:columnwidth]
  }
  spaces := columnwidth - len(text)
  for i := 0; i < spaces; i++ {
    text += " "
  }
  return text
}

func formatDone(taskDone bool) string {
  if taskDone {
    return "[X]"
  }
  return "[ ]"
}

func formatText(text string) string {
  text = strings.Replace(text, "\n", " ", -1)
  maxLength := 25
  if len(text) > maxLength {
    return text[:maxLength] + "..."
  }
  spaces := (maxLength - len(text)) + 3
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

