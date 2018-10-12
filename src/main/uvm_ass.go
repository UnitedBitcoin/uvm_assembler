//package main

//import (
//	"fmt"
//	"io/ioutil"
//	"uvmassembler"
//	"os"
//	"path/filepath"
//	"strings"
//)

//func main() {

//	if len(os.Args) < 2 {
//		fmt.Print("Usage: pass assemble source file as argument\n")
//		return
//	}
//	filename := os.Args[1]

//	fileBytes, _ := ioutil.ReadFile(filename)
//	fileContent := string(fileBytes)
//	asm := uvmassembler.NewAssembler()
//	sep := string(os.PathSeparator)
//	extension := filepath.Ext(filename)
//	outfilePath := filepath.Dir(filename) + sep + strings.TrimSuffix(filepath.Base(filename), extension) + ".out"
//	fmt.Print(outfilePath + "\n")
//	asm.ParseAsmContent(fileContent, outfilePath)
//	fmt.Printf("assemble %s done\n", filename)
//}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"uvmassembler"
)

func main() {
	argsNum := len(os.Args)

	if argsNum == 2 { //.uvm -> .out
		arg := os.Args[1]
		if arg != "-h" && arg != "-help" {
			filename := arg
			fmt.Printf("assemble %s begin--------\n", filename)
			fileBytes, _ := ioutil.ReadFile(filename)
			fileContent := string(fileBytes)
			asm := uvmassembler.NewAssembler()
			sep := string(os.PathSeparator)
			extension := filepath.Ext(filename)
			outfilePath := filepath.Dir(filename) + sep + strings.TrimSuffix(filepath.Base(filename), extension) + ".out"
			fmt.Print(outfilePath + "\n")
			asm.ParseAsmContent(fileContent, outfilePath)
			fmt.Printf("assemble %s done\n", filename)
			return
		}
	} else if argsNum == 3 { // .out -> .uvmasm
		arg := os.Args[1]
		if arg == "-dis" {
			sep := string(os.PathSeparator)
			bytecodefilename := os.Args[2]
			fmt.Printf("disassemble %s begin--------\n", bytecodefilename)
			extension := filepath.Ext(bytecodefilename)
			outfileAsmPath := filepath.Dir(bytecodefilename) + sep + strings.TrimSuffix(filepath.Base(bytecodefilename), extension) + ".uvmasm"
			disasm := uvmassembler.NewDisAssembler()

			fmt.Print(outfileAsmPath + "\n")
			res, err := disasm.ParseByteCode(bytecodefilename)
			if !res {
				fmt.Println(err)
				return
			}
			res, err = disasm.WriteAsm(outfileAsmPath)
			if !res {
				fmt.Println("error: " + err)
				return
			}
			fmt.Printf("disassemble %s done\n", outfileAsmPath)
			return
		}
	}
	fmt.Print("Usage 1:\n")
	fmt.Print("\ttranslate assemble file to bytecode file\n")
	fmt.Print("\targuments: assembleFile\n")
	fmt.Print("Usage 2:\n")
	fmt.Print("\ttranslate bytecode file to assemble file\n")
	fmt.Print("\targuments: -dis bytecodeFile\n")
}
