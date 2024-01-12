			global		main
			extern		printf
			section		.text
main:
			mov ecx, 0
			xor r9, r9
print:
			; print from 4 to 1
			;mov rdi, format
			;xor rsi, rsi
			;mov r9, rcx
			;add rsi, r9
			;xor rax, rax
			;push rcx
			;call printf
			;pop rcx
			;dec ecx
			;jnz print
			;ret

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

			; print only even numbers from 1 to 10
			cmp rcx, 10				; while (counter != 10)
			jg exit
			mov rax, rcx			; prepare statement check
			and rax, 1
			test rax, rax			; if (counter % 2 == 0)
			je .L2
			inc rcx					; counter++
			jmp print


.L2:		; prepare printf
			mov rdi, format			; 1st argument is format
			mov rsi, rcx			; 2nd argument is counter
			xor rax, rax			; vargs			
			push rcx				; save counter
			call printf
			pop rcx					; restore counter
			inc rcx					; counter++
			jmp print
			
exit:
			ret

format:
			db "%i", 10, 0
