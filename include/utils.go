package include

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

