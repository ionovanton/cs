global _start

_start:
	call func_1
	call func_2
	mov eax, 1	; setting system exit code
	int 0x80	; perform system call

func_1:
	mov ebx, 42	; move 42 to ebx
	pop eax		; pop off of the stack return location of function call to eax
	jmp eax		; jump to function call location

func_2:			; performs like func_1
	mov ebx, 21
	ret
