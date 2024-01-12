			global _start

			section		.bss				; writable data (non-init data)
			max_lines	equ		8
			data_size	equ		44
			output		resb	data_size

			section		.text				; actual code
											; section .data is for init data
_start:		
			mov rdx, output		; length of the output will be 44
			mov r8, 0			; i = 1
			mov r9, 0			; j = 0

line:
			mov byte [rdx], '*'
			inc rdx				; advance pointer
			inc r9				; j++
			cmp r9, r8			; j < i
			jng line

line_done:
			mov byte [rdx], 10	; write newline character
			inc rdx				; move rdx pointer to next char
			inc r8				; i++
			xor	r9, r9			; j = 0
			cmp r8, max_lines	; i < max_lines
			jng line

done:
			mov rax, 1
			mov rdi, 1
			mov rsi, output
			mov rdx, data_size
			syscall
			mov rax, 60
			xor rdi, rdi
			syscall


