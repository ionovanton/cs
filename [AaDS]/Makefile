NAME	=compilers_tex
SRC1	=Cormen_solutions
SRC2	=Cormen_notes
SHELL	=/bin/bash
TMP		=$(addsuffix *.log, .tmp/) $(addsuffix *.aux, .tmp/)

all: solutions notes

s:
	@mkdir -p .tmp
	pdflatex -output-directory .tmp -interaction=batchmode $(SRC1).tex
	@mv .tmp/$(SRC1).pdf .

n:
	@mkdir -p .tmp
	@cls; pdflatex -output-directory .tmp -interaction=batchmode $(SRC2).tex
	@mv .tmp/$(SRC2).pdf .

clean:
	@rm .tmp

_s:
	@while true; do inotifywait $(SRC1).tex; sleep 0.01; make s; done

_n:
	@while true; do inotifywait $(SRC2).tex; sleep 0.01; make _n; done

