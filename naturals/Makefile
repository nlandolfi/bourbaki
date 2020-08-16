naturals.pdf: ../format.tex ../macros.tex ../trademark.pdf naturals.tex macros.tex
	pdflatex naturals.tex
	make defs

make:
	make naturals.pdf

open: naturals.pdf
	open naturals.pdf

defs: naturals.tex
	defs naturals.tex
