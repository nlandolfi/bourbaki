nonnegative_integrals.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex nonnegative_integrals.tex
	make defs

make:
	make nonnegative_integrals.pdf

open: nonnegative_integrals.pdf
	open nonnegative_integrals.pdf

defs: nonnegative_integrals.tex
	defs nonnegative_integrals.tex
