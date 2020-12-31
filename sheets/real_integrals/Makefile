real_integrals.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex real_integrals.tex
	make defs

make:
	make real_integrals.pdf

open: real_integrals.pdf
	open real_integrals.pdf

o:
	make open

defs: real_integrals.tex
	defs real_integrals.tex
