arithmetic.pdf: ../format.tex ../macros.tex ../trademark.pdf arithmetic.tex macros.tex
	pdflatex arithmetic.tex
	make defs

make:
	make arithmetic.pdf

open: arithmetic.pdf
	open arithmetic.pdf

defs: arithmetic.tex
	defs arithmetic.tex
