## Progress
### 0 / 1

## 2.3 How `make` Processes a Makefile

By default, make starts with the first target (not targets whose names start with ‘.’ unless they also contain one or more ‘/’). This is called the default goal.


```make
edit : main.o kbd.o command.o display.o \
       insert.o search.o files.o utils.o
        cc -o edit main.o kbd.o command.o display.o \
                   insert.o search.o files.o utils.o

main.o : main.c defs.h
        cc -c main.c
kbd.o : kbd.c defs.h command.h
        cc -c kbd.c
command.o : command.c defs.h command.h
        cc -c command.c
display.o : display.c defs.h buffer.h
        cc -c display.c
insert.o : insert.c defs.h buffer.h
        cc -c insert.c
search.o : search.c defs.h buffer.h
        cc -c search.c
files.o : files.c defs.h buffer.h command.h
        cc -c files.c
utils.o : utils.c defs.h
        cc -c utils.c
clean :
        rm edit main.o kbd.o command.o display.o \
           insert.o search.o files.o utils.o
```

- `make` reads the makefile in the current directory and begins by processing the first rule — `edit`.
 - Before make can fully process this rule, it must process the rules for the files that edit depends on, which in this case are the object files
 - Recompilation must be done if the source file, or any of the header files named as prerequisites, is more recent than the object file, or if the object file does not exist.
 - After recompiling whichever object files need it, make decides whether to relink edit. This must be done if the file edit does not exist, or if any of the object files are newer than it. If an object file was just recompiled, it is now newer than edit, so edit is relinked.
 - Thus, if we change the file `insert.c` and run `make`, make will compile that file to update `insert.o`, and then link `edit`. If we change the file `command.h` and run `make`, make will recompile the object files `kbd.o`, `command.o` and `files.o` and then link the file `edit`.

## 3.7 How `make` Reads a Makefile

- Phase 1: reads all the makefiles, included makefiles, etc. and internalizes all the variables and their values and implicit and explicit rules, and builds a dependency graph of all the targets and their prerequisites
- Phase 2: `make` uses this internalized data to determine which targets need to be updated and run the recipes necessary to update them

We say that expansion is *immediate* if it happens during the first phase: make will expand that part of the construct as the makefile is parsed. We say that expansion is *deferred* if it is not immediate.

Order:

```make
immediate = deferred
immediate ?= deferred
immediate := immediate
immediate ::= immediate
immediate :::= immediate-with-escape
immediate += deferred or immediate
immediate != immediate

define immediate
  deferred
endef

define immediate =
  deferred
endef

define immediate ?=
  deferred
endef

define immediate :=
  immediate
endef

define immediate ::=
  immediate
endef

define immediate :::=
  immediate-with-escape
endef

define immediate +=
  deferred or immediate
endef

define immediate !=
  immediate
endef
```



