simple_integrals.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex simple_integrals.tex
	make defs

make:
	make simple_integrals.pdf

open: simple_integrals.pdf
	open simple_integrals.pdf

defs: simple_integrals.tex
	defs simple_integrals.tex
