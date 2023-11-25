NAME	=compilers_tex
SRC1	=Compilers_solutions
SRC2	=Compilers_notes
SHELL	=/bin/bash
TMP		=$(addsuffix *.log, .tmp/) $(addsuffix *.aux, .tmp/)

all: solutions notes

solutions:
	@cls; pdflatex -output-directory .tmp -interaction=batchmode $(SRC1).tex
	@mv .tmp/$(SRC1).pdf .

notes:
	@cls; pdflatex -output-directory .tmp -interaction=batchmode $(SRC2).tex
	@mv .tmp/$(SRC2).pdf .

clean:
	@rm .tmp

watch_solutions:
	@while true; do inotifywait $(SRC1).tex; sleep 0.01; make all; done

watch_notes:
	@while true; do inotifywait $(SRC2).tex; sleep 0.01; make all; done

