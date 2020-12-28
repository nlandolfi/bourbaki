metric_space_functions.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex metric_space_functions.tex
	make defs

make:
	make metric_space_functions.pdf

open: metric_space_functions.pdf
	open metric_space_functions.pdf

o: open
	make open

defs: metric_space_functions.tex
	defs metric_space_functions.tex

clean:
	rm *.aux
	rm *.out
	rm *.log
