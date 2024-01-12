CC=g++
SRC=main.cpp
NAME=a.out
STD=-std=c++17
ASAN=-fsanitize=address,undefined,leak
PERF=perf stat -d
FLAG=-g -O0
all:
	$(CC) $(FLAG) $(SRC) $(STD) -o $(NAME)

asan:
	$(CC) -ggdb -O0 $(ASAN) $(SRC) $(STD) -o $(NAME)

dbg:
	$(CC) -ggdb -O0 $(SRC) $(STD) -o $(NAME)

p: all
	$(PERF) ./a.out	