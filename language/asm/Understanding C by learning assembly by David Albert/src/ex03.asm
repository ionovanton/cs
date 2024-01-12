global _start

section .text
_start:
	mov ecx, 100
	mov ebx, 42
	mov eax, 1
	cmp ecx, 20		; compare ecx to value
	jl a			; jump if less than (based on previous comparison)
	je b
	jg c
a:
	mov ebx, 1
	int 0x80
b:
	mov ebx, 2
	int 0x80
c:
	mov ebx, 3
	int 0x80
