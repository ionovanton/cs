/* test.c */

#include <stdio.h>
#include <stdlib.h>

int main()
{
	char stack[5];
	char *heap = malloc(sizeof(char) * (5 + 1));

	for (int i = 0; i < 5; i++)
	{
		stack[i] = 97 + i;
		heap[i] = 97 + i;
	}
	stack[5] = 0;
	heap[5] = 0;
	printf("stack:	%s\n", stack);
	printf("heap:	%s\n", heap);
	free(heap);
}
