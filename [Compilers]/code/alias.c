#include <stdlib.h>
#include <stdio.h>

void p(int x, int y) {
	x = 9890;
	printf("p: x=%d	 y=%d\n", x, y);
}

void t(int *x, int *y) {
	x[1] = 377;
	printf("t: x[1]=%d	 y[1]=%d\n", x[1], y[1]);
}

void q(int x[], int y[]) {
	x[1] = 42;
	printf("q: x[1]=%d	 y[1]=%d\n", x[1], y[1]);
}

void main() {
	int a[] = {1, 2, 3};
	int *b = malloc(sizeof(int) * 3);
	int c = 1;
	b[0] = 1; b[1] = 2; b[2] = 3;

	q(a, a); t(b, b); p(c, c);
}

