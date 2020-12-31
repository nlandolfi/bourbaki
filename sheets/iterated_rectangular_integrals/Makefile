iterated_rectangular_integrals.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex iterated_rectangular_integrals.tex
	make defs

make:
	make iterated_rectangular_integrals.pdf

open: iterated_rectangular_integrals.pdf
	open iterated_rectangular_integrals.pdf

o:
	make open

defs: iterated_rectangular_integrals.tex
	defs iterated_rectangular_integrals.tex
