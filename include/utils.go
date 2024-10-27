package include

func contains(slice []string, item string) bool {
  for _, v := range slice {
    if v == item {
      return true
    }
  }
  return false
}
