global _start
_start:
	mov eax, 1
	mov ebx, 22
	sub ebx, 1
	int 0x80

; return exit code 22