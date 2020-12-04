operations.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex operations.tex
	make defs

make:
	make operations.pdf

open: operations.pdf
	open operations.pdf

defs: operations.tex
	defs operations.tex

o:
	make open
