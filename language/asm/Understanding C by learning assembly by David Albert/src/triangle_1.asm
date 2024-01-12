			global _start

			section .bss
			max_lines equ 9

			section .data
			star db '*', 0
			nl db 10

			section .text
_start:		
			mov rsi, star	; move star to rsi
			mov rax, 1		; syscall for write
			mov rdx, 1		; length of a single star
			mov r8, 1		; i = 1
			mov r9, 0		; j = 0

line:		
			syscall			; write from rsi
			inc r9			; i++
			cmp r9, r8		; j != i
			jne line

line_done:	
			mov rsi, nl		; replace star with newline
			syscall
			mov rsi, star 	; replace newline with star
			inc r8			; i++
			xor r9, r9		; j = 0
			cmp r8, max_lines
			jng line		; i < max_lines

done:
			mov rax, 60		; syscall for exit
			mov rdi, 0		; exit code
			syscall
