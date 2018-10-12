package uvmassembler

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

type Head struct {
	uvm_signature        []byte
	uvmc_version         int
	uvmc_format          int
	uvmc_data            []byte
	uvm_int_size         int
	uvm_size_t_size      int
	uvm_instruction_size int
	uvm_integer_size     int
	uvm_number_size      int
}

type DisAssembler struct {
	topFunc     *ParsedFunction
	nUpvals     int
	byteNum     int
	headContent *Head
	fileBytes   []byte
	readCur     int
	functions   map[string]*ParsedFunction
}

func (disassembler *DisAssembler) DisAssemble() (bool, string) {
	return true, ""
}

func NewDisAssembler() *DisAssembler {
	instance := new(DisAssembler)
	instance.nUpvals = 0
	instance.byteNum = 0
	instance.readCur = 0
	instance.topFunc = nil
	instance.headContent = nil
	instance.functions = make(map[string]*ParsedFunction)

	return instance
}

func (disassembler *DisAssembler) ParseHead() (bool, string) {
	sig := []byte(LUA_SIGNATURE)

	if disassembler.byteNum < len(sig) {
		return false, "not bytecode file"
	}
	parsedSig := disassembler.fileBytes[:len(sig)]
	if !bytes.Equal(sig, parsedSig) {
		return false, "not bytecode file, head SIGNATURE wrong"
	}
	disassembler.headContent = new(Head)
	disassembler.readCur += len(sig)
	disassembler.headContent.uvm_signature = make([]byte, len(sig))
	copy(disassembler.headContent.uvm_signature, sig)

	if byte(LUAC_VERSION) != disassembler.fileBytes[disassembler.readCur] {
		return false, "head version wrong"
	}
	disassembler.readCur += 1
	disassembler.headContent.uvmc_version = LUAC_VERSION

	if byte(LUAC_FORMAT) != disassembler.fileBytes[disassembler.readCur] {
		return false, "head format wrong"
	}
	disassembler.readCur += 1
	disassembler.headContent.uvmc_format = LUAC_FORMAT

	cdata := []byte(LUAC_DATA)
	parsedCdata := disassembler.fileBytes[disassembler.readCur:(disassembler.readCur + len(cdata))]
	if !bytes.Equal(cdata, parsedCdata) {
		return false, "head cdata wrong"
	}
	disassembler.readCur += len(cdata)
	disassembler.headContent.uvmc_data = make([]byte, len(cdata))
	copy(disassembler.headContent.uvmc_data, cdata)

	if byte(4) != disassembler.fileBytes[disassembler.readCur] {
		return false, "head int size wrong"
	}
	disassembler.readCur += 1
	disassembler.headContent.uvm_int_size = 4

	if byte(LUA_SIZE_T_TYPE_SIZE) != disassembler.fileBytes[disassembler.readCur] {
		return false, "head size_t size wrong"
	}
	disassembler.readCur += 1
	disassembler.headContent.uvm_size_t_size = 4

	if byte(4) != disassembler.fileBytes[disassembler.readCur] {
		return false, "head instruction size wrong"
	}
	disassembler.readCur += 1
	disassembler.headContent.uvm_instruction_size = 4

	if byte(LUA_INTEGER_TYPE_SIZE) != disassembler.fileBytes[disassembler.readCur] {
		return false, "head integer size wrong"
	}
	disassembler.readCur += 1
	disassembler.headContent.uvm_integer_size = LUA_INTEGER_TYPE_SIZE

	if byte(LUA_NUMBER_TYPE_SIZE) != disassembler.fileBytes[disassembler.readCur] {
		return false, "head number size wrong"
	}
	disassembler.readCur += 1
	disassembler.headContent.uvm_number_size = LUA_NUMBER_TYPE_SIZE

	////////////////
	if LUA_INTEGER_TYPE_SIZE == 4 {
		int32v := binary.LittleEndian.Uint32(disassembler.fileBytes[disassembler.readCur:(disassembler.readCur + 4)])
		if LUAC_INT != int32v {
			return false, "head endianness mismatch"
		}
		disassembler.readCur += 4
	} else if LUA_INTEGER_TYPE_SIZE == 8 {
		int64v := binary.LittleEndian.Uint64(disassembler.fileBytes[disassembler.readCur:(disassembler.readCur + 8)])
		if LUAC_INT != int64v {
			return false, "head endianness mismatch"
		}
		disassembler.readCur += 8
	} else {
		return false, "unsupported lua_Integer size " + strconv.Itoa(LUA_INTEGER_TYPE_SIZE)
	}

	if LUA_NUMBER_TYPE_SIZE == 4 {
		floatbit32s := math.Float32bits(LUAC_NUM)
		bits32 := binary.LittleEndian.Uint32(disassembler.fileBytes[disassembler.readCur:(disassembler.readCur + 4)])
		if floatbit32s != bits32 {
			return false, "head float format mismatch"
		}
	} else if LUA_NUMBER_TYPE_SIZE == 8 {
		floatbit64s := math.Float64bits(LUAC_NUM)
		bits64 := binary.LittleEndian.Uint64(disassembler.fileBytes[disassembler.readCur:(disassembler.readCur + 8)])
		if floatbit64s != bits64 {
			return false, "head float format mismatch"
		}
		disassembler.readCur += 8
	} else {
		return false, "unsupported lua_Number size " + strconv.Itoa(LUA_NUMBER_TYPE_SIZE)
	}

	return true, ""
}

//-----------------util--------------
func (disassembler *DisAssembler) LoadByte() byte {
	byteval := disassembler.fileBytes[disassembler.readCur]
	disassembler.readCur++
	return byteval
}

func (disassembler *DisAssembler) LoadUint32() uint32 {
	uintval := binary.LittleEndian.Uint32(disassembler.fileBytes[disassembler.readCur:(disassembler.readCur + 4)])
	disassembler.readCur += 4
	return uintval
}

func (disassembler *DisAssembler) LoadUint64() uint64 {
	uintval := binary.LittleEndian.Uint64(disassembler.fileBytes[disassembler.readCur:(disassembler.readCur + 8)])
	disassembler.readCur += 8
	return uintval
}

func (disassembler *DisAssembler) LoadNumber() float64 {
	uintval := binary.LittleEndian.Uint64(disassembler.fileBytes[disassembler.readCur:(disassembler.readCur + 8)])
	floatval := math.Float64frombits(uintval)
	disassembler.readCur += 8
	return floatval
}

func (disassembler *DisAssembler) LoadString() string {
	byteval := disassembler.fileBytes[disassembler.readCur]
	size := uint64(byteval)
	disassembler.readCur++
	if byteval == byte(0xFF) {
		size = binary.LittleEndian.Uint64(disassembler.fileBytes[disassembler.readCur:(disassembler.readCur + 8)])
		disassembler.readCur += 8
	}

	if size == 0 {
		return ""
	} else {
		size--
		strval := string(disassembler.fileBytes[disassembler.readCur:(disassembler.readCur + int(size))])
		disassembler.readCur += int(size)
		return strval
	}
}

//----------------code-----------------
func (disassembler *DisAssembler) LoadCode(pf *ParsedFunction) (bool, string) {
	n := disassembler.LoadUint32()
	pf.instructions = make([]Instruction, n)
	for i := uint32(0); i < n; i++ {
		pf.instructions[i] = Instruction(disassembler.LoadUint32())
	}
	if disassembler.readCur > disassembler.byteNum {
		return false, "bytecode file not complete"
	}
	return true, ""

}

func (disassembler *DisAssembler) PreParseInstructionLocation(pf *ParsedFunction) (bool, string, map[int]string) {
	var ins Instruction
	var jmpdest int
	var operand uint
	locationsinfo := make(map[int]string)
	for i := 0; i < len(pf.instructions); i++ {
		ins = pf.instructions[i]
		opcode := GET_OPCODE(ins)

		if int(opcode) >= NUM_OPCODES {
			return false, "unkown opcode", nil
		}
		count := Opcounts[opcode]
		info := Opinfos[opcode]
		for j := 0; j < count; j++ {
			if info[j].limit == LIMIT_LOCATION {
				operand = GETARG_sBx(ins)
				jmpdest = int(operand) + 1 + i
				if (jmpdest >= len(pf.instructions)) || (jmpdest < 0) {
					return false, "jmp dest exceed", nil
				}
				insLabel := pf.name + "_to_dest_br_" + strconv.Itoa(jmpdest)
				locationsinfo[jmpdest] = insLabel
			}
		}
	}

	return true, "", locationsinfo

}

//returns : res, err, insstr, useExtended, extraArg
func (disassembler *DisAssembler) ParseInstruction(pf *ParsedFunction, insIdx int, ins Instruction, locationifos map[int]string) (bool, string, string, bool, bool) {
	opcode := GET_OPCODE(ins)
	if int(opcode) >= NUM_OPCODES {
		return false, "unkown opcode", "", false, false
	}
	opname := UvmPOpnames[opcode]
	count := Opcounts[opcode]
	info := Opinfos[opcode]

	var insstr string = opname
	var operand uint
	var limit int
	//var opValue uint
	var constval string
	var consttype int
	var useExtended bool = false
	var extraArg bool = false

	//{{OPP_Ax, LIMIT_EMBED}}
	if opcode == OP_EXTRAARG {
		extraArg = true
		insstr = ""
	}

	for i := 0; i < count; i++ {
		limit = info[i].limit
		switch info[i].pos {
		case OPP_A:
			operand = GETARG_A(ins)
		case OPP_B:
			operand = GETARG_B(ins)
		case OPP_C:
			operand = GETARG_C(ins)
		case OPP_Bx:
			operand = GETARG_Bx(ins)
		case OPP_Ax:
			operand = GETARG_Ax(ins)
		case OPP_sBx:
			operand = GETARG_sBx(ins)
		case OPP_ARG:
			useExtended = true
		case OPP_C_ARG:
			operand = GETARG_C(ins)
			if operand == 0 {
				useExtended = true //try get next
			}
		}

		if useExtended { //try get next ins
			return true, "", insstr, useExtended, false
		}

		switch limit {
		case LIMIT_STACKIDX:
			insstr = insstr + " %" + strconv.Itoa(int(operand))
		case LIMIT_UPVALUE:
			insstr = insstr + " @" + strconv.Itoa(int(operand))
		case LIMIT_EMBED:
			insstr = insstr + " " + strconv.Itoa(int(operand))
		case LIMIT_CONSTANT:
			constval = (*pf.constants[operand]).str()
			consttype = (*pf.constants[operand]).valueType()
			if consttype == LUA_TSTRING || consttype == LUA_TLNGSTR {
				constval = "\"" + constval + "\""
			}
			insstr = insstr + " const " + constval
		case LIMIT_LOCATION:
			loclabel := locationifos[int(operand)+1+insIdx]
			insstr = insstr + " $" + loclabel
		case LIMIT_CONST_STACK:
			if (int(operand) & BITRK) > 0 {
				constval = (*pf.constants[int(operand)-BITRK]).str()
				consttype = (*pf.constants[int(operand)-BITRK]).valueType()
				if consttype == LUA_TSTRING || consttype == LUA_TLNGSTR {
					constval = "\"" + constval + "\""
				}
				insstr = insstr + " const " + constval
			} else {
				insstr = insstr + " %" + strconv.Itoa(int(operand))
			}

		case LIMIT_PROTO:
			propname := pf.usedSubroutines[operand]
			insstr = insstr + " " + propname
		}

	}

	return true, "", insstr, false, extraArg

}

func (disassembler *DisAssembler) LoadConstants(pf *ParsedFunction) (bool, string) {
	n := disassembler.LoadUint32()
	pf.constants = pf.constants[:0]
	var tval TValue
	var strval *TString
	for i := 0; i < int(n); i++ {
		t := int(disassembler.LoadByte())
		switch t {
		case LUA_TNIL:
			{
				nilval := new(TNil)
				tval = nilval
				temp := tval
				pf.constants = append(pf.constants, &temp)
			}
		case LUA_TBOOLEAN:
			{
				boolval := new(TBool)
				val := uint(disassembler.LoadByte())
				if val == 0 {
					boolval.bool_value = false
				} else {
					boolval.bool_value = true
				}
				tval = boolval
				temp := tval
				pf.constants = append(pf.constants, &temp)
			}
		case LUA_TNUMFLT:
			{
				numberval := new(TNumber)
				numberval.number_value = disassembler.LoadNumber()
				tval = numberval
				temp := tval //new temp prevent tval gc
				pf.constants = append(pf.constants, &temp)
			}
		case LUA_TNUMINT:
			{
				integerval := new(TInteger)
				integerval.int_value = int64(disassembler.LoadUint64())
				tval = integerval
				temp := tval
				pf.constants = append(pf.constants, &temp)
			}
		case LUA_TSTRING:
			{
				strval = new(TString)
				strval.string_value = disassembler.LoadString()
				tval = strval
				temp := tval
				pf.constants = append(pf.constants, &temp)
			}
		case LUA_TLNGSTR:
			{
				strval = new(TString)
				strval.string_value = disassembler.LoadString()
				tval = strval
				temp := tval
				pf.constants = append(pf.constants, &temp)
			}
		default:
			return false, "unkown constant type"
		}

	}

	if disassembler.readCur > disassembler.byteNum {
		return false, "bytecode file not complete"
	}

	return true, ""

}

func (disassembler *DisAssembler) LoadProtos(pf *ParsedFunction) (bool, string) {
	n := disassembler.LoadUint32()

	for i := 0; i < int(n); i++ {
		newfunc := new(ParsedFunction)
		res, err := disassembler.LoadFunc(newfunc, pf, i)
		if !res {
			return res, err
		}
		disassembler.functions[newfunc.name] = newfunc
		pf.usedSubroutines = append(pf.usedSubroutines, newfunc.name)
	}

	if disassembler.readCur > disassembler.byteNum {
		return false, "bytecode file not complete"
	}
	return true, ""
}

func (disassembler *DisAssembler) LoadUpvalues(pf *ParsedFunction) (bool, string) {
	n := disassembler.LoadUint32()
	var upvalue Upvalue
	for i := 0; i < int(n); i++ {
		upvalue.instack = uint8(disassembler.LoadByte())
		upvalue.idx = uint8(disassembler.LoadByte())
		upvalue.name = ""
		pf.upvalues = append(pf.upvalues, upvalue)
	}
	if disassembler.readCur > disassembler.byteNum {
		return false, "bytecode file not complete"
	}
	return true, ""
}

func (disassembler *DisAssembler) LoadDebug(pf *ParsedFunction) (bool, string) {
	n := disassembler.LoadUint32()
	for i := 0; i < int(n); i++ {
		var lineinfo int = int(disassembler.LoadUint32())
		pf.lineinfos = append(pf.lineinfos, lineinfo)
	}

	n = disassembler.LoadUint32()
	var locval LocVar
	for i := 0; i < int(n); i++ {
		locval.varname = disassembler.LoadString()
		locval.startpc = int(disassembler.LoadUint32())
		locval.endpc = int(disassembler.LoadUint32())
		pf.locals = append(pf.locals, locval)
	}

	n = disassembler.LoadUint32()
	for i := 0; i < int(n); i++ {
		pf.upvalues[i].name = disassembler.LoadString()
	}
	if disassembler.readCur > disassembler.byteNum {
		return false, "bytecode file not complete"
	}
	return true, ""
}

//------------------------------
func (disassembler *DisAssembler) LoadFunc(pf *ParsedFunction, parent *ParsedFunction, idx int) (bool, string) {
	name := disassembler.LoadString()
	if name == "" {
		if parent == nil {
			name = "main_fake"
		} else {
			name = parent.name + "_" + strconv.Itoa(idx) + "_fake"
		}
	}

	pf.name = name

	pf.linedefined = uint(disassembler.LoadUint32())
	pf.lastlinedefined = uint(disassembler.LoadUint32())

	pf.params = uint(disassembler.LoadByte())
	pf.vararg = uint(disassembler.LoadByte())
	pf.maxstacksize = uint(disassembler.LoadByte())

	res, err := disassembler.LoadCode(pf)
	if !res {
		return res, err
	}

	res, err = disassembler.LoadConstants(pf)
	if !res {
		return res, err
	}

	res, err = disassembler.LoadUpvalues(pf)
	if !res {
		return res, err
	}

	res, err = disassembler.LoadProtos(pf)
	if !res {
		return res, err
	}

	res, err = disassembler.LoadDebug(pf)
	if !res {
		fmt.Println("debug info not found")
	}

	if disassembler.readCur > disassembler.byteNum {
		return false, "trucated bytecode file"
	}

	return true, ""
}

//-----------------------------------------------------------
func (disassembler *DisAssembler) WriteAsm(outfilePath string) (bool, string) {
	outfile, err := os.Create(outfilePath)
	if err != nil {
		return false, err.Error()
	}
	res, errinfo := disassembler.WriteFuncAsm(outfile, disassembler.topFunc, true)
	if !res {
		outfile.Close()
		return res, errinfo
	}
	outfile.Close()
	return true, ""
}

func (disassembler *DisAssembler) WriteFuncAsm(outfile *os.File, pf *ParsedFunction, isTop bool) (bool, string) {
	if isTop {
		outfile.WriteString(".upvalues " + strconv.Itoa(disassembler.nUpvals) + "\r\n")
	}

	outfile.WriteString(".func " + pf.name + " " + strconv.Itoa(int(pf.maxstacksize)) + " " + strconv.Itoa(int(pf.params)) + " " + strconv.Itoa(len(pf.locals)) + "\r\n")

	//write constants  --------------------------------------
	outfile.WriteString(".begin_const\r\n")
	var valtype int
	var val string
	for i := 0; i < len(pf.constants); i++ {
		constantValue := *(pf.constants[i])
		outfile.WriteString("\t")
		valtype = constantValue.valueType()
		val = constantValue.str()
		switch valtype {
		case LUA_TSTRING:
			outfile.WriteString("\"" + val + "\"\r\n")
		case LUA_TLNGSTR:
			outfile.WriteString("\"" + val + "\"\r\n")
		default:
			outfile.WriteString(constantValue.str() + "\r\n")
		}
	}
	outfile.WriteString(".end_const\r\n")

	//write upvalue -------------------------------------------
	outfile.WriteString(".begin_upvalue\r\n")
	for i := 0; i < len(pf.upvalues); i++ {
		upvalue := pf.upvalues[i]
		outfile.WriteString("\t" + strconv.Itoa(int(upvalue.instack)) + " " + strconv.Itoa(int(upvalue.idx)) + " \"" + upvalue.name + "\"\r\n")
	}
	outfile.WriteString(".end_upvalue\r\n")

	//write locals  ---------------------------------------------
	var local LocVar
	outfile.WriteString(".begin_local\r\n")
	for i := 0; i < len(pf.locals); i++ {
		local = pf.locals[i]
		outfile.WriteString("\t\"" + local.varname + "\" " + strconv.Itoa(local.startpc) + " " + strconv.Itoa(local.endpc) + "\r\n")
	}
	outfile.WriteString(".end_local\r\n")

	//pre parse location  ----------------------------------------
	res, err, locationinfos := disassembler.PreParseInstructionLocation(pf)
	if !res {
		return res, err
	}

	//write code  ---------------------------------------------
	outfile.WriteString(".begin_code\r\n")
	lineinfoNum := len(pf.lineinfos)
	var sourceCodeLine int
	for i := 0; i < len(pf.instructions); i++ {
		//add location label
		if inslabel, ok := locationinfos[i]; ok {
			outfile.WriteString(inslabel + ":\r\n")
		}

		instruction := pf.instructions[i]

		res, err, insstr, useExtended, extraArg := disassembler.ParseInstruction(pf, i, instruction, locationinfos)
		if !res {
			return res, err
		}
		if extraArg {
			outfile.WriteString(insstr)
		} else {
			outfile.WriteString("\t" + insstr)
		}

		if !useExtended {
			if i < lineinfoNum {
				sourceCodeLine = pf.lineinfos[i]
				outfile.WriteString(";L" + strconv.Itoa(sourceCodeLine) + ";")
			}
			outfile.WriteString("\r\n")
		}

	}
	outfile.WriteString(".end_code\r\n")

	//write subprotos   -----------------------------------------
	for i := 0; i < len(pf.usedSubroutines); i++ {
		outfile.WriteString("\r\n")
		sub := pf.usedSubroutines[i]
		subpf := disassembler.functions[sub]
		if subpf == nil {
			return false, "subproto not exist"
		}
		res, err := disassembler.WriteFuncAsm(outfile, subpf, false)
		if !res {
			return res, err
		}
	}
	outfile.WriteString("\r\n")
	return true, ""
}

func (disassembler *DisAssembler) ParseByteCode(inputFilepath string) (bool, string) {
	tempFileBytes, _ := ioutil.ReadFile(inputFilepath)
	disassembler.byteNum = len(tempFileBytes)
	disassembler.fileBytes = make([]byte, disassembler.byteNum)
	copy(disassembler.fileBytes, tempFileBytes)

	res, err := disassembler.ParseHead()
	if !res {
		return res, err
	}

	disassembler.nUpvals = int(disassembler.fileBytes[disassembler.readCur])
	disassembler.readCur++

	topfunc := new(ParsedFunction)
	res, err = disassembler.LoadFunc(topfunc, nil, 0)
	if !res {
		return res, err
	}

	if disassembler.readCur != disassembler.byteNum {
		return false, "file byte num wrong"
	}
	disassembler.topFunc = topfunc
	return true, ""
}
