global _start

_start:
	call func
	mov eax, 1		; setting system exit code
	mov ebx, 0
	int 0x80		; perform system call

func:
	mov ebp, esp	; let ebp remember when top of stack was when the function
					; was called (esp - stack pointer, ebp - base pointer)
	sub esp, 4		; allocate 2 bytes on stack
	mov [esp], byte 'H'
	mov [esp + 1], byte 'i'
	mov [esp + 2], byte 10
	mov eax, 4
	mov ebx, 1
	mov ecx, esp
	mov edx, 3
	int 0x80
	mov esp, ebp	; restore esp like it was before func call
	ret
	

