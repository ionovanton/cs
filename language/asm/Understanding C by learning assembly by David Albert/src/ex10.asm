section .data
	msg db "Testing %i...", 0x0a, 0x00

section .text
global main
extern printf
main:
	push ebp		; prologue
	mov ebp, esp	; prologue
	push 123
	push msg
	call printf
	mov eax, 0
	mov esp, ebp	; epilogue
	pop ebp			; epilogue
	ret				; epilogue
