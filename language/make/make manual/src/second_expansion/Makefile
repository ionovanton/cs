main_OBJS := main.o try.o test.o
lib_OBJS := lib.o api.o

all: main lib

clean:
	rm main lib $(main_OBJS) $(lib_OBJS)

.SECONDEXPANSION:
main lib: $$($$@_OBJS)
