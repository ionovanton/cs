			global _start

			section .text
_start:		mov rax, 1		; syscall for write
							; values are returned from functions in this register
							; rax is the 64-bit, "long" size register
			mov rdi, 1		; file handle 1 to stdout
							; scratch register / temporary register — a register used to hold an intermediate value during a calculation
							; (usually, such values are not named in the program source and have a limited lifetime)
			mov rsi, msg	; scratch register
			mov rdx, len	; scratch register
			syscall			; invoke OS to do write
			mov rax, 60		; syscall for exit
			mov rdi, 0		; exit code
			syscall
			
			section .data
			msg db "Hello world!", 10, 0
			len equ $ - msg
