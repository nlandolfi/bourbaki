iterated_integrals.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex iterated_integrals.tex
	make defs

make:
	make iterated_integrals.pdf

open: iterated_integrals.pdf
	open iterated_integrals.pdf

o:
	make open

defs: iterated_integrals.tex
	defs iterated_integrals.tex
