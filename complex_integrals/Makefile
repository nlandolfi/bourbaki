complex_integrals.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex complex_integrals.tex
	make defs

make:
	make complex_integrals.pdf

open: complex_integrals.pdf
	open complex_integrals.pdf

o:
	make open

defs: complex_integrals.tex
	defs complex_integrals.tex
