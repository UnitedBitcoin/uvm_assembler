package main

import (
	"fmt"
	"io/ioutil"
	"uvmassembler"
  "os"
  "path/filepath"
  "strings"
)

func main() {

  if len(os.Args)<2 {
    fmt.Print("Usage: pass assemble source file as argument\n")
    return
  }
  filename := os.Args[1]

	fileBytes, _ := ioutil.ReadFile(filename)
	fileContent := string(fileBytes)
	asm := uvmassembler.NewAssembler()
  sep := string(os.PathSeparator)
  extension := filepath.Ext(filename)
  outfilePath := filepath.Dir(filename) + sep + strings.TrimSuffix(filepath.Base(filename), extension) + ".out"
  fmt.Print(outfilePath + "\n")
	asm.ParseAsmContent(fileContent, outfilePath)
  fmt.Printf("assemble %s done", filename)
}
