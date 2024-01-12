## Progress
### 1 / 1

# **[Understanding C by learning assembly by David Albert](https://www.recurse.com/blog/7-understanding-c-by-learning-assembly#footnote_p7f3)**
## **Learning assembly with GDB**
Consider next code:
``` c
/* simple.c */
int main()
{
	int a = 5;
	int b = a + 6;
	return 0;
}
```
Compile it with debugging symbols and no optimization:
```
$ CFLAGS="-g -O0" make simple
cc -g -O0 simple.c -o simple
$ gdb simple
```

Now let's use the `disassemble` command to show the assembly instructions for the current function. You can also pass a function name to `disassemble` to specify a different function to examine.

### *intel flavor*
``` assembly
Dump of assembler code for function main:
   0x0000000000001119 <+0>:     push   rbp
   0x000000000000111a <+1>:     mov    rbp,rsp
   0x000000000000111d <+4>:     mov    DWORD PTR [rbp-0x8],0x5
   0x0000000000001124 <+11>:    mov    eax,DWORD PTR [rbp-0x8]
   0x0000000000001127 <+14>:    add    eax,0x6
   0x000000000000112a <+17>:    mov    DWORD PTR [rbp-0x4],eax
   0x000000000000112d <+20>:    mov    eax,0x0
   0x0000000000001132 <+25>:    pop    rbp
   0x0000000000001133 <+26>:    ret    
End of assembler dump.
```

### *att flavor*
``` assembly
Dump of assembler code for function main:
   0x0000000000001119 <+0>:     push   %rbp
   0x000000000000111a <+1>:     mov    %rsp,%rbp
   0x000000000000111d <+4>:     movl   $0x5,-0x8(%rbp)
   0x0000000000001124 <+11>:    mov    -0x8(%rbp),%eax
   0x0000000000001127 <+14>:    add    $0x6,%eax
   0x000000000000112a <+17>:    mov    %eax,-0x4(%rbp)
   0x000000000000112d <+20>:    mov    $0x0,%eax
   0x0000000000001132 <+25>:    pop    %rbp
   0x0000000000001133 <+26>:    ret
End of assembler dump.
```

## **Registers**
Registers are data storage locations directly on the CPU. With some exceptions, the size, or width, of a CPU’s registers define its architecture. So if you have a 64-bit CPU, your registers will be 64 bits wide. The same is true of 32-bit CPUs (32-bit registers), 16-bit CPUs, and so on. Registers are very fast to access and are often the operands for arithmetic and logic operations.

The x86 family has a number of general and special purpose registers. General purpose registers can be used for any operation and their value has no particular meaning to the CPU. On the other hand, the CPU relies on special purpose registers for its own operation and the values stored in them have a specific meaning depending on the register. In our example above, `%eax` and `%ecx` are general purpose registers, while `%rbp` and `%rsp` are special purpose registers. `%rbp` is the base pointer, which points to the base of the current stack frame, and `%rsp` is the stack pointer, which points to the top of the current stack frame. `%rbp` always has a higher value than `%rsp` because the stack starts at a high memory address and grows downwards. 


The first two unstructions are called the function prologue. First we push the old base pointer onto the stack to save it for later. Then we copy the value of the stack pointer to the base pointer. After this, `%rbp` points to the base of `main`'s stack frame and `%rsp` is pointing to the top of the stack.

``` assembly
; intel
0x000000000000112d <+20>:    mov    eax,0x0
; att
0x000000000000112d <+20>:    mov    $0x0,%eax
```
This instruction copies 0 into `%eax`. The x86 calling convention dictates that a function’s return value is stored in `%eax`, so the above instruction sets us up to return 0 at the end of our function.

``` assembly
; intel
0x000000000000112a <+17>:    mov    DWORD PTR [rbp-0x4],eax
; att
0x000000000000112a <+17>:    mov    %eax,-0x4(%rbp)
```
The parentheses let us know that this is a memory address. Here, `%rbp` is called the base register, and -0x4 is the displacement. This is equivalent to `%rbp + -0x4`. Because the stack grows downwards, subtracting 4 from the base of the current stack frame moves us into the current frame itself, where local variables are stored. It seems that clang allocates a hidden local variable for an implicit return value from main.

``` assembly
; intel
0x000000000000111d <+4>:     mov    DWORD PTR [rbp-0x8],0x5
; att
0x000000000000111d <+4>:     movl   $0x5,-0x8(%rbp)
```
Notice that the mnemonic has the suffix `l`. This signifies that the operands will be long (32 bits for integers). Other valid suffixes are `b`yte, `s`hort, `w`ord, `q`uad, and `t`en. If you see an instruction that does not have a suffix, the size of the operands are inferred from the size of the source or destination register. For instance, in the previous line, `%eax` is 32 bits wide, so the `mov` instruction is inferred to be `movl`.

The first line of assembly is the first line of C in `main` and stores the number 5 in the next available local variable slot `(%rbp - 0x8)`, 4 bytes down from our last local variable. That’s the location of `a`. We can use GDB to verify this:
``` gdb
(gdb) x &a
0x7fffffffe268: 0x00000005
(gdb) x $rbp-8
0x7fffffffe268: 0x00000005
```

Note that the memory addresses are the same. You’ll notice that GDB sets up variables for our registers, but like all variables in GDB, we prefix it with a `$` rather than the `%` used in AT&T assembly.

``` assembly
; intel
0x0000000000001124 <+11>:    mov    eax,DWORD PTR [rbp-0x8]
0x0000000000001127 <+14>:    add    eax,0x6
0x000000000000112a <+17>:    mov    DWORD PTR [rbp-0x4],eax
; att
0x0000000000001124 <+11>:    mov    -0x8(%rbp),%eax	; move a to %eax
0x0000000000001127 <+14>:    add    $0x6,%eax		; add 6 to %eax
0x000000000000112a <+17>:    mov    %eax,-0x4(%rbp)	; move result of eax to %rbp - 0x4
```

We then move `a` into `%eax`, one of our general purpose registers, add 6 to it and store the result in `%rbp - 0x4`. You've maybe figured it out that `%rbp - 4` is `b`, which we can verify in GDB:
``` gdb
(gdb) x &b
0x7fffffffe26c: 0x0000000b
(gdb) x $rbp-4
0x7fffffffe26c: 0x0000000b
```

The rest of `main` is just clean up, called the function epilogue:
``` assembly
; intel
0x0000000000001132 <+25>:    pop    rbp
0x0000000000001133 <+26>:    ret
; att
0x0000000000001132 <+25>:    pop    %rbp
0x0000000000001133 <+26>:    ret 
```

We `pop` the old base pointer off the stack and store it back in `%rbp` and then `ret` jumps back to out return address, wheich is also stored in the stack frame.

Now let's figure out how static variables work.

## **Understanding static local variables**

``` c
/* static.c */

#include <stdio.h>

int natural_generator()
{
	int a = 1;
	static int b = -1;
	b += 1;
	return a + b;
}

int main()
{
	printf("%d\n", natural_generator());
	printf("%d\n", natural_generator());
	printf("%d\n", natural_generator());
}
```

Here is disassembly of `natural_generator`:

``` assembly
; intel
Dump of assembler code for function natural_generator:
   0x0000555555555139 <+0>:     push   rbp
   0x000055555555513a <+1>:     mov    rbp,rsp
   0x000055555555513d <+4>:     mov    DWORD PTR [rbp-0x4],0x1
   0x0000555555555144 <+11>:    mov    eax,DWORD PTR [rip+0x2ece]        # 0x555555558018 <b.0>
   0x000055555555514a <+17>:    add    eax,0x1
   0x000055555555514d <+20>:    mov    DWORD PTR [rip+0x2ec5],eax        # 0x555555558018 <b.0>
   0x0000555555555153 <+26>:    mov    edx,DWORD PTR [rip+0x2ebf]        # 0x555555558018 <b.0>
   0x0000555555555159 <+32>:    mov    eax,DWORD PTR [rbp-0x4]
   0x000055555555515c <+35>:    add    eax,edx
   0x000055555555515e <+37>:    pop    rbp
   0x000055555555515f <+38>:    ret    
End of assembler dump.
; att
Dump of assembler code for function natural_generator:
   0x0000555555555139 <+0>:     push   %rbp
   0x000055555555513a <+1>:     mov    %rsp,%rbp
   0x000055555555513d <+4>:     movl   $0x1,-0x4(%rbp)
   0x0000555555555144 <+11>:    mov    0x2ece(%rip),%eax        # 0x555555558018 <b.0>
   0x000055555555514a <+17>:    add    $0x1,%eax
   0x000055555555514d <+20>:    mov    %eax,0x2ec5(%rip)        # 0x555555558018 <b.0>
   0x0000555555555153 <+26>:    mov    0x2ebf(%rip),%edx        # 0x555555558018 <b.0>
   0x0000555555555159 <+32>:    mov    -0x4(%rbp),%eax
   0x000055555555515c <+35>:    add    %edx,%eax
   0x000055555555515e <+37>:    pop    %rbp
   0x000055555555515f <+38>:    ret    
End of assembler dump.
```

We can also find out on what instruction we're currently on by examining register `%rip`, or alternatively we can use the architecture independent `$pc`:
``` gdb
(gdb) x/i $pc
0x55555555513d <natural_generator+4>:	movl   $0x1,-0x4(%rbp)
```

The instruction pointer always contains the address of the next instruction to be run, which means the third instruction hasn’t been run yet, but is about to be.

We can see that `a` is stored at `$rbp - 4`:
``` assembly
0x000055555555513d <+4>:     movl   $0x1,-0x4(%rbp)
```

But then we'd expect to find the line `static int b = -1`, but it looks substantially different than anything we’ve seen before. For one thing, there's no reference to the stack frame where we'd normally expect to find local variables. There's not even a `-0x1`. Instead, we have an instruction that loads `0x555555558018`, located somewhere after the instruction pointer, into `%eax`. GDB gives us a helpful comment with *the result of the memory operand calculation* and a hint telling us that `natural_generator.b` is stored at this address.

``` gdb
(gdb) p $rax
$11 = -1
```
It looks like we've found `b`. We can double check this by using the `x` command:
``` gdb
(gdb) x/d 0x555555558018
0x555555558018 <b.0>:	-1
(gdb) x/d &b
0x555555558018 <b.0>:	-1
```

So not only is `b` stored at a low memory address outside of the stack, it’s also initialized to -1 before `natural_generator` is even called. In fact, even if you disassembled the entire program, you wouldn’t find any code that sets `b` to -1. This is because the value for `b` is hardcoded in a different section of the `sample` executable, and it’s loaded into memory along with all the machine code by the operating system’s loader when the process is launched.

Now we understand how `b` is initialized, let’s see what happens when we run `natural_generator` again:
``` gdb
(gdb) continue
(gdb) x &b
0x555555558018 <b.0>:	0
```
Because `b` is not stored on the stack with other local variables, it’s still zero when `natural_generator` is called again. No matter how many times our generator is called, `b` will always retain its previous value. This is because it’s stored outside the stack and initialized when the loader moves the program into memory, rather than by any of our machine code.

## **Conclusion**
We began by going over how to read assembly and how to disassemble a program with GDB. Afterwards, we covered how static local variables work, which we could not have done without disassembling our executable.

We spent a lot of time alternating between reading the assembly instructions and verifying our hypotheses in GDB. It may seem repetitive, but there’s a very important reason for doing things this way: **the best way to learn something abstract is to make it more concrete, and one of the best way to make something more concrete is to use tools that let you peel back layers of abstraction. The best way to to learn these tools is to force yourself to use them until they’re second nature**.

# **Other**
## **Calling Conventions**
From left to right, pass as many parameters as will fit in registers. The order in which registers are allocated, are:
- For integers and pointers:
	- `rdi`, `rsi`, `rdx`, `rcx`, `r8`, `r9`
- For floating-point:
	- `xmm0`, `xmm1`, `xmm2`, `xmm3`, `xmm4`, `xmm5`, `xmm6`, `xmm7`

Additional parameters are pushed on stack, right to left, and are to be removed by the caller after the call.

After parameters are pushed, the call instruction is made, so when the called function gets control, the return address at `[rsp]`, the first memory parameter is at `[rsp + 8]`, etc.

The `rsp` must be aligned to a 16-byte boundary before making a call.

The only registers that the called function is required to preserve are: `rbp`, `rbx`, `r12`, `r13`, `r14`, `r15`. All others are free to bechanged by the called function.

Integers are returned in `rax` or `rdx:rax`, and floating-point values are returned in `xmm0` or `xmm1:xmm0`.