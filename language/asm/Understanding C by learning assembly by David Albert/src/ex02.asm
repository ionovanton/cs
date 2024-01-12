global _start

section .data
	msg db "Hello World!", 0x0a
	len equ $ - msg

section .text
_start:
	mov eax, 4		; 4 to denote that is a system write call
	mov ebx, 1		; 1 is stdout file descriptor
	mov ecx, msg	; bytes to write
	mov edx, len	; number of bytes to write
	int 0x80
	mov eax, 1
	mov ebx, 0
	int 0x80
