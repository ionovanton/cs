global _start

section .data
	addr db "yellow"
	len equ $-addr

section .text
_start:
	mov [addr], byte 'H'
	mov [addr + 5], byte '!'
	mov eax, 4
	mov ebx, 1
	mov ecx, addr
	mov edx, len
	int 0x80
	mov eax, 1
	mov ebx, 0
	int 0x80