package main

import (
  shared "github.com/TimoKats/nt/include/shared"
  notebook "github.com/TimoKats/nt/include/notebook"
  server "github.com/TimoKats/nt/include/server"

  "errors"
  "os"
)

func run(arguments shared.Arguments) error {
  switch arguments.Command {
  case shared.Add:
    return notebook.AddNote(arguments)
  case shared.List:
    return notebook.ReadNotebook(arguments)
  case shared.Clear:
    return notebook.ClearNotebook(arguments)
  case shared.Move:
    return notebook.MoveNote(arguments)
  case shared.Comment:
    return notebook.AddComment(arguments)
  case shared.Tags:
    return notebook.ReadTags(arguments)
  case shared.Server:
    server.RunServer()
    return nil
  default:
    return errors.New("No valid command found. Use <<ls, add, mv>>")
  }
}

func main() {
  loadErr := notebook.LoadNotebook()
  if loadErr != nil {
    shared.Error.Println(loadErr)
    return
  }
  arguments := shared.ParseArgs(os.Args)
  runErr := run(arguments)
  if runErr != nil {
    shared.Error.Println(runErr)
  }
}
