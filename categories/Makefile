categories.pdf: ../format.tex ../macros.tex ../trademark.pdf categories.tex macros.tex
	pdflatex categories.tex
	make defs

make:
	make categories.pdf

open: categories.pdf
	open categories.pdf

defs: categories.tex
	defs categories.tex
