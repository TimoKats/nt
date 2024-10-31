package main

import (
  nt "github.com/TimoKats/nt/include"
  "errors"
  "os"
)

func run(arguments nt.Arguments) error {
  switch arguments.Command {
  case nt.Add:
    return nt.AddNote(arguments)
  case nt.List:
    return nt.ReadNotebook(arguments)
  case nt.Clear:
    return nt.ClearNotebook(arguments)
  case nt.Move:
    return nt.MoveNote(arguments)
  default:
    return errors.New("No valid command found. Use <<ls, add, mv>>")
  }
}

func main() {
  loadErr := nt.LoadNotebook()
  if loadErr != nil {
    nt.Error.Println(loadErr)
    return
  }
  arguments := nt.ParseArgs(os.Args)
  runErr := run(arguments)
  if runErr != nil {
    nt.Error.Println(runErr)
  }
}
