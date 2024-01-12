			global		main
			extern		printf
			section		.text
main:
			push rbx		; we have to save it since we use it
			mov ecx, 4		; counter
			xor rax, rax	; current number
			xor rbx, rbx	; next number
			inc rbx			; next number = 1

print:
			; We need to call printf, but we are using rax, rbx, and rcx. Printf
        	; may destroy rax and rcx so we will save these before the call and
        	; restore them afterwards.

			push rax		; caller-save register (1)
			push rcx		; caller-save register
			mov rdi, format	; set 1st parameter (format)
			mov rsi, rax	; set 2nd parameter (current number)
			xor rax, rax	; because printf use vargs (3)

			; stack is already aligned because we pushed three 8 bytes registers (2)
			call printf
			pop rcx			; restore
			pop rax			; restore

			mov rdx, rax	; save current number
			mov rax, rbx	; current = current
			add rbx, rdx	; get the new next number
			dec ecx			; counter--
			jnz print		; jump if not zero

			pop rbx
			ret

format:
			db "%20ld", 10, 0

			; compiled with:
			; nasm -felf64 $1 -o program.o; gcc program.o -o a.out -static && ./a.out
			; 
			; (1)
			; Caller-saved registers are used to hold temporary quantities that need not be preserved across calls.
			; For that reason, it is the caller's responsibility to push these registers onto the stack 
			; or copy them somewhere else if it wants to restore this value after a procedure call.
			;
			; Callee-saved registers are used to hold long-lived values that should be preserved across calls.
			; When the caller makes a procedure call, it can expect that those registers 
			; will hold the same value after the callee returns, making it the responsibility of the callee 
			; to save them and restore them before returning to the caller. Or to not touch them.
			;
			; (2)
			; Word addresses end in a number divisible by 4: 0, 4, 8, c (hex)
			;
			; (3)
			; To call printf from assembly language, you just pass the format string in rdi as usual 
			; for the first argument, pass any format specifier arguments in the next argument register rsi, 
			; then rdx, etc.  There is one surprise: you need to zero out the al register 
			; (the low 8 bits of eax/rax), or else printf will assume you're passing it vector registers and crash.
