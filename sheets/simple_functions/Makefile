simple_functions.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex simple_functions.tex
	make defs

make:
	make simple_functions.pdf

open: simple_functions.pdf
	open simple_functions.pdf

defs: simple_functions.tex
	defs simple_functions.tex
