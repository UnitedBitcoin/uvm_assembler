package uvmassembler

import (
	"fmt"
	"testing"
  //"text/scanner"
  //"strings"
  "io/ioutil"
  "strings"
  "os"
)

func TestParseInt(t *testing.T) {
  //const src = `
  //    5 0 2 hello "world" ; hello world
  //`
  //var s scanner.Scanner
  //// Don't skip comments: we need to count newlines.
  //s.Mode = scanner.ScanChars |
  //scanner.ScanFloats |
  // 	scanner.ScanIdents |
  // scanner.ScanInts |
  // scanner.ScanStrings |
  //  scanner.ScanComments
  ////s.Filename = "test1.luas"
  //s.Init(strings.NewReader(src))
  //var tok rune
  //for tok != scanner.EOF {
  //  tok = s.Scan()
  //  fmt.Println("At position", s.Pos(), " Line ", s.Line, ":", s.TokenText())
  //}

	line := "    5 0 2 hello \"world\"; hello"

	res1, pos1, value1 := ParseInt(line, 0, len(line))
  fmt.Printf("d1: pos %d value %d\n", pos1, value1)
	if !res1 {
		t.Error("parse int error")
	}
  res2, pos2, value2 := ParseInt(line, pos1, len(line))
  fmt.Printf("d2: pos %d value %d\n", pos2, value2)
  if !res2 {
    t.Error("parse int error")
  }
  res3, pos3, value3 := ParseInt(line, pos2, len(line))
  fmt.Printf("d3: pos %d value %d\n", pos3, value3)
  if !res3 {
    t.Error("parse int error")
  }

  res4, pos4, value4 := ParseInt(line, pos3, len(line))
  fmt.Printf("d4: pos %d value %d\n", pos4, value4)
  if res4 {
    t.Error("parse empty int error")
  }

  res5, pos5, value5 := ParseLabel(line, pos3, len(line))
  fmt.Printf("d4: pos %d value %s\n", pos5, value5)
  if !res5 {
    t.Error("parse label error")
  }
}

func TestParseLabel(t *testing.T) {
  line := "    hello world"
  res1, pos1, value1 := ParseLabel(line, 0, len(line))
  fmt.Printf("d1: pos %d value %s\n", pos1, value1)
  if !res1 {
    t.Error("parse int error")
  }
  res2, pos2, value2 := ParseLabel(line, pos1, len(line))
  fmt.Printf("d2: pos %d value %s\n", pos2, value2)
  if !res2 {
    t.Error("parse int error")
  }
}

func TestParseOperand(t *testing.T) {
  line1 := " closure %1 subroutine_2"
  //line2:= "call %2 3 1"
  res1, pos1, value1 := ParseLabel(line1, 0, len(line1))
  fmt.Printf("cmd: pos %d value %s\n", pos1, value1)
  if !res1 {
    t.Error("parse label error")
  }
  asm := NewAssembler()
  var op1 Operand
  res2, pos2, err2 := asm.ParseOperand(&op1, line1[pos1:], LIMIT_STACKIDX)
  fmt.Printf("op: pos: %d value: %d error: %s\n", pos2, op1.value, err2)
  if !res2 {
    t.Error("parse label error")
  }

  var op2 Operand
  res3, pos3, err3 := asm.ParseOperand(&op2, line1[pos1+pos2:], LIMIT_PROTO)
  fmt.Printf("op: pos: %d value: %d error: %s\n", pos3, op2.value, err3)
  if !res3 {
    t.Error("parse label error")
  }

}

func TestParseLine(t *testing.T) {
  fileBytes, _ := ioutil.ReadFile("D:\\projects\\lua_assembler\\test1.luas")
  fileContent := string(fileBytes)
  lines := strings.Split(fileContent, "\n")
  asm := NewAssembler()

  for i := range lines {
    line := lines[i]
    line = Trim(line)
    if len(line) < 1 {
      continue
    }
    res, ParseLineError := asm.ParseLine(line, len(line))
    if !res {
      t.Error(ParseLineError + ".\t" + line)
    }
    fmt.Printf("line %d is %s\n", (i + 1), line)
  }
}

func TestParseUpvalue(t *testing.T) {

}

func TestAssembler(t *testing.T) {
  fmt.Printf("hello world\n")

  wd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  sep := string(os.PathSeparator)
  fileBytes, _ := ioutil.ReadFile(wd + sep + ".." + sep + ".." + sep + "test1.uvms")
  fileContent := string(fileBytes)
  //fmt.Print(fileContent)
  asm := NewAssembler()
  outFilepath := wd + sep + ".." + sep + ".." + sep + "test1.out"
  asm.ParseAsmContent(fileContent, outFilepath)
}
