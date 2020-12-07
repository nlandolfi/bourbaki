associative_operations.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex associative_operations.tex
	make defs

make:
	make associative_operations.pdf

open: associative_operations.pdf
	open associative_operations.pdf

defs: associative_operations.tex
	defs associative_operations.tex

o:
	make open
