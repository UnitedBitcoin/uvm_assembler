Pseudo-assembly Syntax of uvm
================================

* One instruction per line, unless it is a cross-line of a string
* String literals need to be escaped
* In addition to the contents of the string, the current line after the semicolon is (optional 'L'+ line number + ';') + current line comment
* There may be some blanks at the beginning of a one-line instruction. There may be some blanks between the instruction name and instruction parameters.
* `.local varname`  syntax to declare variable name
* syntax to declare const value:

```
  .begin_const
     "asd"
     "Hello"
     " world!"
  .end_const
```

* syntax to declare upvalues:

```
  .begin_upvalue
     1 0
  .end_upvalue
```

* syntax to declare code section

```
.begin_code
   loadk %0 const "asd"    ;L12;  %0 means slot 0
   closure %1 subroutine_2
   move %2 %1
   loadk %3 const "Hello"
   loadk %4 const " world!"
   call %2 3 1
   return %0 1
.end_code
```

* `.upvalues upvalues_count` syntax to declare upvalues count of that proto
* `.func func_name maxstacksize params_count use_vararg` declare start of a proto, with proto's func name, arguments, etc.
* loadk, setglobal, return, move and other instructions correspond to the corresponding uvm bytecode instructions, followed by the corresponding parameters
* `.end_func` syntax to finish declaration of a proto
* A `do end` code section generates only one `return` instruction


# Instruction Set

```

  MOVE Copy a value between registers
  LOADK Load a constant into a register
  LOADBOOL Load a boolean into a register
  LOADNIL Load nil values into a range of registers
  GETUPVAL Read an upvalue into a register
  GETTABLE Read a table element into a register
  SETUPVAL Write a register value into an upvalue
  SETTABLE Write a register value into a table element
  NEWTABLE Create a new table
  SELF Prepare an object method for calling
  ADD Addition operator
  SUB Subtraction operator
  MUL Multiplication operator
  DIV Division operator
  MOD Modulus (remainder) operator
  POW Exponentiation operator
  UNM Unary minus operator
  NOT Logical NOT operator
  LEN Length operator
  CONCAT Concatenate a range of registers
  JMP Unconditional jump
  EQ Equality test
  LT Less than test
  LE Less than or equal to test
  TEST Boolean test, with conditional jump
  TESTSET Boolean test, with conditional jump and assignment
  CALL Call a closure
  TAILCALL Perform a tail call
  RETURN Return from function call
  FORLOOP Iterate a numeric for loop
  FORPREP Initialization for a numeric for loop
  TFORLOOP Iterate a generic for loop
  SETLIST Set a range of array elements for a table
  CLOSURE Create a closure of a function prototype
  VARARG Assign vararg function arguments to registers
  PUSH push register value to evaluation stack
  POP pop top value in evaluation stack to register
  GETTOP get top value in evaluation stack to register
  CMP R(A) = 1 if RK(B) > RK(C), 0 if RK(B) == RK(C), -1 if RK(B) < RK(C) */

```
