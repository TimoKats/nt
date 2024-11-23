package include

import(
  "golang.org/x/term"
  "syscall"
  "bufio"
  "fmt"
  "os"
)

var Cyan = "\033[34m"

func InsecureInput(message string) []byte {
  var command []byte
  var readErr error
  var in *bufio.Reader = bufio.NewReader(os.Stdin)
  fmt.Print(Cyan + message + Reset)
  command, readErr = in.ReadBytes('\n')
  if readErr != nil { Error.Println(readErr) }
  return command[0:len(command)-1]
}

func SecureInput(message string) []byte {
  var byteInput []byte
  var readErr error
  fmt.Print(Cyan + message + Reset)
  byteInput, readErr = term.ReadPassword(int(syscall.Stdin))
  if readErr != nil { Error.Println(readErr) }
  fmt.Print("\n")
  return byteInput
}
