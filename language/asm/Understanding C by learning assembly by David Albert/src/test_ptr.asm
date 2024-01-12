			global		main
			extern		printf
			section		.text
main:
			mov ecx, 0
			xor r9, r9
job:
			; print only even numbers from 1 to 10
			cmp rcx, 10				; while (counter != 10)
			jg exit
			mov rax, rcx			; prepare statement check
			and rax, 1
			test rax, rax			; if (counter % 2 == 0)
			je print
			inc rcx					; counter++
			jmp job


print:		; prepare printf
			mov rdi, format			; 1st argument is format
			mov rsi, rcx			; 2nd argument is counter
			xor rax, rax			; vargs			
			push rcx				; save counter
			call printf
			pop rcx					; restore counter
			inc rcx					; counter++
			jmp job
			
exit:
			ret

format:
			db "%i", 10, 0
