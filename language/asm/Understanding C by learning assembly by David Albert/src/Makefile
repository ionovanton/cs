
include ../settings.env
export

all:
	nasm -felf64 $(FILE) -o program.o && ld program.o -o a.out && ./a.out &&

d:
	nasm -gdwarf -felf64 $(FILE) -o program.o && ld -g program.o -o a.out && gdb ./a.out

extern:
	nasm -felf64 $(FILE) -o program.o && gcc -O0 program.o -o a.out -static && ./a.out

extern_d:
	nasm -felf64 $(FILE) -o program.o && gcc -g -O0 program.o -o a.out -static && gdb ./a.out -tui

c:
	gcc -g -O0 $(FILE) -o a.out && ./a.out

c_mem:
	gcc -g -O0 $(FILE) -o a.out && valgrind ./a.out

c_d:
	gcc -g -O0 $(FILE) -o a.out && gdb ./a.out -tui

clean:
	rm -rf a.out program.o

test:
	echo $(FILE)
