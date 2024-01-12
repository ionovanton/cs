			global		main
			extern		printf
			section		.text
main:
			mov ecx, 4			; counter is 4
print:
			; print from 4 to 1
			mov rdi, format		; 1st arguments is format
			mov rsi, rcx		; 2nd argument is counter
			xor rax, rax		; printf vargs
			push rcx			; save counter before system call
			call printf
			pop rcx				; restore counter
			dec cl				; counter--
			jnz print
			ret

			; print from 1 to 4
			;mov rdi, format
			;xor rsi, rsi
			;mov r9, rcx
			;add rsi, r9
			;xor rax, rax
			;push rcx
			;call printf
			;pop rcx
			;inc ecx
			;cmp ecx, 4
			;jng print
			;ret
			
exit:
			ret

format:
			db "%i", 10, 0
