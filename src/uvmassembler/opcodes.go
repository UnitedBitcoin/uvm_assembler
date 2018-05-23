package uvmassembler

// OpCodes
const (
	/*----------------------------------------------------------------------
	  name		args	description
	  ------------------------------------------------------------------------*/
	OP_MOVE     = 0 /*	A B	R(A) := R(B)					*/
	OP_LOADK    = 1 /*	A Bx	R(A) := Kst(Bx)					*/
	OP_LOADKX   = 2 /*	A 	R(A) := Kst(extra arg)				*/
	OP_LOADBOOL = 3 /*	A B C	R(A) := (Bool)B; if (C) pc++			*/
	OP_LOADNIL  = 4 /*	A B	R(A), R(A+1), ..., R(A+B) := nil		*/
	OP_GETUPVAL = 5 /*	A B	R(A) := UpValue[B]				*/

	OP_GETTABUP = 6 /*	A B C	R(A) := UpValue[B][RK(C)]			*/
	OP_GETTABLE = 7 /*	A B C	R(A) := R(B)[RK(C)]				*/

	OP_SETTABUP = 8  /*	A B C	UpValue[A][RK(B)] := RK(C)			*/
	OP_SETUPVAL = 9  /*	A B	UpValue[B] := R(A)				*/
	OP_SETTABLE = 10 /*	A B C	R(A)[RK(B)] := RK(C)				*/

	OP_NEWTABLE = 11 /*	A B C	R(A) := {} (size = B,C)				*/

	OP_SELF = 12 /*	A B C	R(A+1) := R(B); R(A) := R(B)[RK(C)]		*/

	OP_ADD  = 13 /*	A B C	R(A) := RK(B) + RK(C)				*/
	OP_SUB  = 14 /*	A B C	R(A) := RK(B) - RK(C)				*/
	OP_MUL  = 15 /*	A B C	R(A) := RK(B) * RK(C)				*/
	OP_MOD  = 16 /*	A B C	R(A) := RK(B) % RK(C)				*/
	OP_POW  = 17 /*	A B C	R(A) := RK(B) ^ RK(C)				*/
	OP_DIV  = 18 /*	A B C	R(A) := RK(B) / RK(C)				*/
	OP_IDIV = 19 /*	A B C	R(A) := RK(B) // RK(C)				*/
	OP_BAND = 20 /*	A B C	R(A) := RK(B) & RK(C)				*/
	OP_BOR  = 21 /*	A B C	R(A) := RK(B) | RK(C)				*/
	OP_BXOR = 22 /*	A B C	R(A) := RK(B) ~ RK(C)				*/
	OP_SHL  = 23 /*	A B C	R(A) := RK(B) << RK(C)				*/
	OP_SHR  = 24 /*	A B C	R(A) := RK(B) >> RK(C)				*/
	OP_UNM  = 25 /*	A B	R(A) := -R(B)					*/
	OP_BNOT = 26 /*	A B	R(A) := ~R(B)					*/
	OP_NOT  = 27 /*	A B	R(A) := not R(B)				*/
	OP_LEN  = 28 /*	A B	R(A) := length of R(B)				*/

	OP_CONCAT = 29 /*	A B C	R(A) := R(B).. ... ..R(C)			*/

	OP_JMP = 30 /*	A sBx	pc+=sBx; if (A) close all upvalues >= R(A - 1)	*/
	OP_EQ  = 31 /*	A B C	if ((RK(B) == RK(C)) ~= A) then pc++		*/
	OP_LT  = 32 /*	A B C	if ((RK(B) <  RK(C)) ~= A) then pc++		*/
	OP_LE  = 33 /*	A B C	if ((RK(B) <= RK(C)) ~= A) then pc++		*/

	OP_TEST    = 34 /*	A C	if not (R(A) <=> C) then pc++			*/
	OP_TESTSET = 35 /*	A B C	if (R(B) <=> C) then R(A) := R(B) else pc++	*/

	OP_CALL     = 36 /*	A B C	R(A), ... ,R(A+C-2) := R(A)(R(A+1), ... ,R(A+B-1)) */
	OP_TAILCALL = 37 /*	A B C	return R(A)(R(A+1), ... ,R(A+B-1))		*/
	OP_RETURN   = 38 /*	A B	return R(A), ... ,R(A+B-2)	(see note)	*/

	OP_FORLOOP = 39 /*	A sBx	R(A)+=R(A+2);
		if R(A) <?= R(A+1) then { pc+=sBx; R(A+3)=R(A) }*/
	OP_FORPREP = 40 /*	A sBx	R(A)-=R(A+2); pc+=sBx				*/

	OP_TFORCALL = 41 /*	A C	R(A+3), ... ,R(A+2+C) := R(A)(R(A+1), R(A+2));	*/
	OP_TFORLOOP = 42 /*	A sBx	if R(A+1) ~= nil then { R(A)=R(A+1); pc += sBx }*/

	OP_SETLIST = 43 /*	A B C	R(A)[(C-1)*FPF+i] := R(A+i), 1 <= i <= B	*/

	OP_CLOSURE = 44 /*	A Bx	R(A) := closure(KPROTO[Bx])			*/

	OP_VARARG = 45 /*	A B	R(A), R(A+1), ..., R(A+B-2) = vararg		*/

	OP_EXTRAARG = 46 /*	Ax	extra (larger) argument for previous opcode	*/

	UOP_PUSH   = 47 /* A   top++, evalstack(top) = R(A)  */
	UOP_POP    = 48 /* A   R(A) := evalstack(top), top-- */
	UOP_GETTOP = 49 /* A   R(A) := evalstack(top) */
	UOP_CMP    = 50 /* A B C   R(A) = 1 if RK(B) > RK(C), 0 if RK(B) == RK(C), -1 if RK(B) < RK(C) */

	UOP_CMP_EQ = 51 /* A B C R(A) = 1 if RK(B) == RK(C), else 0 */
	UOP_CMP_NE = 52 /* A B C R(A) = 1 if RK(B) != RK(C), else 0 */
	UOP_CMP_GT = 53 /* A B C R(A) = 1 if RK(B) > RK(C), else 0 */
	UOP_CMP_LT = 54 /* A B C R(A) = 1 if RK(B) < RK(C), else 0 */

)

const NUM_OPCODES = int(UOP_CMP_LT) + 1

var UvmPOpnames = []string{
	"move",
	"loadk",
	"loadkx",
	"loadbool",
	"loadnil",
	"getupval",
	"gettabup",
	"gettable",
	"settabup",
	"setupval",
	"settable",
	"newtable",
	"self",
	"add",
	"sub",
	"mul",
	"mod",
	"pow",
	"div",
	"idiv",
	"band",
	"bor",
	"bxor",
	"shl",
	"shr",
	"unm",
	"bnot",
	"not",
	"len",
	"concat",
	"jmp",
	"eq",
	"lt",
	"le",
	"test",
	"testset",
	"call",
	"tailcall",
	"return",
	"forloop",
	"forprep",
	"tforcall",
	"tforloop",
	"setlist",
	"closure",
	"vararg",

	"extraarg",

	"push",
	"pop",
	"gettop",
	"cmp",

	"cmp_eq",
	"cmp_ne",
	"cmp_gt",
	"cmp_lt",


	""}

// count of parameters for each instruction
var Opcounts = []int{
	2, // MOVE
	2, // LOADK
	2, // LOADKX
	3, // LOADBOOL
	2, // LOADNIL
	2, // GETUPVAL
	3, // GETTABUP
	3, // GETTABLE
	3, // SETTABUP
	2, // SETUPVAL
	3, // SETTABLE
	3, // NEWTABLE
	3, // SELF
	3, // ADD
	3, // SUB
	3, // MUL
	3, // DIV
	3, // BAND
	3, // BOR
	3, // BXOR
	3, // SHL
	3, // SHR
	3, // MOD
	3, // IDIV
	3, // POW
	2, // UNM
	2, // BNOT
	2, // NOT
	2, // LEN
	3, // CONCAT
	2, // JMP
	3, // EQ
	3, // LT
	3, // LE
	2, // TEST
	3, // TESTSET
	3, // CALL
	3, // TAILCALL
	2, // RETURN
	2, // FORLOOP
	2, // FORPREP
	2, // TFORCALL
	2, // TFORLOOP
	3, // SETLIST
	2, // CLOSURE
	2, // VARARG
	1, // EXTRAARG
	1, // PUSH
	1, // POP
	1, // GETTOP
	3, // CMP
	3, // CMP_EQ
	3, // CMP_NE
	3, // CMP_GT
	3, // CMP_LT
}

const (
	LIMIT_STACKIDX    = 1
	LIMIT_UPVALUE     = 2
	LIMIT_LOCATION    = 4
	LIMIT_CONSTANT    = 8
	LIMIT_EMBED       = 0x10
	LIMIT_PROTO       = 0x20
	LIMIT_CONST_STACK = LIMIT_CONSTANT | LIMIT_STACKIDX
)

// OpPos
const (
	OPP_A     = 0
	OPP_B     = 1
	OPP_C     = 2
	OPP_Ax    = 3
	OPP_Bx    = 4
	OPP_sBx   = 5
	OPP_ARG   = 6
	OPP_C_ARG = 7
)

type OpPos int

type OpInfo struct {
	pos   OpPos
	limit int
}

var Opinfos = [][]OpInfo{ // Maximum of 3 operands
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_STACKIDX}},                                // MOVE
	{{OPP_A, LIMIT_STACKIDX}, {OPP_Bx, LIMIT_CONSTANT}},                               // LOADK
	{{OPP_A, LIMIT_STACKIDX}, {OPP_ARG, LIMIT_CONSTANT}},                              // LOADKX
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_EMBED}, {OPP_C, LIMIT_EMBED}},             // LOADBOOL
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_EMBED}},                                   // LOADNIL
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_UPVALUE}},                                 // GETUPVAL
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_UPVALUE}, {OPP_C, LIMIT_CONST_STACK}},     // GETTABUP
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_STACKIDX}, {OPP_C, LIMIT_CONST_STACK}},    // GETTABLE
	{{OPP_A, LIMIT_UPVALUE}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}},  // SETTABUP
	{{OPP_B, LIMIT_UPVALUE}, {OPP_A, LIMIT_STACKIDX}},                                 /// SETUPVAL
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // SETTABLE
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_EMBED}, {OPP_C, LIMIT_EMBED}},             // NEWTABLE
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_STACKIDX}, {OPP_C, LIMIT_CONST_STACK}},    // SELF
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // ADD
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // SUB
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // MUL
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // DIV
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // BAND
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // BOR
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // BXOR
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // SHL
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // SHR
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // MOD
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // IDIV
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // POW
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_STACKIDX}},                                // UNM
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_STACKIDX}},                                // BNOT
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_STACKIDX}},                                // NOT
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_STACKIDX}},                                // LEN
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_STACKIDX}, {OPP_C, LIMIT_STACKIDX}},       // CONCAT
	{{OPP_A, LIMIT_EMBED}, {OPP_sBx, LIMIT_LOCATION}},                                 // JMP
	{{OPP_A, LIMIT_EMBED}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}},    // EQ
	{{OPP_A, LIMIT_EMBED}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}},    // LT
	{{OPP_A, LIMIT_EMBED}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}},    // LE
	{{OPP_A, LIMIT_STACKIDX}, {OPP_C, LIMIT_EMBED}},                                   // TEST
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_STACKIDX}, {OPP_C, LIMIT_EMBED}},          // TESTSET
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_EMBED}, {OPP_C, LIMIT_EMBED}},             // CALL
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_EMBED}, {OPP_C, LIMIT_EMBED}},             // TAILCALL
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_EMBED}},                                   // RETURN
	{{OPP_A, LIMIT_STACKIDX}, {OPP_sBx, LIMIT_LOCATION}},                              // FORLOOP
	{{OPP_A, LIMIT_STACKIDX}, {OPP_sBx, LIMIT_LOCATION}},                              // FORPREP
	{{OPP_A, LIMIT_STACKIDX}, {OPP_C, LIMIT_EMBED}},                                   // TFORCALL
	{{OPP_A, LIMIT_STACKIDX}, {OPP_sBx, LIMIT_LOCATION}},                              // TFORLOOP
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_EMBED}, {OPP_C_ARG, LIMIT_EMBED}},         // SETLIST
	{{OPP_A, LIMIT_STACKIDX}, {OPP_Bx, LIMIT_PROTO}},                                  // CLOSURE
	{{OPP_A, LIMIT_STACKIDX}, {OPP_Bx, LIMIT_EMBED}},                                  // VARARG

	{{OPP_Ax, LIMIT_EMBED}}, // EXTRAARG

	{{OPP_A, LIMIT_STACKIDX}}, // PUSH
	{{OPP_A, LIMIT_STACKIDX}}, // POP
	{{OPP_A, LIMIT_STACKIDX}}, // GETTOP
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // CMP

	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // CMP_EQ
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // CMP_NE
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // CMP_GT
	{{OPP_A, LIMIT_STACKIDX}, {OPP_B, LIMIT_CONST_STACK}, {OPP_C, LIMIT_CONST_STACK}}, // CMP_LT
}
