global _start

_start:
	push 21				; push argument 21 to stack
	call times2			; push return address to stack
	mov ebx, eax
	mov eax, 1
	int 0x80

times2:
	push ebp
	mov ebp, esp		; write to ebp current top of the stack
	mov eax, [ebp + 8]	; copy to eax value (21) that were pushed earlier to the stack
	add eax, eax		; add eax to itself
	mov esp, ebp
	pop ebp
	ret

