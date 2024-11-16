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
  in := bufio.NewReader(os.Stdin)
  fmt.Print(Cyan + message + Reset)
  command, _ := in.ReadBytes('\n')
  return command[0:len(command)-1]
}

func SecureInput(message string) []byte {
  fmt.Print(Cyan + message + Reset)
  bytepw, _ := term.ReadPassword(int(syscall.Stdin))
  fmt.Print("\n")
  return bytepw
}
